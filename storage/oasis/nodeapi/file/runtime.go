package file

import (
	"context"

	roothash "github.com/oasisprotocol/nexus/coreapi/v22.2.11/roothash/api"

	"github.com/oasisprotocol/nexus/common"
	"github.com/oasisprotocol/nexus/log"
	"github.com/oasisprotocol/nexus/metrics"
	"github.com/oasisprotocol/nexus/storage/oasis/nodeapi"
)

type FileRuntimeApiLite struct {
	runtime    common.Runtime
	db         KVStore
	runtimeApi nodeapi.RuntimeApiLite
}

type RuntimeApiMethod func() (interface{}, error)

var _ nodeapi.RuntimeApiLite = (*FileRuntimeApiLite)(nil)

func NewFileRuntimeApiLite(runtime common.Runtime, cacheDir string, runtimeApi nodeapi.RuntimeApiLite) (*FileRuntimeApiLite, error) {
	db, err := OpenKVStore(
		log.NewDefaultLogger("cached-node-api").With("runtime", runtime),
		cacheDir,
		common.Ptr(metrics.NewDefaultStorageMetrics(string(runtime))),
	)
	if err != nil {
		return nil, err
	}
	return &FileRuntimeApiLite{
		runtime:    runtime,
		db:         db,
		runtimeApi: runtimeApi,
	}, nil
}

func (r *FileRuntimeApiLite) Close() error {
	// Close all resources and return the first encountered error, if any.
	var firstErr error
	if r.runtimeApi != nil {
		firstErr = r.runtimeApi.Close()
	}
	if err := r.db.Close(); err != nil && firstErr == nil {
		firstErr = err
	}
	return firstErr
}

func (r *FileRuntimeApiLite) GetBlockHeader(ctx context.Context, round uint64) (*nodeapi.RuntimeBlockHeader, error) {
	return GetFromCacheOrCall(
		r.db, round == roothash.RoundLatest,
		generateCacheKey("GetBlockHeader", r.runtime, round),
		func() (*nodeapi.RuntimeBlockHeader, error) { return r.runtimeApi.GetBlockHeader(ctx, round) },
	)
}

func (r *FileRuntimeApiLite) GetTransactionsWithResults(ctx context.Context, round uint64) ([]nodeapi.RuntimeTransactionWithResults, error) {
	return GetSliceFromCacheOrCall(
		r.db, round == roothash.RoundLatest,
		generateCacheKey("GetTransactionsWithResults", r.runtime, round),
		func() ([]nodeapi.RuntimeTransactionWithResults, error) {
			return r.runtimeApi.GetTransactionsWithResults(ctx, round)
		},
	)
}

func (r *FileRuntimeApiLite) GetEventsRaw(ctx context.Context, round uint64) ([]nodeapi.RuntimeEvent, error) {
	return GetSliceFromCacheOrCall(
		r.db, round == roothash.RoundLatest,
		generateCacheKey("GetEventsRaw", r.runtime, round),
		func() ([]nodeapi.RuntimeEvent, error) { return r.runtimeApi.GetEventsRaw(ctx, round) },
	)
}

func (r *FileRuntimeApiLite) GetNativeBalance(ctx context.Context, round uint64, addr nodeapi.Address) (*common.BigInt, error) {
	return GetFromCacheOrCall(
		r.db, round == roothash.RoundLatest,
		generateCacheKey("GetNativeBalance", r.runtime, round, addr),
		func() (*common.BigInt, error) { return r.runtimeApi.GetNativeBalance(ctx, round, addr) },
	)
}

func (r *FileRuntimeApiLite) EVMSimulateCall(ctx context.Context, round uint64, gasPrice []byte, gasLimit uint64, caller []byte, address []byte, value []byte, data []byte) (*nodeapi.FallibleResponse, error) {
	return GetFromCacheOrCall(
		r.db, round == roothash.RoundLatest,
		generateCacheKey("EVMSimulateCall", r.runtime, round, gasPrice, gasLimit, caller, address, value, data),
		func() (*nodeapi.FallibleResponse, error) {
			return r.runtimeApi.EVMSimulateCall(ctx, round, gasPrice, gasLimit, caller, address, value, data)
		},
	)
}

func (r *FileRuntimeApiLite) EVMGetCode(ctx context.Context, round uint64, address []byte) ([]byte, error) {
	return GetSliceFromCacheOrCall(
		r.db, round == roothash.RoundLatest,
		generateCacheKey("EVMGetCode", r.runtime, round, address),
		func() ([]byte, error) {
			return r.runtimeApi.EVMGetCode(ctx, round, address)
		},
	)
}
