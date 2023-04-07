package file

import (
	"context"
	"encoding/json"

	"github.com/akrylysov/pogreb"
	"google.golang.org/grpc"

	coreCommon "github.com/oasisprotocol/oasis-core/go/common"
	"github.com/oasisprotocol/oasis-sdk/client-sdk/go/connection"

	"github.com/oasisprotocol/oasis-indexer/storage/oasis/nodeapi"
)

type FileRuntimeApiLite struct {
	db         pogreb.DB
	runtimeApi nodeapi.RuntimeApiLite
}

type RuntimeApiMethod func() (interface{}, error)

var _ nodeapi.RuntimeApiLite = (*FileRuntimeApiLite)(nil)

func NewFileRuntimeApiLite(filename string, runtimeID coreCommon.Namespace, grpcConn *grpc.ClientConn, sdkClient *connection.RuntimeClient) (*FileRuntimeApiLite, error) {
	db, err := pogreb.Open(filename, &pogreb.Options{BackgroundSyncInterval: -1})
	if err != nil {
		return nil, err
	}
	var runtimeApi nodeapi.RuntimeApiLite
	if grpcConn != nil && sdkClient != nil {
		runtimeApi = nodeapi.NewUniversalRuntimeApiLite(runtimeID, grpcConn, sdkClient)
	}
	return &FileRuntimeApiLite{
		db:         *db,
		runtimeApi: runtimeApi,
	}, nil
}

func (r *FileRuntimeApiLite) get(key []byte, result interface{}) error {
	res, err := r.db.Get(key)
	if err != nil {
		return err
	}
	return json.Unmarshal(res, result)
}

func (r *FileRuntimeApiLite) put(key []byte, val interface{}) error {
	valBytes, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return r.db.Put(key, valBytes)
}

func (r *FileRuntimeApiLite) updateCache(key []byte, method NodeApiMethod) error {
	exists, err := r.db.Has(key)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	val, err := method()
	if err != nil {
		return err
	}

	return r.put(key, val)
}

func (r *FileRuntimeApiLite) GetBlockHeader(ctx context.Context, round uint64) (*nodeapi.RuntimeBlockHeader, error) {
	key := generateCacheKey("GetBlockHeader", round)
	if r.runtimeApi != nil {
		if err := r.updateCache(key, func() (interface{}, error) { return r.runtimeApi.GetBlockHeader(ctx, round) }); err != nil {
			return nil, err
		}
	}
	var blockHeader nodeapi.RuntimeBlockHeader
	err := r.get(key, &blockHeader)
	if err != nil {
		return nil, err
	}
	return &blockHeader, nil
}

func (r *FileRuntimeApiLite) GetTransactionsWithResults(ctx context.Context, round uint64) ([]*nodeapi.RuntimeTransactionWithResults, error) {
	key := generateCacheKey("GetRuntimeTransactionsWithResults", round) // todo: maybe remove "runtime" from key; if we don't need to worry about collisions with the consensus api method of the same name. are we always guaranteed the diff db between consensus/runtime?
	if r.runtimeApi != nil {
		if err := r.updateCache(key, func() (interface{}, error) { return r.runtimeApi.GetTransactionsWithResults(ctx, round) }); err != nil {
			return nil, err
		}
	}
	txr := []*nodeapi.RuntimeTransactionWithResults{}
	err := r.get(key, &txr)
	if err != nil {
		return nil, err
	}
	return txr, nil
}

func (r *FileRuntimeApiLite) GetEventsRaw(ctx context.Context, round uint64) ([]*nodeapi.RuntimeEvent, error) {
	key := generateCacheKey("GetEventsRaw", round)
	if r.runtimeApi != nil {
		if err := r.updateCache(key, func() (interface{}, error) { return r.runtimeApi.GetEventsRaw(ctx, round) }); err != nil {
			return nil, err
		}
	}
	events := []*nodeapi.RuntimeEvent{}
	err := r.get(key, &events)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *FileRuntimeApiLite) EVMSimulateCall(ctx context.Context, round uint64, gasPrice []byte, gasLimit uint64, caller []byte, address []byte, value []byte, data []byte) ([]byte, error) {
	key := generateCacheKey("EVMSimulateCall", round, gasPrice, gasLimit, caller, address, value, data)
	if r.runtimeApi != nil {
		if err := r.updateCache(key, func() (interface{}, error) {
			return r.runtimeApi.EVMSimulateCall(ctx, round, gasPrice, gasLimit, caller, address, value, data)
		}); err != nil {
			return nil, err
		}
	}
	res := []byte{}
	err := r.get(key, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
