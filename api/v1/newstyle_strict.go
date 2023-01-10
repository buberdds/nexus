package v1

import (
	"context"

	apiTypes "github.com/oasisprotocol/oasis-indexer/api/v1/types"
	"github.com/oasisprotocol/oasis-indexer/log"
)

type StrictServerImpl struct {
	client *storageClient
	logger *log.Logger
}

var _ apiTypes.StrictServerInterface = (*StrictServerImpl)(nil)

func NewStrictServerImpl(client *storageClient, logger *log.Logger) *StrictServerImpl {
	return &StrictServerImpl{
		client: client,
		logger: logger,
	}
}

// Stubs generated with:
//
//	sed -n '/type StrictServerInterface interface/,/^\}/p' api/v1/types/server.gen.go | grep -v // | head -n-1 | tail -n+2 | sed -E 's/^\s+(\w+)(.*)/func (srv *StrictServerImpl) \1\2 { return apiTypes.\1200JSONResponse{}, nil } /g; s/[a-zA-Z]+(Params|RequestObject|ResponseObject)/apiTypes.\0/g;'

func (srv *StrictServerImpl) Get(ctx context.Context, request apiTypes.GetRequestObject) (apiTypes.GetResponseObject, error) {
	status, err := srv.client.Storage.Status(ctx)
	if err != nil {
		return nil, err
	}
	return apiTypes.Get200JSONResponse(*status), nil
}

