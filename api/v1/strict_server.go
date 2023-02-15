package v1

import (
	"context"
	"fmt"

	apiTypes "github.com/oasisprotocol/oasis-indexer/api/v1/types"
	"github.com/oasisprotocol/oasis-indexer/log"
	"github.com/oasisprotocol/oasis-indexer/storage/client"
)

// StrictServerImpl implements the oapi-codegen StrictServerInterface interface,
// which exposes our API endpoints as functions with strongly-typed params.
// This struct is a thin layer over the DB-querying client; it mostly just forwards
// the inputs and outputs, but sometimes it additionally processes them.
type StrictServerImpl struct {
	dbClient client.StorageClient
	logger   log.Logger
}

var _ apiTypes.StrictServerInterface = (*StrictServerImpl)(nil)

func NewStrictServerImpl(client client.StorageClient, logger log.Logger) *StrictServerImpl {
	return &StrictServerImpl{
		dbClient: client,
		logger:   logger,
	}
}

// Stubs of these functions were derived from the autogenerated interface with:
//	sed -n '/type StrictServerInterface interface/,/^\}/p' api/v1/types/server.gen.go | grep -v // | head -n-1 | tail -n+2 | sed -E 's/^\s+(\w+)(.*)/func (srv *StrictServerImpl) \1\2 { return apiTypes.\1200JSONResponse{}, nil }\n/g; s/[a-zA-Z]+(Params|RequestObject|ResponseObject)/apiTypes.\0/g;'

func (srv *StrictServerImpl) GetStatus(ctx context.Context, request apiTypes.GetStatusRequestObject) (apiTypes.GetStatusResponseObject, error) {
	status, err := srv.dbClient.Status(ctx)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetStatus200JSONResponse(*status), nil
}

