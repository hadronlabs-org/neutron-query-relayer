package app

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"

	cosmosrelayer "github.com/cosmos/relayer/v2/relayer"

	"github.com/hadronlabs-org/neutron-query-relayer/internal/config"
	"github.com/hadronlabs-org/neutron-query-relayer/internal/kvprocessor"
	"github.com/hadronlabs-org/neutron-query-relayer/internal/raw"
	"github.com/hadronlabs-org/neutron-query-relayer/internal/relay"
	"github.com/hadronlabs-org/neutron-query-relayer/internal/submit"
	"github.com/hadronlabs-org/neutron-query-relayer/internal/tmquerier"
	"github.com/hadronlabs-org/neutron-query-relayer/internal/trusted_headers"
	nlogger "github.com/neutron-org/neutron-logger"
)

type DependencyContainer struct {
	kvProcessor          relay.KVProcessor
	proofSubmitter       relay.Submitter
	trustedHeaderFetcher relay.TrustedHeaderFetcher
	targetChain          *cosmosrelayer.Chain
	neutronChain         *cosmosrelayer.Chain
	targetQuerier        *tmquerier.Querier
}

func NewDefaultDependencyContainer(ctx context.Context,
	cfg config.NeutronQueryRelayerConfig,
	logRegistry *nlogger.Registry) (*DependencyContainer, error) {
	targetClient, err := raw.NewRPCClient(cfg.TargetChain.RPCAddr, cfg.TargetChain.Timeout)
	if err != nil {
		return nil, fmt.Errorf("could not initialize target rpc client: %w", err)
	}

	neutronClient, err := raw.NewRPCClient(cfg.NeutronChain.RPCAddr, cfg.NeutronChain.Timeout)
	if err != nil {
		return nil, fmt.Errorf("cannot create neutron client: %w", err)
	}

	connParams, err := loadConnParams(ctx, neutronClient, targetClient, cfg.NeutronChain.RESTAddr,
		cfg.NeutronChain.ConnectionID, logRegistry.Get(AppContext))
	if err != nil {
		return nil, fmt.Errorf("cannot load network params: %w", err)
	}

	targetQuerier, err := tmquerier.NewQuerier(targetClient, connParams.targetChainID)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to target chain: %w", err)
	}

	cdc := raw.MakeCodecDefault()
	keybase, err := submit.TestKeybase(connParams.neutronChainID, cfg.NeutronChain.HomeDir, codec.NewProtoCodec(cdc.InterfaceRegistry))
	if err != nil {
		return nil, fmt.Errorf("cannot initialize keybase: %w", err)
	}

	txSender, err := submit.NewTxSender(ctx,
		neutronClient,
		cdc.Marshaller,
		keybase,
		*cfg.NeutronChain,
		logRegistry.Get(TxSenderContext),
		connParams.neutronChainID)
	if err != nil {
		return nil, fmt.Errorf("cannot create tx sender: %w", err)
	}

	neutronChain, targetChain, err := loadChains(ctx, cfg, logRegistry, connParams)
	if err != nil {
		return nil, fmt.Errorf("failed to loadChains: %w", err)
	}

	proofSubmitter := submit.NewSubmitterImpl(txSender, cfg.AllowKVCallbacks, neutronChain.PathEnd.ClientID)
	trustedHeaderFetcher := trusted_headers.NewTrustedHeaderFetcher(neutronChain, targetChain, logRegistry.Get(TrustedHeadersFetcherContext))
	kvProcessor := kvprocessor.NewKVProcessor(
		trustedHeaderFetcher,
		targetQuerier,
		logRegistry.Get(KVProcessorContext),
		proofSubmitter,
		targetChain,
		neutronChain,
	)
	return &DependencyContainer{
		kvProcessor:          kvProcessor,
		proofSubmitter:       proofSubmitter,
		trustedHeaderFetcher: trustedHeaderFetcher,
		targetChain:          targetChain,
		neutronChain:         neutronChain,
		targetQuerier:        targetQuerier,
	}, nil
}

func (c DependencyContainer) GetKvProcessor() relay.KVProcessor {
	return c.kvProcessor
}

func (c DependencyContainer) GetTargetChain() *cosmosrelayer.Chain {
	return c.targetChain
}

func (c DependencyContainer) GetNeutronChain() *cosmosrelayer.Chain {
	return c.neutronChain
}

func (c DependencyContainer) GetProofSubmitter() relay.Submitter {
	return c.proofSubmitter
}

func (c DependencyContainer) GetTrustedHeaderFetcher() relay.TrustedHeaderFetcher {
	return c.trustedHeaderFetcher
}

func (c DependencyContainer) GetTargetQuerier() *tmquerier.Querier {
	return c.targetQuerier
}
