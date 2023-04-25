package file

import (
	"context"

	"github.com/akrylysov/pogreb"
	beacon "github.com/oasisprotocol/oasis-core/go/beacon/api"
	coreCommon "github.com/oasisprotocol/oasis-core/go/common"
	consensus "github.com/oasisprotocol/oasis-core/go/consensus/api"
	genesis "github.com/oasisprotocol/oasis-core/go/genesis/api"

	"github.com/oasisprotocol/oasis-indexer/storage/oasis/nodeapi"
)

// FileConsensusApiLite provides access to the consensus API of an Oasis node.
// Since FileConsensusApiLite is backed by a file containing the cached responses
// to `ConsensusApiLite` calls, this data is inherently compatible with the
// current indexer and can thus handle heights from both Cobalt/Damask.
type FileConsensusApiLite struct {
	db           KVStore
	consensusApi nodeapi.ConsensusApiLite
}

var _ nodeapi.ConsensusApiLite = (*FileConsensusApiLite)(nil)

func NewFileConsensusApiLite(filename string, consensusApi nodeapi.ConsensusApiLite) (*FileConsensusApiLite, error) {
	db, err := pogreb.Open(filename, &pogreb.Options{BackgroundSyncInterval: -1})
	if err != nil {
		return nil, err
	}
	return &FileConsensusApiLite{
		db:           KVStore{*db},
		consensusApi: consensusApi,
	}, nil
}

func (c *FileConsensusApiLite) GetGenesisDocument(ctx context.Context) (*genesis.Document, error) {
	return GetFromCacheOrCall(
		c.db, false,
		generateCacheKey("GetGenesisDocument"),
		func() (*genesis.Document, error) { return c.consensusApi.GetGenesisDocument(ctx) },
	)
}

func (c *FileConsensusApiLite) StateToGenesis(ctx context.Context, height int64) (*genesis.Document, error) {
	return GetFromCacheOrCall(
		c.db, height == consensus.HeightLatest,
		generateCacheKey("StateToGenesis", height),
		func() (*genesis.Document, error) { return c.consensusApi.StateToGenesis(ctx, height) },
	)
}

func (c *FileConsensusApiLite) GetBlock(ctx context.Context, height int64) (*consensus.Block, error) {
	return GetFromCacheOrCall(
		c.db, height == consensus.HeightLatest,
		generateCacheKey("GetBlock", height),
		func() (*consensus.Block, error) { return c.consensusApi.GetBlock(ctx, height) },
	)
}

func (c *FileConsensusApiLite) GetTransactionsWithResults(ctx context.Context, height int64) ([]nodeapi.TransactionWithResults, error) {
	return GetSliceFromCacheOrCall(
		c.db, height == consensus.HeightLatest,
		generateCacheKey("GetTransactionsWithResults", height),
		func() ([]nodeapi.TransactionWithResults, error) {
			return c.consensusApi.GetTransactionsWithResults(ctx, height)
		},
	)
}

func (c *FileConsensusApiLite) GetEpoch(ctx context.Context, height int64) (beacon.EpochTime, error) {
	time, err := GetFromCacheOrCall(
		c.db, height == consensus.HeightLatest,
		generateCacheKey("GetEpoch", height),
		func() (*beacon.EpochTime, error) {
			time, err := c.consensusApi.GetEpoch(ctx, height)
			return &time, err
		},
	)
	return *time, err
}

func (c *FileConsensusApiLite) RegistryEvents(ctx context.Context, height int64) ([]nodeapi.Event, error) {
	return GetSliceFromCacheOrCall(
		c.db, height == consensus.HeightLatest,
		generateCacheKey("RegistryEvents", height),
		func() ([]nodeapi.Event, error) { return c.consensusApi.RegistryEvents(ctx, height) },
	)
}

func (c *FileConsensusApiLite) StakingEvents(ctx context.Context, height int64) ([]nodeapi.Event, error) {
	return GetSliceFromCacheOrCall(
		c.db, height == consensus.HeightLatest,
		generateCacheKey("StakingEvents", height),
		func() ([]nodeapi.Event, error) { return c.consensusApi.StakingEvents(ctx, height) },
	)
}

func (c *FileConsensusApiLite) GovernanceEvents(ctx context.Context, height int64) ([]nodeapi.Event, error) {
	return GetSliceFromCacheOrCall(
		c.db, height == consensus.HeightLatest,
		generateCacheKey("GovernanceEvents", height),
		func() ([]nodeapi.Event, error) { return c.consensusApi.GovernanceEvents(ctx, height) },
	)
}

func (c *FileConsensusApiLite) RoothashEvents(ctx context.Context, height int64) ([]nodeapi.Event, error) {
	return GetSliceFromCacheOrCall(
		c.db, height == consensus.HeightLatest,
		generateCacheKey("RoothashEvents", height),
		func() ([]nodeapi.Event, error) { return c.consensusApi.RoothashEvents(ctx, height) },
	)
}

func (c *FileConsensusApiLite) GetValidators(ctx context.Context, height int64) ([]nodeapi.Validator, error) {
	return GetSliceFromCacheOrCall(
		c.db, height == consensus.HeightLatest,
		generateCacheKey("GetValidators", height),
		func() ([]nodeapi.Validator, error) { return c.consensusApi.GetValidators(ctx, height) },
	)
}

func (c *FileConsensusApiLite) GetCommittees(ctx context.Context, height int64, runtimeID coreCommon.Namespace) ([]nodeapi.Committee, error) {
	return GetSliceFromCacheOrCall(
		c.db, height == consensus.HeightLatest,
		generateCacheKey("GetCommittee", height, runtimeID),
		func() ([]nodeapi.Committee, error) { return c.consensusApi.GetCommittees(ctx, height, runtimeID) },
	)
}

func (c *FileConsensusApiLite) GetProposal(ctx context.Context, height int64, proposalID uint64) (*nodeapi.Proposal, error) {
	return GetFromCacheOrCall(
		c.db, height == consensus.HeightLatest,
		generateCacheKey("GetProposal", height, proposalID),
		func() (*nodeapi.Proposal, error) { return c.consensusApi.GetProposal(ctx, height, proposalID) },
	)
}