func (srv *StrictServerImpl) GetConsensusAccounts(ctx context.Context, request apiTypes.GetConsensusAccountsRequestObject) (apiTypes.GetConsensusAccountsResponseObject, error) {
	accounts, err := srv.dbClient.Accounts(ctx, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusAccounts200JSONResponse(*accounts), nil
}

func (srv *StrictServerImpl) GetConsensusAccountsAddress(ctx context.Context, request apiTypes.GetConsensusAccountsAddressRequestObject) (apiTypes.GetConsensusAccountsAddressResponseObject, error) {
	account, err := srv.dbClient.Account(ctx, request.Address)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusAccountsAddress200JSONResponse(*account), nil
}

func (srv *StrictServerImpl) GetConsensusAccountsAddressDebondingDelegations(ctx context.Context, request apiTypes.GetConsensusAccountsAddressDebondingDelegationsRequestObject) (apiTypes.GetConsensusAccountsAddressDebondingDelegationsResponseObject, error) {
	delegations, err := srv.dbClient.DebondingDelegations(ctx, request.Address, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusAccountsAddressDebondingDelegations200JSONResponse(*delegations), nil
}

func (srv *StrictServerImpl) GetConsensusAccountsAddressDelegations(ctx context.Context, request apiTypes.GetConsensusAccountsAddressDelegationsRequestObject) (apiTypes.GetConsensusAccountsAddressDelegationsResponseObject, error) {
	delegations, err := srv.dbClient.Delegations(ctx, request.Address, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusAccountsAddressDelegations200JSONResponse(*delegations), nil
}

func (srv *StrictServerImpl) GetConsensusBlocks(ctx context.Context, request apiTypes.GetConsensusBlocksRequestObject) (apiTypes.GetConsensusBlocksResponseObject, error) {
	blocks, err := srv.dbClient.Blocks(ctx, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusBlocks200JSONResponse(*blocks), nil
}

func (srv *StrictServerImpl) GetConsensusBlocksHeight(ctx context.Context, request apiTypes.GetConsensusBlocksHeightRequestObject) (apiTypes.GetConsensusBlocksHeightResponseObject, error) {
	block, err := srv.dbClient.Block(ctx, request.Height)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusBlocksHeight200JSONResponse(*block), nil
}

func (srv *StrictServerImpl) GetConsensusEntities(ctx context.Context, request apiTypes.GetConsensusEntitiesRequestObject) (apiTypes.GetConsensusEntitiesResponseObject, error) {
	entities, err := srv.dbClient.Entities(ctx, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusEntities200JSONResponse(*entities), nil
}

func (srv *StrictServerImpl) GetConsensusEntitiesEntityId(ctx context.Context, request apiTypes.GetConsensusEntitiesEntityIdRequestObject) (apiTypes.GetConsensusEntitiesEntityIdResponseObject, error) {
	entity, err := srv.dbClient.Entity(ctx, request.EntityId)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusEntitiesEntityId200JSONResponse(*entity), nil
}

func (srv *StrictServerImpl) GetConsensusEntitiesEntityIdNodes(ctx context.Context, request apiTypes.GetConsensusEntitiesEntityIdNodesRequestObject) (apiTypes.GetConsensusEntitiesEntityIdNodesResponseObject, error) {
	nodes, err := srv.dbClient.EntityNodes(ctx, request.EntityId, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusEntitiesEntityIdNodes200JSONResponse(*nodes), nil
}

func (srv *StrictServerImpl) GetConsensusEntitiesEntityIdNodesNodeId(ctx context.Context, request apiTypes.GetConsensusEntitiesEntityIdNodesNodeIdRequestObject) (apiTypes.GetConsensusEntitiesEntityIdNodesNodeIdResponseObject, error) {
	node, err := srv.dbClient.EntityNode(ctx, request.EntityId, request.NodeId)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusEntitiesEntityIdNodesNodeId200JSONResponse(*node), nil
}

func (srv *StrictServerImpl) GetConsensusEpochs(ctx context.Context, request apiTypes.GetConsensusEpochsRequestObject) (apiTypes.GetConsensusEpochsResponseObject, error) {
	epochs, err := srv.dbClient.Epochs(ctx, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusEpochs200JSONResponse(*epochs), nil
}

func (srv *StrictServerImpl) GetConsensusEpochsEpoch(ctx context.Context, request apiTypes.GetConsensusEpochsEpochRequestObject) (apiTypes.GetConsensusEpochsEpochResponseObject, error) {
	epoch, err := srv.dbClient.Epoch(ctx, request.Epoch)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusEpochsEpoch200JSONResponse(*epoch), nil
}

func (srv *StrictServerImpl) GetConsensusEvents(ctx context.Context, request apiTypes.GetConsensusEventsRequestObject) (apiTypes.GetConsensusEventsResponseObject, error) {
	// Additional param validation.
	if request.Params.Type != nil && !request.Params.Type.IsValid() {
		return nil, &apiTypes.InvalidParamFormatError{ParamName: "type", Err: fmt.Errorf("not a valid enum value: %s", *request.Params.Type)}
	}
	if request.Params.TxIndex != nil && request.Params.Block == nil {
		return nil, &apiTypes.InvalidParamFormatError{ParamName: "block", Err: fmt.Errorf("must be specified when tx_index is specified")}
	}

	events, err := srv.dbClient.Events(ctx, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusEvents200JSONResponse(*events), nil
}

func (srv *StrictServerImpl) GetConsensusProposals(ctx context.Context, request apiTypes.GetConsensusProposalsRequestObject) (apiTypes.GetConsensusProposalsResponseObject, error) {
	proposals, err := srv.dbClient.Proposals(ctx, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusProposals200JSONResponse(*proposals), nil
}

func (srv *StrictServerImpl) GetConsensusProposalsProposalId(ctx context.Context, request apiTypes.GetConsensusProposalsProposalIdRequestObject) (apiTypes.GetConsensusProposalsProposalIdResponseObject, error) {
	proposal, err := srv.dbClient.Proposal(ctx, request.ProposalId)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusProposalsProposalId200JSONResponse(*proposal), nil
}

func (srv *StrictServerImpl) GetConsensusProposalsProposalIdVotes(ctx context.Context, request apiTypes.GetConsensusProposalsProposalIdVotesRequestObject) (apiTypes.GetConsensusProposalsProposalIdVotesResponseObject, error) {
	votes, err := srv.dbClient.ProposalVotes(ctx, request.ProposalId, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusProposalsProposalIdVotes200JSONResponse(*votes), nil
}

func (srv *StrictServerImpl) GetLayerStatsTxVolume(ctx context.Context, request apiTypes.GetLayerStatsTxVolumeRequestObject) (apiTypes.GetLayerStatsTxVolumeResponseObject, error) {
	// Additional param validation.
	if !request.Layer.IsValid() {
		return nil, &apiTypes.InvalidParamFormatError{ParamName: "layer", Err: fmt.Errorf("not a valid enum value: %s", request.Layer)}
	}

	volumeList, err := srv.dbClient.TxVolumes(ctx, request.Layer, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetLayerStatsTxVolume200JSONResponse(*volumeList), nil
}

func (srv *StrictServerImpl) GetConsensusTransactions(ctx context.Context, request apiTypes.GetConsensusTransactionsRequestObject) (apiTypes.GetConsensusTransactionsResponseObject, error) {
	txs, err := srv.dbClient.Transactions(ctx, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusTransactions200JSONResponse(*txs), nil
}

func (srv *StrictServerImpl) GetConsensusTransactionsTxHash(ctx context.Context, request apiTypes.GetConsensusTransactionsTxHashRequestObject) (apiTypes.GetConsensusTransactionsTxHashResponseObject, error) {
	tx, err := srv.dbClient.Transaction(ctx, request.TxHash)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusTransactionsTxHash200JSONResponse(*tx), nil
}

func (srv *StrictServerImpl) GetConsensusValidators(ctx context.Context, request apiTypes.GetConsensusValidatorsRequestObject) (apiTypes.GetConsensusValidatorsResponseObject, error) {
	validators, err := srv.dbClient.Validators(ctx, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusValidators200JSONResponse(*validators), nil
}

func (srv *StrictServerImpl) GetConsensusValidatorsEntityId(ctx context.Context, request apiTypes.GetConsensusValidatorsEntityIdRequestObject) (apiTypes.GetConsensusValidatorsEntityIdResponseObject, error) {
	validator, err := srv.dbClient.Validator(ctx, request.EntityId)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetConsensusValidatorsEntityId200JSONResponse(*validator), nil
}

func (srv *StrictServerImpl) GetRuntimeBlocks(ctx context.Context, request apiTypes.GetRuntimeBlocksRequestObject) (apiTypes.GetRuntimeBlocksResponseObject, error) {
	blocks, err := srv.dbClient.RuntimeBlocks(ctx, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetRuntimeBlocks200JSONResponse(*blocks), nil
}

func (srv *StrictServerImpl) GetRuntimeEvmTokens(ctx context.Context, request apiTypes.GetRuntimeEvmTokensRequestObject) (apiTypes.GetRuntimeEvmTokensResponseObject, error) {
	tokens, err := srv.dbClient.RuntimeTokens(ctx, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetRuntimeEvmTokens200JSONResponse(*tokens), nil
}

func (srv *StrictServerImpl) GetRuntimeTransactions(ctx context.Context, request apiTypes.GetRuntimeTransactionsRequestObject) (apiTypes.GetRuntimeTransactionsResponseObject, error) {
	storageTransactions, err := srv.dbClient.RuntimeTransactions(ctx, request.Params, nil)
	if err != nil {
		return nil, err
	}

	// Perform additional tx body parsing on the fly; DB stores only partially-parsed txs.
	apiTransactions := apiTypes.RuntimeTransactionList{
		Transactions: []apiTypes.RuntimeTransaction{},
	}
	for _, storageTransaction := range storageTransactions.Transactions {
		apiTransaction, err2 := renderRuntimeTransaction(storageTransaction)
		if err2 != nil {
			return nil, fmt.Errorf("round %d tx %d: %w", storageTransaction.Round, storageTransaction.Index, err2)
		}
		apiTransactions.Transactions = append(apiTransactions.Transactions, apiTransaction)
	}

	return apiTypes.GetRuntimeTransactions200JSONResponse(apiTransactions), nil
}

func (srv *StrictServerImpl) GetRuntimeTransactionsTxHash(ctx context.Context, request apiTypes.GetRuntimeTransactionsTxHashRequestObject) (apiTypes.GetRuntimeTransactionsTxHashResponseObject, error) {
	storageTransactions, err := srv.dbClient.RuntimeTransactions(ctx, apiTypes.GetRuntimeTransactionsParams{}, &request.TxHash)
	if err != nil {
		return nil, err
	}

	if len(storageTransactions.Transactions) == 0 {
		return apiTypes.GetRuntimeTransactionsTxHash404JSONResponse{}, nil
	}

	// Perform additional tx body parsing on the fly; DB stores only partially-parsed txs.
	var apiTransactions apiTypes.RuntimeTransactionList
	for _, storageTransaction := range storageTransactions.Transactions {
		apiTransaction, err2 := renderRuntimeTransaction(storageTransaction)
		if err2 != nil {
			return nil, fmt.Errorf("round %d tx %d: %w", storageTransaction.Round, storageTransaction.Index, err2)
		}
		apiTransactions.Transactions = append(apiTransactions.Transactions, apiTransaction)
	}

	return apiTypes.GetRuntimeTransactionsTxHash200JSONResponse(apiTransactions), nil
}

func (srv *StrictServerImpl) GetRuntimeEvents(ctx context.Context, request apiTypes.GetRuntimeEventsRequestObject) (apiTypes.GetRuntimeEventsResponseObject, error) {
	events, err := srv.dbClient.RuntimeEvents(ctx, request.Params)
	if err != nil {
		return nil, err
	}
	return apiTypes.GetRuntimeEvents200JSONResponse(*events), nil
}
