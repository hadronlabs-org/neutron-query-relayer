package txprocessor

import (
	"context"
	"encoding/hex"
	"fmt"
	"regexp"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	tmtypes "github.com/tendermint/tendermint/types"
	"go.uber.org/zap"

	neutronmetrics "github.com/neutron-org/neutron-query-relayer/cmd/neutron_query_relayer/metrics"
	"github.com/neutron-org/neutron-query-relayer/internal/relay"
	neutrontypes "github.com/neutron-org/neutron/x/interchainqueries/types"
)

type TXProcessor struct {
	trustedHeaderFetcher        relay.TrustedHeaderFetcher
	storage                     relay.Storage
	submitter                   relay.Submitter
	logger                      *zap.Logger
	checkSubmittedTxStatusDelay time.Duration
	ignoreErrorsRegexp          *regexp.Regexp
}

func NewTxProcessor(
	trustedHeaderFetcher relay.TrustedHeaderFetcher,
	storage relay.Storage,
	submitter relay.Submitter,
	logger *zap.Logger,
	checkSubmittedTxStatusDelay time.Duration,
	ignoreErrorsRegexp string,
) TXProcessor {
	txProcessor := TXProcessor{
		trustedHeaderFetcher:        trustedHeaderFetcher,
		storage:                     storage,
		submitter:                   submitter,
		logger:                      logger,
		checkSubmittedTxStatusDelay: checkSubmittedTxStatusDelay,
		ignoreErrorsRegexp:          regexp.MustCompile(ignoreErrorsRegexp),
	}

	return txProcessor
}

func (r TXProcessor) ProcessAndSubmit(
	ctx context.Context,
	queryID uint64,
	tx relay.Transaction,
	submittedTxsTasksQueue chan relay.PendingSubmittedTxInfo,
) error {
	hash := hex.EncodeToString(tmtypes.Tx(tx.Tx.Data).Hash())
	txExists, err := r.storage.TxExists(queryID, hash)
	if err != nil {
		return fmt.Errorf("failed to check tx existence: %w", err)
	}

	if txExists {
		r.logger.Debug("transaction already submitted",
			zap.Uint64("query_id", queryID),
			zap.String("hash", hash),
			zap.Uint64("height", tx.Height))
		return nil
	}

	block, err := r.txToBlock(ctx, tx)
	if err != nil {
		return fmt.Errorf("failed to prepare block: %w", err)
	}

	if err = r.submitTxWithProofs(ctx, queryID, block, submittedTxsTasksQueue); err != nil {
		return fmt.Errorf("failed to submit block: %w", err)
	}
	return nil
}

func (r TXProcessor) submitTxWithProofs(
	ctx context.Context,
	queryID uint64,
	block *neutrontypes.Block,
	submittedTxsTasksQueue chan relay.PendingSubmittedTxInfo,
) error {
	proofStart := time.Now()
	hash := hex.EncodeToString(tmtypes.Tx(block.Tx.Data).Hash())
	neutronTxHash, err := r.submitter.SubmitTxProof(ctx, queryID, block)

	if err == nil {
		neutronmetrics.AddSuccessProof(string(neutrontypes.InterchainQueryTypeTX), time.Since(proofStart).Seconds())
		err = r.storage.SetTxStatus(queryID, hash, neutronTxHash, relay.SubmittedTxInfo{
			Status: relay.Submitted,
		})
		if err != nil {
			return fmt.Errorf("failed to store tx: %w", err)
		}

		// We submit the PendingSubmittedTxInfo only after checkSubmittedTxStatusDelay to reduce the possibility of
		// unsuccessful checks (the block is 100% not ready here yet).
		go func() {
			var t = time.NewTimer(r.checkSubmittedTxStatusDelay)
			select {
			case <-t.C:
				submittedTxsTasksQueue <- relay.PendingSubmittedTxInfo{
					QueryID:         queryID,
					SubmittedTxHash: hash,
					NeutronHash:     neutronTxHash,
				}
			case <-ctx.Done():
				r.logger.Info("Cancelled PendingSubmittedTxInfo delayed checking",
					zap.Uint64("query_id", queryID),
					zap.String("submitted_tx_hash", hash))
			}
		}()

		r.logger.Info("proof for query_id submitted successfully", zap.Uint64("query_id", queryID))
		return nil
	}

	// check error with regexp
	if !r.ignoreErrorsRegexp.MatchString(err.Error()) {
		r.logger.Error("failed to submit tx proof", zap.Error(err))
		return err
	}

	neutronmetrics.AddFailedProof(string(neutrontypes.InterchainQueryTypeTX), time.Since(proofStart).Seconds())
	errSetStatus := r.storage.SetTxStatus(
		queryID, hash, neutronTxHash, relay.SubmittedTxInfo{Status: relay.ErrorOnSubmit, Message: err.Error()})
	if errSetStatus != nil {
		return fmt.Errorf("failed to store tx: %w", errSetStatus)
	}

	r.logger.Error("\"could not submit proof", zap.Error(err), zap.Uint64("query_id", queryID))
	return nil
}

func (r TXProcessor) txToBlock(ctx context.Context, tx relay.Transaction) (*neutrontypes.Block, error) {
	packedHeader, packedNextHeader, err := r.prepareHeaders(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare headers: %w", err)
	}
	block := neutrontypes.Block{
		Header:          packedHeader,
		NextBlockHeader: packedNextHeader,
		Tx:              tx.Tx,
	}
	return &block, nil
}

func (r TXProcessor) prepareHeaders(ctx context.Context, txStruct relay.Transaction) (
	packedHeader *codectypes.Any, packedNextHeader *codectypes.Any, err error) {
	packedHeader, packedNextHeader, err = r.trustedHeaderFetcher.FetchTrustedHeadersForHeights(ctx, txStruct.Height)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get header for src chain: %w", err)
	}

	return
}
