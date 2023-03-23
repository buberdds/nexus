package oasis

import (
	"context"

	"github.com/oasisprotocol/oasis-indexer/storage/oasis/nodeapi"
	config "github.com/oasisprotocol/oasis-sdk/client-sdk/go/config"
	sdkTypes "github.com/oasisprotocol/oasis-sdk/client-sdk/go/types"

	"github.com/oasisprotocol/oasis-indexer/storage"
)

// RuntimeClient is a client to a runtime. Unlike RuntimeApiLite implementations,
// which provide a 1:1 mapping to the Oasis node's runtime RPCs, this client
// is higher-level and provides a more convenient interface for the indexer.
//
// TODO: Get rid of this struct, it hardly provides any value.
type RuntimeClient struct {
	nodeApi nodeapi.RuntimeApiLite
	info    *sdkTypes.RuntimeInfo
}

// AllData returns all relevant data to the given round.
func (rc *RuntimeClient) AllData(ctx context.Context, round uint64) (*storage.RuntimeAllData, error) {
	blockHeader, err := rc.nodeApi.GetBlockHeader(ctx, round)
	if err != nil {
		return nil, err
	}
	rawEvents, err := rc.nodeApi.GetEventsRaw(ctx, round)
	if err != nil {
		return nil, err
	}
	transactionsWithResults, err := rc.nodeApi.GetTransactionsWithResults(ctx, round)
	if err != nil {
		return nil, err
	}

	data := storage.RuntimeAllData{
		Round:                   round,
		BlockHeader:             *blockHeader,
		RawEvents:               rawEvents,
		TransactionsWithResults: transactionsWithResults,
	}
	return &data, nil
}

func (rc *RuntimeClient) EVMSimulateCall(ctx context.Context, round uint64, gasPrice []byte, gasLimit uint64, caller []byte, address []byte, value []byte, data []byte) ([]byte, error) {
	return rc.nodeApi.EVMSimulateCall(ctx, round, gasPrice, gasLimit, caller, address, value, data)
}

func (rc *RuntimeClient) nativeTokenSymbol() string {
	for _, network := range config.DefaultNetworks.All {
		// Iterate over all networks and find the one that contains the runtime.
		// Any network will do; we assume that paratime IDs are unique across networks.
		for _, paratime := range network.ParaTimes.All {
			if paratime.ID == rc.info.ID.Hex() {
				return paratime.Denominations[config.NativeDenominationKey].Symbol
			}
		}
	}
	panic("Cannot find native token symbol for runtime")
}

func (rc *RuntimeClient) StringifyDenomination(d sdkTypes.Denomination) string {
	if d.IsNative() {
		return rc.nativeTokenSymbol()
	}

	return d.String()
}
