package history

import (
	"context"
	"fmt"

	sdkConfig "github.com/oasisprotocol/oasis-sdk/client-sdk/go/config"

	"github.com/oasisprotocol/nexus/common"
	"github.com/oasisprotocol/nexus/config"
	"github.com/oasisprotocol/nexus/storage/oasis/connections"
	"github.com/oasisprotocol/nexus/storage/oasis/nodeapi"
)

var _ nodeapi.RuntimeApiLite = (*HistoryRuntimeApiLite)(nil)

type HistoryRuntimeApiLite struct {
	Runtime common.Runtime
	History *config.History
	APIs    map[string]nodeapi.RuntimeApiLite
}

func NewHistoryRuntimeApiLite(ctx context.Context, history *config.History, sdkPT *sdkConfig.ParaTime, nodes map[string]*config.ArchiveConfig, fastStartup bool, runtime common.Runtime) (*HistoryRuntimeApiLite, error) {
	apis := map[string]nodeapi.RuntimeApiLite{}
	for _, record := range history.Records {
		if archiveConfig, ok := nodes[record.ArchiveName]; ok {
			sdkConn, err := connections.SDKConnect(ctx, record.ChainContext, archiveConfig.ResolvedRuntimeNode(runtime), fastStartup)
			if err != nil {
				return nil, err
			}
			sdkClient := sdkConn.Runtime(sdkPT)
			rawConn := connections.NewLazyGrpcConn(*archiveConfig.ResolvedRuntimeNode(runtime))
			apis[record.ArchiveName] = nodeapi.NewUniversalRuntimeApiLite(sdkPT.Namespace(), rawConn, &sdkClient)
		}
	}
	return &HistoryRuntimeApiLite{
		Runtime: runtime,
		History: history,
		APIs:    apis,
	}, nil
}

func (rc *HistoryRuntimeApiLite) Close() error {
	var firstErr error
	for _, api := range rc.APIs {
		if err := api.Close(); err != nil && firstErr == nil {
			firstErr = err
			// Do not return yet; keep closing others.
		}
	}
	if firstErr != nil {
		return fmt.Errorf("closing apis failed, first encountered error was: %w", firstErr)
	}
	return nil
}

func (rc *HistoryRuntimeApiLite) APIForRound(round uint64) (nodeapi.RuntimeApiLite, error) {
	record, err := rc.History.RecordForRuntimeRound(rc.Runtime, round)
	if err != nil {
		return nil, fmt.Errorf("determining archive: %w", err)
	}
	api, ok := rc.APIs[record.ArchiveName]
	if !ok {
		return nil, fmt.Errorf("archive %s has no node configured", record.ArchiveName)
	}
	return api, nil
}

func (rc *HistoryRuntimeApiLite) GetEventsRaw(ctx context.Context, round uint64) ([]nodeapi.RuntimeEvent, error) {
	api, err := rc.APIForRound(round)
	if err != nil {
		return nil, fmt.Errorf("getting api for runtime %s round %d: %w", rc.Runtime, round, err)
	}
	return api.GetEventsRaw(ctx, round)
}

func (rc *HistoryRuntimeApiLite) EVMSimulateCall(ctx context.Context, round uint64, gasPrice []byte, gasLimit uint64, caller []byte, address []byte, value []byte, data []byte) (*nodeapi.FallibleResponse, error) {
	api, err := rc.APIForRound(round)
	if err != nil {
		return nil, fmt.Errorf("getting api for runtime %s round %d: %w", rc.Runtime, round, err)
	}
	return api.EVMSimulateCall(ctx, round, gasPrice, gasLimit, caller, address, value, data)
}

func (rc *HistoryRuntimeApiLite) EVMGetCode(ctx context.Context, round uint64, address []byte) ([]byte, error) {
	api, err := rc.APIForRound(round)
	if err != nil {
		return nil, fmt.Errorf("getting api for runtime %s round %d: %w", rc.Runtime, round, err)
	}
	return api.EVMGetCode(ctx, round, address)
}

func (rc *HistoryRuntimeApiLite) GetBlockHeader(ctx context.Context, round uint64) (*nodeapi.RuntimeBlockHeader, error) {
	api, err := rc.APIForRound(round)
	if err != nil {
		return nil, fmt.Errorf("getting api for runtime %s round %d: %w", rc.Runtime, round, err)
	}
	return api.GetBlockHeader(ctx, round)
}

func (rc *HistoryRuntimeApiLite) GetNativeBalance(ctx context.Context, round uint64, addr nodeapi.Address) (*common.BigInt, error) {
	api, err := rc.APIForRound(round)
	if err != nil {
		return nil, fmt.Errorf("getting api for runtime %s round %d: %w", rc.Runtime, round, err)
	}
	return api.GetNativeBalance(ctx, round, addr)
}

func (rc *HistoryRuntimeApiLite) GetTransactionsWithResults(ctx context.Context, round uint64) ([]nodeapi.RuntimeTransactionWithResults, error) {
	api, err := rc.APIForRound(round)
	if err != nil {
		return nil, fmt.Errorf("getting api for runtime %s round %d: %w", rc.Runtime, round, err)
	}
	return api.GetTransactionsWithResults(ctx, round)
}
