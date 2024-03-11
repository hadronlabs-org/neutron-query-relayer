package app

import (
	"context"
	"fmt"

	"github.com/neutron-org/neutron-query-relayer/internal/kvprocessor"

	"github.com/avast/retry-go/v4"
	cosmosrelayer "github.com/cosmos/relayer/v2/relayer"

	"time"

	rpcclienthttp "github.com/cometbft/cometbft/rpc/client/http"
	"go.uber.org/zap"

	nlogger "github.com/neutron-org/neutron-logger"
	"github.com/neutron-org/neutron-query-relayer/internal/chain_client"
	"github.com/neutron-org/neutron-query-relayer/internal/chain_client/querier/client/query"
	"github.com/neutron-org/neutron-query-relayer/internal/config"
	"github.com/neutron-org/neutron-query-relayer/internal/raw"
	"github.com/neutron-org/neutron-query-relayer/internal/relay"
)

var (
	Version = ""
	Commit  = ""
)

const (
	AppContext                   = "app"
	ChainClientContext           = "chain_client"
	RelayerContext               = "relayer"
	TargetChainRPCClientContext  = "target_chain_rpc"
	NeutronChainRPCClientContext = "neutron_chain_rpc"
	TargetChainProviderContext   = "target_chain_provider"
	NeutronChainProviderContext  = "neutron_chain_provider"
	TxSenderContext              = "tx_sender"
	TrustedHeadersFetcherContext = "trusted_headers_fetcher"
	KVProcessorContext           = "kv_processor"
)

// retries configuration for fetching connection info
var (
	rtyAtt = retry.Attempts(uint(5))
	rtyDel = retry.Delay(time.Second * 10)
	rtyErr = retry.LastErrorOnly(true)
)

func NewDefaultChainClient(cfg config.NeutronQueryRelayerConfig, logRegistry *nlogger.Registry) (relay.Subscriber, error) {

	subscriber, err := chain_client.NewChainClient(
		&chain_client.ChainClientConfig{
			RESTAddress:  cfg.NeutronChain.RESTAddr,
			Timeout:      cfg.NeutronChain.Timeout,
			ConnectionID: cfg.NeutronChain.ConnectionID,
		},
		logRegistry.Get(ChainClientContext),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create a NewChainClient: %s", err)
	}

	return subscriber, nil
}

// NewDefaultRelayer returns a relayer built with cfg.
func NewDefaultRelayer(
	cfg config.NeutronQueryRelayerConfig,
	logRegistry *nlogger.Registry,
	storage relay.Storage,
	deps *DependencyContainer,
) (*relay.Relayer, error) {
	var (
		kvProcessor = kvprocessor.NewKVProcessor(
			deps.GetTrustedHeaderFetcher(),
			deps.GetTargetQuerier(),
			logRegistry.Get(KVProcessorContext),
			deps.GetProofSubmitter(),
			deps.GetTargetChain(),
			deps.GetNeutronChain(),
		)
		relayer = relay.NewRelayer(
			cfg,
			storage,
			kvProcessor,
			deps.GetTargetChain(),
			logRegistry.Get(RelayerContext),
		)
	)
	return relayer, nil
}

// NewDefaultRelayer returns a relayer built with cfg.
func NewDefaultKVProcessor(
	logRegistry *nlogger.Registry,
	deps *DependencyContainer,
) (*kvprocessor.KVProcessor, error) {
	var (
		kvProcessor = kvprocessor.NewKVProcessor(
			deps.GetTrustedHeaderFetcher(),
			deps.GetTargetQuerier(),
			logRegistry.Get(KVProcessorContext),
			deps.GetProofSubmitter(),
			deps.GetTargetChain(),
			deps.GetNeutronChain(),
		)
	)
	return kvProcessor, nil
}

