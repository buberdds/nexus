package evmtokens

import (
	"context"
	"fmt"
	"time"

	oasisConfig "github.com/oasisprotocol/oasis-sdk/client-sdk/go/config"
	"golang.org/x/sync/errgroup"

	"github.com/oasisprotocol/oasis-indexer/analyzer"
	"github.com/oasisprotocol/oasis-indexer/analyzer/modules"
	"github.com/oasisprotocol/oasis-indexer/analyzer/util"
	"github.com/oasisprotocol/oasis-indexer/config"
	"github.com/oasisprotocol/oasis-indexer/log"
	"github.com/oasisprotocol/oasis-indexer/storage"
	"github.com/oasisprotocol/oasis-indexer/storage/oasis"
)

// The token analyzer (1) gets a list from the database of tokens to download
// the info for, (2) downloads that info, and (3) saves the info in the
// database. Note that steps (1) and (3) don't happen in a single transaction,
// so that we don't hold the tables locked and disrupt the block scanner. So
// be careful of block analyzer potentially making further updates in between.
// That's why, for example, there are separate last_mutate_round and
// last_download_round columns. The block analyzer updates last_mutate_round,
// this token analyzer updates last_download_round.

const (
	EmeraldDamaskTokensAnalyzerName = "emerald_damask_tokens"
	MaxDownloadBatch                = 20
	DownloadTimeout                 = 61 * time.Second
)

type Main struct {
	source storage.RuntimeSourceStorage
	target storage.TargetStorage
	qf     analyzer.QueryFactory
	logger *log.Logger
}

var _ analyzer.Analyzer = (*Main)(nil)

func NewMain(cfg *config.AnalyzerConfig, target storage.TargetStorage, logger *log.Logger) (*Main, error) {
	ctx := context.Background()

	// Initialize source storage.
	networkCfg := oasisConfig.Network{
		ChainContext: cfg.ChainContext,
		RPC:          cfg.RPC,
	}
	factory, err := oasis.NewClientFactory(ctx, &networkCfg, cfg.FastStartup)
	if err != nil {
		logger.Error("error creating client factory",
			"err", err,
		)
		return nil, err
	}

	network, err := analyzer.FromChainContext(cfg.ChainContext)
	if err != nil {
		return nil, err
	}

	id, err := analyzer.RuntimeEmerald.ID(network)
	if err != nil {
		return nil, err
	}
	logger.Info("Emerald runtime ID determined", "runtime_id", id)

	client, err := factory.Runtime(id)
	if err != nil {
		logger.Error("error creating runtime client",
			"err", err,
		)
		return nil, err
	}

	return &Main{
		source: client,
		target: target,
		qf:     analyzer.NewQueryFactory("oasis_3", analyzer.RuntimeEmerald.String()),
		logger: logger.With("analyzer", EmeraldDamaskTokensAnalyzerName),
	}, nil
}

type StaleToken struct {
	Addr                  string
	LastMutateRound       uint64
	LastDownloadRound     *uint64
	Type                  *modules.EVMTokenType
	AddrContextIdentifier string
	AddrContextVersion    int
	AddrData              []byte
}

func (m Main) getStaleTokens(ctx context.Context, limit int) ([]*StaleToken, error) {
	var staleTokens []*StaleToken
	rows, err := m.target.Query(ctx, m.qf.RuntimeEVMTokensAnalysisStaleQuery(), limit)
	if err != nil {
		return nil, fmt.Errorf("querying discovered tokens: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var staleToken StaleToken
		if err = rows.Scan(
			&staleToken.Addr,
			&staleToken.LastMutateRound,
			&staleToken.LastDownloadRound,
			&staleToken.Type,
			&staleToken.AddrContextIdentifier,
			&staleToken.AddrContextVersion,
			&staleToken.AddrData,
		); err != nil {
			return nil, fmt.Errorf("scanning discovered token: %w", err)
		}
		staleTokens = append(staleTokens, &staleToken)
	}
	return staleTokens, nil
}

func (m Main) processBatch(ctx context.Context) error {
	staleTokens, err := m.getStaleTokens(ctx, MaxDownloadBatch)
	if err != nil {
		return fmt.Errorf("getting discovered tokens: %w", err)
	}
	m.logger.Info("processing", "num_stale_tokens", len(staleTokens))
	if len(staleTokens) == 0 {
		return 0, nil
	}

	ctxWithTimeout, cancel := context.WithTimeout(ctx, DownloadTimeout)
	defer cancel()
	group, groupCtx := errgroup.WithContext(ctxWithTimeout)

	batch := &storage.QueryBatch{}

	for _, st := range staleTokens {
		// Redeclare `st` for unclobbered use within goroutine.
		staleToken := st
		group.Go(func() error {
			m.logger.Info("downloading", "stale_token", staleToken)
			// todo: assert that addr context is secp256k1
			if staleToken.LastDownloadRound == nil {
				tokenData, err := modules.EVMDownloadNewToken(
					groupCtx,
					m.logger,
					m.source,
					staleToken.LastMutateRound,
					staleToken.AddrData,
				)
				if err != nil {
					return fmt.Errorf("downloading new token %s: %w", staleToken.Addr, err)
				}
				if tokenData != nil {
					batch.Queue(m.qf.RuntimeEVMTokenInsertQuery(),
						staleToken.Addr,
						tokenData.Type,
						tokenData.Name,
						tokenData.Symbol,
						tokenData.Decimals,
						tokenData.TotalSupply.String(),
					)
				}
			} else if staleToken.Type != nil {
				mutable, err := modules.EVMDownloadMutatedToken(
					groupCtx,
					m.logger,
					m.source,
					staleToken.LastMutateRound,
					staleToken.AddrData,
					*staleToken.Type,
				)
				if err != nil {
					return fmt.Errorf("downloading mutated token %s: %w", staleToken.Addr, err)
				}
				batch.Queue(m.qf.RuntimeEVMTokenUpdateQuery(),
					staleToken.Addr,
					mutable.TotalSupply.String(),
				)
			}
			batch.Queue(m.qf.RuntimeEVMTokenAnalysisUpdateQuery(), staleToken.Addr, staleToken.LastMutateRound)
			return nil
		})
	}

	if err := group.Wait(); err != nil {
		return err
	}

	if err := m.target.SendBatch(ctx, batch); err != nil {
		return fmt.Errorf("sending batch: %w", err)
	}
	return nil
}

func (m Main) Start() {
	ctx := context.Background()

	backoff, err := util.NewBackoff(
		100*time.Millisecond,
		6*time.Second,
		// ^cap the timeout at the expected
		// emerald round time
	)
	if err != nil {
		m.logger.Error("error configuring indexer backoff policy",
			"err", err,
		)
		return
	}

	for {
		backoff.Wait()

		if err := m.processBatch(ctx); err != nil {
			m.logger.Error("error processing batch", "err", err)
			backoff.Failure()
			continue
		}

		backoff.Success()
	}
}

func (m Main) Name() string {
	return EmeraldDamaskTokensAnalyzerName
}