func (srv *StrictServerImpl) GetConsensusAccounts(ctx context.Context, request apiTypes.GetConsensusAccountsRequestObject) (apiTypes.GetConsensusAccountsResponseObject, error) {
	accounts, err := srv.client.Storage.Accounts(ctx, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusAccounts200JSONResponse(*accounts), nil
}

func (srv *StrictServerImpl) GetConsensusAccountsAddress(ctx context.Context, request apiTypes.GetConsensusAccountsAddressRequestObject) (apiTypes.GetConsensusAccountsAddressResponseObject, error) {
	account, err := srv.client.Storage.Account(ctx, request.Address)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusAccountsAddress200JSONResponse(*account), nil
}

func (srv *StrictServerImpl) GetConsensusAccountsAddressDebondingDelegations(ctx context.Context, request apiTypes.GetConsensusAccountsAddressDebondingDelegationsRequestObject) (apiTypes.GetConsensusAccountsAddressDebondingDelegationsResponseObject, error) {
	return apiTypes.GetConsensusAccountsAddressDebondingDelegations200JSONResponse{}, nil
}

func (srv *StrictServerImpl) GetConsensusAccountsAddressDelegations(ctx context.Context, request apiTypes.GetConsensusAccountsAddressDelegationsRequestObject) (apiTypes.GetConsensusAccountsAddressDelegationsResponseObject, error) {
	delegations, err := srv.client.Storage.Delegations(ctx, request.Address, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusAccountsAddressDelegations200JSONResponse(*delegations), nil
}

func (srv *StrictServerImpl) GetConsensusBlocks(ctx context.Context, request apiTypes.GetConsensusBlocksRequestObject) (apiTypes.GetConsensusBlocksResponseObject, error) {
	blocks, err := srv.client.Storage.Blocks(ctx, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusBlocks200JSONResponse(*blocks), nil
}

func (srv *StrictServerImpl) GetConsensusBlocksHeight(ctx context.Context, request apiTypes.GetConsensusBlocksHeightRequestObject) (apiTypes.GetConsensusBlocksHeightResponseObject, error) {
	block, err := srv.client.Storage.Block(ctx, request.Height)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusBlocksHeight200JSONResponse(*block), nil
}

func (srv *StrictServerImpl) GetConsensusEntities(ctx context.Context, request apiTypes.GetConsensusEntitiesRequestObject) (apiTypes.GetConsensusEntitiesResponseObject, error) {
	entities, err := srv.client.Storage.Entities(ctx, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusEntities200JSONResponse(*entities), nil
}

func (srv *StrictServerImpl) GetConsensusEntitiesEntityId(ctx context.Context, request apiTypes.GetConsensusEntitiesEntityIdRequestObject) (apiTypes.GetConsensusEntitiesEntityIdResponseObject, error) {
	entity, err := srv.client.Storage.Entity(ctx, request.EntityId)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusEntitiesEntityId200JSONResponse(*entity), nil
}

func (srv *StrictServerImpl) GetConsensusEntitiesEntityIdNodes(ctx context.Context, request apiTypes.GetConsensusEntitiesEntityIdNodesRequestObject) (apiTypes.GetConsensusEntitiesEntityIdNodesResponseObject, error) {
	nodes, err := srv.client.Storage.EntityNodes(ctx, request.EntityId, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusEntitiesEntityIdNodes200JSONResponse(*nodes), nil
}

func (srv *StrictServerImpl) GetConsensusEntitiesEntityIdNodesNodeId(ctx context.Context, request apiTypes.GetConsensusEntitiesEntityIdNodesNodeIdRequestObject) (apiTypes.GetConsensusEntitiesEntityIdNodesNodeIdResponseObject, error) {
	node, err := srv.client.Storage.EntityNode(ctx, request.EntityId, request.NodeId)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusEntitiesEntityIdNodesNodeId200JSONResponse(*node), nil
}

func (srv *StrictServerImpl) GetConsensusEpochs(ctx context.Context, request apiTypes.GetConsensusEpochsRequestObject) (apiTypes.GetConsensusEpochsResponseObject, error) {
	epochs, err := srv.client.Storage.Epochs(ctx, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusEpochs200JSONResponse(*epochs), nil
}

func (srv *StrictServerImpl) GetConsensusEpochsEpoch(ctx context.Context, request apiTypes.GetConsensusEpochsEpochRequestObject) (apiTypes.GetConsensusEpochsEpochResponseObject, error) {
	epoch, err := srv.client.Storage.Epoch(ctx, request.Epoch)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusEpochsEpoch200JSONResponse(*epoch), nil
}

func (srv *StrictServerImpl) GetConsensusEvents(ctx context.Context, request apiTypes.GetConsensusEventsRequestObject) (apiTypes.GetConsensusEventsResponseObject, error) {
	events, err := srv.client.Storage.Events(ctx, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusEvents200JSONResponse(*events), nil
}

func (srv *StrictServerImpl) GetConsensusProposals(ctx context.Context, request apiTypes.GetConsensusProposalsRequestObject) (apiTypes.GetConsensusProposalsResponseObject, error) {
	proposals, err := srv.client.Storage.Proposals(ctx, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusProposals200JSONResponse(*proposals), nil
}

func (srv *StrictServerImpl) GetConsensusProposalsProposalId(ctx context.Context, request apiTypes.GetConsensusProposalsProposalIdRequestObject) (apiTypes.GetConsensusProposalsProposalIdResponseObject, error) {
	proposal, err := srv.client.Storage.Proposal(ctx, request.ProposalId)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusProposalsProposalId200JSONResponse(*proposal), nil
}

func (srv *StrictServerImpl) GetConsensusProposalsProposalIdVotes(ctx context.Context, request apiTypes.GetConsensusProposalsProposalIdVotesRequestObject) (apiTypes.GetConsensusProposalsProposalIdVotesResponseObject, error) {
	votes, err := srv.client.Storage.ProposalVotes(ctx, request.ProposalId, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusProposalsProposalIdVotes200JSONResponse(*votes), nil
}

func (srv *StrictServerImpl) GetConsensusStatsTxVolume(ctx context.Context, request apiTypes.GetConsensusStatsTxVolumeRequestObject) (apiTypes.GetConsensusStatsTxVolumeResponseObject, error) {
	volumeList, err := srv.client.Storage.TxVolumes(ctx, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusStatsTxVolume200JSONResponse(*volumeList), nil
}

func (srv *StrictServerImpl) GetConsensusTransactions(ctx context.Context, request apiTypes.GetConsensusTransactionsRequestObject) (apiTypes.GetConsensusTransactionsResponseObject, error) {
	txs, err := srv.client.Storage.Transactions(ctx, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusTransactions200JSONResponse(*txs), nil
}

func (srv *StrictServerImpl) GetConsensusTransactionsTxHash(ctx context.Context, request apiTypes.GetConsensusTransactionsTxHashRequestObject) (apiTypes.GetConsensusTransactionsTxHashResponseObject, error) {
	tx, err := srv.client.Storage.Transaction(ctx, request.TxHash)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusTransactionsTxHash200JSONResponse(*tx), nil
}

func (srv *StrictServerImpl) GetConsensusValidators(ctx context.Context, request apiTypes.GetConsensusValidatorsRequestObject) (apiTypes.GetConsensusValidatorsResponseObject, error) {
	validators, err := srv.client.Storage.Validators(ctx, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusValidators200JSONResponse(*validators), nil
}

func (srv *StrictServerImpl) GetConsensusValidatorsEntityId(ctx context.Context, request apiTypes.GetConsensusValidatorsEntityIdRequestObject) (apiTypes.GetConsensusValidatorsEntityIdResponseObject, error) {
	validator, err := srv.client.Storage.Validator(ctx, request.EntityId)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusValidatorsEntityId200JSONResponse(*validator), nil
}

func (srv *StrictServerImpl) GetEmeraldBlocks(ctx context.Context, request apiTypes.GetEmeraldBlocksRequestObject) (apiTypes.GetEmeraldBlocksResponseObject, error) {
	return apiTypes.GetEmeraldBlocks200JSONResponse{}, nil
}

func (srv *StrictServerImpl) GetEmeraldTokens(ctx context.Context, request apiTypes.GetEmeraldTokensRequestObject) (apiTypes.GetEmeraldTokensResponseObject, error) {
	return apiTypes.GetEmeraldTokens200JSONResponse{}, nil
}

func (srv *StrictServerImpl) GetEmeraldTransactions(ctx context.Context, request apiTypes.GetEmeraldTransactionsRequestObject) (apiTypes.GetEmeraldTransactionsResponseObject, error) {
	return apiTypes.GetEmeraldTransactions200JSONResponse{}, nil
}