func loadChains(
	ctx context.Context,
	cfg config.NeutronQueryRelayerConfig,
	logRegistry *nlogger.Registry,
	connParams *connectionParams,
) (neutronChain *cosmosrelayer.Chain, targetChain *cosmosrelayer.Chain, err error) {
	targetChain, err = relay.GetTargetChain(logRegistry.Get(TargetChainProviderContext), cfg.TargetChain, connParams.targetChainID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load target chain from env: %w", err)
	}

	if err := targetChain.AddPath(connParams.targetClientID, connParams.targetConnectionID); err != nil {
		return nil, nil, fmt.Errorf("failed to AddPath to source chain: %w", err)
	}

	if err := targetChain.ChainProvider.Init(ctx); err != nil {
		return nil, nil, fmt.Errorf("failed to Init source chain provider: %w", err)
	}

	neutronChain, err = relay.GetNeutronChain(logRegistry.Get(NeutronChainProviderContext), cfg.NeutronChain, connParams.neutronChainID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load neutron chain from env: %w", err)
	}

	if err := neutronChain.AddPath(connParams.neutronClientID, cfg.NeutronChain.ConnectionID); err != nil {
		return nil, nil, fmt.Errorf("failed to AddPath to destination chain: %w", err)
	}

	if err := neutronChain.ChainProvider.Init(ctx); err != nil {
		return nil, nil, fmt.Errorf("failed to Init source chain provider: %w", err)
	}

	return neutronChain, targetChain, nil
}

type connectionParams struct {
	neutronChainID  string
	neutronClientID string

	targetChainID      string
	targetClientID     string
	targetConnectionID string
}

func loadConnParams(ctx context.Context, neutronClient, targetClient *rpcclienthttp.HTTP, neutronRestAddress string, neutronConnectionId string, logger *zap.Logger) (*connectionParams, error) {
	restClient, err := raw.NewRESTClient(neutronRestAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get newRESTClient: %w", err)
	}

	targetStatus, err := targetClient.Status(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch target chain status: %w", err)
	}

	neutronStatus, err := neutronClient.Status(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch neutron chain status: %w", err)
	}

	var queryResponse *query.IbcCoreConnectionV1ConnectionOK
	if err := retry.Do(func() error {
		var err error

		queryResponse, err = restClient.Query.IbcCoreConnectionV1Connection(&query.IbcCoreConnectionV1ConnectionParams{
			ConnectionID: neutronConnectionId,
			Context:      ctx,
		})
		if err != nil {
			return err
		}

		if queryResponse.GetPayload().Connection.Counterparty.ConnectionID == "" {
			return fmt.Errorf("empty target connection ID")
		}

		if queryResponse.GetPayload().Connection.Counterparty.ClientID == "" {
			return fmt.Errorf("empty target client ID")
		}

		return nil
	}, retry.Context(ctx), rtyAtt, rtyDel, rtyErr, retry.OnRetry(func(n uint, err error) {
		logger.Info(
			"failed to query ibc connection info", zap.Error(err))
	})); err != nil {
		return nil, err
	}

	if err != nil {
		return nil, fmt.Errorf("failed to query ibc connection info: %w", err)
	}

	connParams := connectionParams{
		neutronChainID:     neutronStatus.NodeInfo.Network,
		targetChainID:      targetStatus.NodeInfo.Network,
		neutronClientID:    queryResponse.GetPayload().Connection.ClientID,
		targetClientID:     queryResponse.GetPayload().Connection.Counterparty.ClientID,
		targetConnectionID: queryResponse.GetPayload().Connection.Counterparty.ConnectionID,
	}

	logger.Info("loaded conn params",
		zap.String("neutron_chain_id", connParams.neutronChainID),
		zap.String("target_chain_id", connParams.targetChainID),
		zap.String("neutron_client_id", connParams.neutronClientID),
		zap.String("target_client_id", connParams.targetClientID),
		zap.String("target_connection_id", connParams.targetConnectionID))

	return &connParams, nil
}
