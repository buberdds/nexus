// Package consensus implements an analyzer for the consensus layer.
package consensus

import (
	"context"

	coreCommon "github.com/oasisprotocol/oasis-core/go/common"
	sdkConfig "github.com/oasisprotocol/oasis-sdk/client-sdk/go/config"

	beacon "github.com/oasisprotocol/nexus/coreapi/v22.2.11/beacon/api"
	consensus "github.com/oasisprotocol/nexus/coreapi/v22.2.11/consensus/api"
	"github.com/oasisprotocol/nexus/storage/oasis/nodeapi"
)

// fetchAllData returns all relevant data related to the given height.
func fetchAllData(ctx context.Context, cc nodeapi.ConsensusApiLite, network sdkConfig.Network, height int64, fastSync bool) (*allData, error) {
	blockData, err := fetchBlockData(ctx, cc, height)
	if err != nil {
		return nil, err
	}

	beaconData, err := fetchBeaconData(ctx, cc, height)
	if err != nil {
		return nil, err
	}

	registryData, err := fetchRegistryData(ctx, cc, height)
	if err != nil {
		return nil, err
	}

	stakingData, err := fetchStakingData(ctx, cc, height)
	if err != nil {
		return nil, err
	}

	var schedulerData *schedulerData
	// Scheduler data is not needed during fast sync. It contains no events,
	// only a complete snapshot validators/committees. Since we don't store historical data,
	// any single snapshot during slow-sync is sufficient to reconstruct the state.
	if !fastSync {
		schedulerData, err = fetchSchedulerData(ctx, cc, network, height)
		if err != nil {
			return nil, err
		}
	}

	governanceData, err := fetchGovernanceData(ctx, cc, height)
	if err != nil {
		return nil, err
	}

	rootHashData, err := fetchRootHashData(ctx, cc, height)
	if err != nil {
		return nil, err
	}

	data := allData{
		BlockData:      blockData,
		BeaconData:     beaconData,
		RegistryData:   registryData,
		RootHashData:   rootHashData,
		StakingData:    stakingData,
		SchedulerData:  schedulerData,
		GovernanceData: governanceData,
	}
	return &data, nil
}

// fetchBlockData retrieves data about a consensus block at the provided block height.
func fetchBlockData(ctx context.Context, cc nodeapi.ConsensusApiLite, height int64) (*consensusBlockData, error) {
	block, err := cc.GetBlock(ctx, height)
	if err != nil {
		return nil, err
	}

	epoch, err := cc.GetEpoch(ctx, height)
	if err != nil {
		return nil, err
	}

	transactionsWithResults, err := cc.GetTransactionsWithResults(ctx, height)
	if err != nil {
		return nil, err
	}

	return &consensusBlockData{
		Height:                  height,
		BlockHeader:             block,
		Epoch:                   epoch,
		TransactionsWithResults: transactionsWithResults,
	}, nil
}

// fetchBeaconData retrieves the beacon for the provided block height.
func fetchBeaconData(ctx context.Context, cc nodeapi.ConsensusApiLite, height int64) (*beaconData, error) {
	// NOTE: The random beacon endpoint is in flux.
	// GetBeacon() consistently errors out (at least for heights soon after Damask genesis) with "beacon: random beacon not available".
	// beacon, err := cc.client.Beacon().GetBeacon(ctx, height)
	// if err != nil {
	// 	return nil, err
	// }

	epoch, err := cc.GetEpoch(ctx, height)
	if err != nil {
		return nil, err
	}

	return &beaconData{
		Height: height,
		Epoch:  epoch,
		Beacon: nil,
	}, nil
}

