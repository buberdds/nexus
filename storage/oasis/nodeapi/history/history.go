package history

import (
	"context"
	"fmt"

	beacon "github.com/oasisprotocol/oasis-core/go/beacon/api"
	"github.com/oasisprotocol/oasis-core/go/common"
	consensus "github.com/oasisprotocol/oasis-core/go/consensus/api"
	genesis "github.com/oasisprotocol/oasis-core/go/genesis/api"

	"github.com/oasisprotocol/oasis-indexer/config"
	"github.com/oasisprotocol/oasis-indexer/storage/oasis/nodeapi"
	"github.com/oasisprotocol/oasis-indexer/storage/oasis/nodeapi/cobalt"
	"github.com/oasisprotocol/oasis-indexer/storage/oasis/nodeapi/damask"
)

var _ nodeapi.ConsensusApiLite = (*HistoryConsensusApiLite)(nil)

type APIConstructor func(ctx context.Context, chainContext string, nodeConfig *config.NodeConfig, fastStartup bool) (nodeapi.ConsensusApiLite, error)

var APIConstructors = map[string]APIConstructor{
	"damask": func(ctx context.Context, chainContext string, nodeConfig *config.NodeConfig, fastStartup bool) (nodeapi.ConsensusApiLite, error) {
		sdkConn, err := SDKConnect(ctx, chainContext, nodeConfig, fastStartup)
		if err != nil {
			return nil, err
		}
		return damask.NewDamaskConsensusApiLite(sdkConn.Consensus()), nil
	},
	"cobalt": func(ctx context.Context, chainContext string, nodeConfig *config.NodeConfig, fastStartup bool) (nodeapi.ConsensusApiLite, error) {
		rawGRPCConnection, err := ConnectNoVerify(nodeConfig.RPC)
		if err != nil {
			return nil, fmt.Errorf("indexer ConnectNoVerify: %w", err)
		}
		return cobalt.NewCobaltConsensusApiLite(rawGRPCConnection), nil
	},
}

type HistoryConsensusApiLite struct {
	History *config.History
	// APIs. Keys are "archive names," which are named after mainnet releases,
	// in lowercase e.g. "cobalt" and "damask."
	APIs map[string]nodeapi.ConsensusApiLite
}

func NewHistoryConsensusApiLite(ctx context.Context, history *config.History, nodes map[string]*config.NodeConfig, fastStartup bool) (*HistoryConsensusApiLite, error) {
	APIs := map[string]nodeapi.ConsensusApiLite{}
	for _, record := range history.Records {
		if nodeConfig, ok := nodes[record.ArchiveName]; ok {
			apiConstructor := APIConstructors[record.ArchiveName]
			if apiConstructor == nil {
				return nil, fmt.Errorf("historical API for archive %s not implemented", record.ArchiveName)
			}
			api, err := apiConstructor(ctx, record.ChainContext, nodeConfig, fastStartup)
			if err != nil {
				return nil, fmt.Errorf("connecting to archive %s: %w", record.ArchiveName, err)
			}
			APIs[record.ArchiveName] = api
		}
	}
	return &HistoryConsensusApiLite{
		History: history,
		APIs:    APIs,
	}, nil
}

func (c *HistoryConsensusApiLite) APIForHeight(height int64) (nodeapi.ConsensusApiLite, error) {
	record, err := c.History.RecordForHeight(height)
	if err != nil {
		return nil, fmt.Errorf("dertermining archive: %w", err)
	}
	api, ok := c.APIs[record.ArchiveName]
	if !ok {
		return nil, fmt.Errorf("archive %s for has no node configured", record.ArchiveName)
	}
	return api, nil
}

func (c *HistoryConsensusApiLite) GetGenesisDocument(ctx context.Context) (*genesis.Document, error) {
	// Use latest.
	height := consensus.HeightLatest
	api, err := c.APIForHeight(height)
	if err != nil {
		return nil, fmt.Errorf("getting api for height %d: %w", height, err)
	}
	return api.GetGenesisDocument(ctx)
}