// fetchRegistryData retrieves registry events at the provided block height.
func fetchRegistryData(ctx context.Context, cc nodeapi.ConsensusApiLite, height int64) (*registryData, error) {
	events, err := cc.RegistryEvents(ctx, height)
	if err != nil {
		return nil, err
	}

	var runtimeStartedEvents []nodeapi.RuntimeStartedEvent
	var runtimeSuspendedEvents []nodeapi.RuntimeSuspendedEvent
	var entityEvents []nodeapi.EntityEvent
	var nodeEvents []nodeapi.NodeEvent
	var nodeUnfrozenEvents []nodeapi.NodeUnfrozenEvent

	for _, event := range events {
		switch e := event; {
		case e.RegistryRuntimeStarted != nil:
			runtimeStartedEvents = append(runtimeStartedEvents, *e.RegistryRuntimeStarted)
		case e.RegistryRuntimeSuspended != nil:
			runtimeSuspendedEvents = append(runtimeSuspendedEvents, *e.RegistryRuntimeSuspended)
		case e.RegistryEntity != nil:
			entityEvents = append(entityEvents, *e.RegistryEntity)
		case e.RegistryNode != nil:
			nodeEvents = append(nodeEvents, *e.RegistryNode)
		case e.RegistryNodeUnfrozen != nil:
			nodeUnfrozenEvents = append(nodeUnfrozenEvents, *e.RegistryNodeUnfrozen)
		}
	}

	return &registryData{
		Height:                 height,
		Events:                 events,
		RuntimeStartedEvents:   runtimeStartedEvents,
		RuntimeSuspendedEvents: runtimeSuspendedEvents,
		EntityEvents:           entityEvents,
		NodeEvents:             nodeEvents,
		NodeUnfrozenEvents:     nodeUnfrozenEvents,
	}, nil
}

// fetchStakingData retrieves staking events at the provided block height.
func fetchStakingData(ctx context.Context, cc nodeapi.ConsensusApiLite, height int64) (*stakingData, error) {
	events, err := cc.StakingEvents(ctx, height)
	if err != nil {
		return nil, err
	}

	epoch, err := cc.GetEpoch(ctx, height)
	if err != nil {
		return nil, err
	}

	var transfers []nodeapi.TransferEvent
	var burns []nodeapi.BurnEvent
	var addEscrows []nodeapi.AddEscrowEvent
	var reclaimEscrows []nodeapi.ReclaimEscrowEvent
	var debondingStartEscrows []nodeapi.DebondingStartEscrowEvent
	var takeEscrows []nodeapi.TakeEscrowEvent
	var allowanceChanges []nodeapi.AllowanceChangeEvent

	for _, event := range events {
		switch e := event; {
		case e.StakingTransfer != nil:
			transfers = append(transfers, *event.StakingTransfer)
		case e.StakingBurn != nil:
			burns = append(burns, *event.StakingBurn)
		case e.StakingAddEscrow != nil:
			addEscrows = append(addEscrows, *event.StakingAddEscrow)
		case e.StakingReclaimEscrow != nil:
			reclaimEscrows = append(reclaimEscrows, *event.StakingReclaimEscrow)
		case e.StakingDebondingStart != nil:
			debondingStartEscrows = append(debondingStartEscrows, *event.StakingDebondingStart)
		case e.StakingTakeEscrow != nil:
			takeEscrows = append(takeEscrows, *event.StakingTakeEscrow)
		case e.StakingAllowanceChange != nil:
			allowanceChanges = append(allowanceChanges, *event.StakingAllowanceChange)
		}
	}

	return &stakingData{
		Height:                height,
		Epoch:                 epoch,
		Events:                events,
		Transfers:             transfers,
		Burns:                 burns,
		AddEscrows:            addEscrows,
		ReclaimEscrows:        reclaimEscrows,
		DebondingStartEscrows: debondingStartEscrows,
		TakeEscrows:           takeEscrows,
		AllowanceChanges:      allowanceChanges,
	}, nil
}

// fetchSchedulerData retrieves validators and runtime committees at the provided block height.
func fetchSchedulerData(ctx context.Context, cc nodeapi.ConsensusApiLite, network sdkConfig.Network, height int64) (*schedulerData, error) {
	validators, err := cc.GetValidators(ctx, height)
	if err != nil {
		return nil, err
	}

	committees := make(map[coreCommon.Namespace][]nodeapi.Committee, len(network.ParaTimes.All))

	for name := range network.ParaTimes.All {
		var runtimeID coreCommon.Namespace
		if err := runtimeID.UnmarshalHex(network.ParaTimes.All[name].ID); err != nil {
			return nil, err
		}

		consensusCommittees, err := cc.GetCommittees(ctx, height, runtimeID)
		if err != nil {
			return nil, err
		}
		committees[runtimeID] = consensusCommittees
	}

	return &schedulerData{
		Height:     height,
		Validators: validators,
		Committees: committees,
	}, nil
}

// fetchGovernanceData retrieves governance events at the provided block height.
func fetchGovernanceData(ctx context.Context, cc nodeapi.ConsensusApiLite, height int64) (*governanceData, error) {
	events, err := cc.GovernanceEvents(ctx, height)
	if err != nil {
		return nil, err
	}

	var submissions []nodeapi.Proposal
	var executions []nodeapi.ProposalExecutedEvent
	var finalizations []nodeapi.Proposal
	var votes []nodeapi.VoteEvent

	for _, event := range events {
		switch {
		case event.GovernanceProposalSubmitted != nil:
			proposal, err := cc.GetProposal(ctx, height, event.GovernanceProposalSubmitted.ID)
			if err != nil {
				return nil, err
			}
			submissions = append(submissions, *proposal)
		case event.GovernanceProposalExecuted != nil:
			executions = append(executions, *event.GovernanceProposalExecuted)
		case event.GovernanceProposalFinalized != nil:
			proposal, err := cc.GetProposal(ctx, height, event.GovernanceProposalFinalized.ID)
			if err != nil {
				return nil, err
			}
			finalizations = append(finalizations, *proposal)
		case event.GovernanceVote != nil:
			votes = append(votes, *event.GovernanceVote)
		}
	}
	return &governanceData{
		Height:                height,
		Events:                events,
		ProposalSubmissions:   submissions,
		ProposalExecutions:    executions,
		ProposalFinalizations: finalizations,
		Votes:                 votes,
	}, nil
}

// fetchRootHashData retrieves roothash events at the provided block height.
func fetchRootHashData(ctx context.Context, cc nodeapi.ConsensusApiLite, height int64) (*rootHashData, error) {
	events, err := cc.RoothashEvents(ctx, height)
	if err != nil {
		return nil, err
	}

	return &rootHashData{
		Height: height,
		Events: events,
	}, nil
}

type allData struct {
	BlockData      *consensusBlockData
	BeaconData     *beaconData
	RegistryData   *registryData
	RootHashData   *rootHashData
	StakingData    *stakingData
	SchedulerData  *schedulerData
	GovernanceData *governanceData
}

// consensusBlockData represents data for a consensus block at a given height.
type consensusBlockData struct {
	Height int64

	BlockHeader             *consensus.Block
	Epoch                   beacon.EpochTime
	TransactionsWithResults []nodeapi.TransactionWithResults
}

// beaconData represents data for the random beacon at a given height.
type beaconData struct {
	Height int64

	Epoch  beacon.EpochTime
	Beacon []byte
}

// registryData represents data for the node registry at a given height.
type registryData struct {
	Height int64

	Events []nodeapi.Event

	// Below: Same events as in `Events` but grouped by type.
	RuntimeStartedEvents   []nodeapi.RuntimeStartedEvent
	RuntimeSuspendedEvents []nodeapi.RuntimeSuspendedEvent
	EntityEvents           []nodeapi.EntityEvent
	NodeEvents             []nodeapi.NodeEvent
	NodeUnfrozenEvents     []nodeapi.NodeUnfrozenEvent
}

// stakingData represents data for accounts at a given height.
type stakingData struct {
	Height int64
	Epoch  beacon.EpochTime

	Events []nodeapi.Event

	// Below: Same events as in `Events` but grouped by type.
	Transfers             []nodeapi.TransferEvent
	Burns                 []nodeapi.BurnEvent
	AddEscrows            []nodeapi.AddEscrowEvent
	TakeEscrows           []nodeapi.TakeEscrowEvent
	ReclaimEscrows        []nodeapi.ReclaimEscrowEvent
	DebondingStartEscrows []nodeapi.DebondingStartEscrowEvent
	AllowanceChanges      []nodeapi.AllowanceChangeEvent
}

// rootHashData represents data for runtime processing at a given height.
type rootHashData struct {
	Height int64

	Events []nodeapi.Event
}

// schedulerData represents data for elected committees and validators at a given height.
type schedulerData struct {
	Height int64

	Validators []nodeapi.Validator
	Committees map[coreCommon.Namespace][]nodeapi.Committee
}

// governanceData represents governance data for proposals at a given height.
type governanceData struct {
	Height int64

	Events []nodeapi.Event

	// Below: Same events as in `Events` but grouped by type.
	ProposalSubmissions   []nodeapi.Proposal
	ProposalExecutions    []nodeapi.ProposalExecutedEvent
	ProposalFinalizations []nodeapi.Proposal
	Votes                 []nodeapi.VoteEvent
}