func (c *HistoryConsensusApiLite) StateToGenesis(ctx context.Context, height int64) (*genesis.Document, error) {
	api, err := c.APIForHeight(height)
	if err != nil {
		return nil, fmt.Errorf("getting api for height %d: %w", height, err)
	}
	return api.StateToGenesis(ctx, height)
}

func (c *HistoryConsensusApiLite) GetBlock(ctx context.Context, height int64) (*consensus.Block, error) {
	api, err := c.APIForHeight(height)
	if err != nil {
		return nil, fmt.Errorf("getting api for height %d: %w", height, err)
	}
	return api.GetBlock(ctx, height)
}

func (c *HistoryConsensusApiLite) GetTransactionsWithResults(ctx context.Context, height int64) ([]nodeapi.TransactionWithResults, error) {
	api, err := c.APIForHeight(height)
	if err != nil {
		return nil, fmt.Errorf("getting api for height %d: %w", height, err)
	}
	return api.GetTransactionsWithResults(ctx, height)
}

func (c *HistoryConsensusApiLite) GetEpoch(ctx context.Context, height int64) (beacon.EpochTime, error) {
	api, err := c.APIForHeight(height)
	if err != nil {
		return 0, fmt.Errorf("getting api for height %d: %w", height, err)
	}
	return api.GetEpoch(ctx, height)
}

func (c *HistoryConsensusApiLite) RegistryEvents(ctx context.Context, height int64) ([]nodeapi.Event, error) {
	api, err := c.APIForHeight(height)
	if err != nil {
		return nil, fmt.Errorf("getting api for height %d: %w", height, err)
	}
	return api.RegistryEvents(ctx, height)
}

func (c *HistoryConsensusApiLite) StakingEvents(ctx context.Context, height int64) ([]nodeapi.Event, error) {
	api, err := c.APIForHeight(height)
	if err != nil {
		return nil, fmt.Errorf("getting api for height %d: %w", height, err)
	}
	return api.StakingEvents(ctx, height)
}

func (c *HistoryConsensusApiLite) GovernanceEvents(ctx context.Context, height int64) ([]nodeapi.Event, error) {
	api, err := c.APIForHeight(height)
	if err != nil {
		return nil, fmt.Errorf("getting api for height %d: %w", height, err)
	}
	return api.GovernanceEvents(ctx, height)
}

func (c *HistoryConsensusApiLite) RoothashEvents(ctx context.Context, height int64) ([]nodeapi.Event, error) {
	api, err := c.APIForHeight(height)
	if err != nil {
		return nil, fmt.Errorf("getting api for height %d: %w", height, err)
	}
	return api.RoothashEvents(ctx, height)
}

func (c *HistoryConsensusApiLite) GetValidators(ctx context.Context, height int64) ([]nodeapi.Validator, error) {
	api, err := c.APIForHeight(height)
	if err != nil {
		return nil, fmt.Errorf("getting api for height %d: %w", height, err)
	}
	return api.GetValidators(ctx, height)
}

func (c *HistoryConsensusApiLite) GetCommittees(ctx context.Context, height int64, runtimeID common.Namespace) ([]nodeapi.Committee, error) {
	api, err := c.APIForHeight(height)
	if err != nil {
		return nil, fmt.Errorf("getting api for height %d: %w", height, err)
	}
	return api.GetCommittees(ctx, height, runtimeID)
}

func (c *HistoryConsensusApiLite) GetProposal(ctx context.Context, height int64, proposalID uint64) (*nodeapi.Proposal, error) {
	api, err := c.APIForHeight(height)
	if err != nil {
		return nil, fmt.Errorf("getting api for height %d: %w", height, err)
	}
	return api.GetProposal(ctx, height, proposalID)
}
