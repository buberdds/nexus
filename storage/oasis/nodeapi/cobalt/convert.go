package cobalt

import (
	"github.com/oasisprotocol/oasis-core/go/common/quantity"

	// indexer-internal data types.
	"github.com/oasisprotocol/oasis-core/go/common"
	genesis "github.com/oasisprotocol/oasis-core/go/genesis/api"
	governance "github.com/oasisprotocol/oasis-core/go/governance/api"
	registry "github.com/oasisprotocol/oasis-core/go/registry/api"
	staking "github.com/oasisprotocol/oasis-core/go/staking/api"
	apiTypes "github.com/oasisprotocol/oasis-indexer/api/v1/types"
	"github.com/oasisprotocol/oasis-indexer/storage/oasis/nodeapi"

	// data types for Cobalt gRPC APIs.
	txResultsCobalt "github.com/oasisprotocol/oasis-indexer/coreapi/v21.1.1/consensus/api/transaction/results"
	genesisCobalt "github.com/oasisprotocol/oasis-indexer/coreapi/v21.1.1/genesis/api"
	governanceCobalt "github.com/oasisprotocol/oasis-indexer/coreapi/v21.1.1/governance/api"
	registryCobalt "github.com/oasisprotocol/oasis-indexer/coreapi/v21.1.1/registry/api"
	stakingCobalt "github.com/oasisprotocol/oasis-indexer/coreapi/v21.1.1/staking/api"
)

func convertProposal(p *governanceCobalt.Proposal) *governance.Proposal {
	results := make(map[governance.Vote]quantity.Quantity)
	for k, v := range p.Results {
		results[governance.Vote(k)] = v
	}

	return &governance.Proposal{
		ID:        p.ID,
		Submitter: p.Submitter,
		State:     governance.ProposalState(p.State),
		Deposit:   p.Deposit,
		Content: governance.ProposalContent{
			Upgrade:          (*governance.UpgradeProposal)(p.Content.Upgrade),
			CancelUpgrade:    (*governance.CancelUpgradeProposal)(p.Content.CancelUpgrade),
			ChangeParameters: nil, // not present in cobalt
		},
		CreatedAt:    p.CreatedAt,
		ClosesAt:     p.ClosesAt,
		Results:      results,
		InvalidVotes: 0,
	}
}

func convertAccount(a *stakingCobalt.Account) *staking.Account {
	rateSteps := make([]staking.CommissionRateStep, len(a.Escrow.CommissionSchedule.Rates))
	for i, r := range a.Escrow.CommissionSchedule.Rates {
		rateSteps[i] = staking.CommissionRateStep(r)
	}
	rateBoundSteps := make([]staking.CommissionRateBoundStep, len(a.Escrow.CommissionSchedule.Bounds))
	for i, r := range a.Escrow.CommissionSchedule.Bounds {
		rateBoundSteps[i] = staking.CommissionRateBoundStep(r)
	}
	return &staking.Account{
		General: staking.GeneralAccount(a.General),
		Escrow: staking.EscrowAccount{
			Active:    staking.SharePool(a.Escrow.Active),
			Debonding: staking.SharePool(a.Escrow.Debonding),
			CommissionSchedule: staking.CommissionSchedule{
				Rates:  rateSteps,
				Bounds: rateBoundSteps,
			},
		},
	}
}

func convertRuntime(r *registryCobalt.Runtime) *registry.Runtime {
	return &registry.Runtime{
		ID:          r.ID,
		EntityID:    r.EntityID,
		Kind:        registry.RuntimeKind(r.Kind),
		KeyManager:  r.KeyManager,
		TEEHardware: r.TEEHardware,
	}
}

// ConvertGenesis converts a genesis document from the Cobalt format to the
// indexer-internal (= current oasis-core) format.
// WARNING: This is a partial conversion, only the fields that are used by
// the indexer are filled in the output document.
func ConvertGenesis(d genesisCobalt.Document) *genesis.Document {
	proposals := make([]*governance.Proposal, len(d.Governance.Proposals))
	for i, p := range d.Governance.Proposals {
		proposals[i] = convertProposal(p)
	}

	voteEntries := make(map[uint64][]*governance.VoteEntry, len(d.Governance.VoteEntries))
	for k, v := range d.Governance.VoteEntries {
		voteEntries[k] = make([]*governance.VoteEntry, len(v))
		for i, ve := range v {
			voteEntries[k][i] = &governance.VoteEntry{
				Voter: ve.Voter,
				Vote:  governance.Vote(ve.Vote),
			}
		}
	}

	ledger := make(map[staking.Address]*staking.Account, len(d.Staking.Ledger))
	for k, v := range d.Staking.Ledger {
		ledger[k] = convertAccount(v)
	}

	delegations := make(map[staking.Address]map[staking.Address]*staking.Delegation, len(d.Staking.Delegations))
	for k, v := range d.Staking.Delegations {
		delegations[k] = make(map[staking.Address]*staking.Delegation, len(v))
		for k2, v2 := range v {
			delegations[k][k2] = &staking.Delegation{
				Shares: v2.Shares,
			}
		}
	}

	debondingDelegations := make(map[staking.Address]map[staking.Address][]*staking.DebondingDelegation, len(d.Staking.DebondingDelegations))
	for k, v := range d.Staking.DebondingDelegations {
		debondingDelegations[k] = make(map[staking.Address][]*staking.DebondingDelegation, len(v))
		for k2, v2 := range v {
			debondingDelegations[k][k2] = make([]*staking.DebondingDelegation, len(v2))
			for i, v3 := range v2 {
				debondingDelegations[k][k2][i] = &staking.DebondingDelegation{
					Shares: v3.Shares,
				}
			}
		}
	}

	runtimes := make([]*registry.Runtime, len(d.Registry.Runtimes))
	for i, r := range d.Registry.Runtimes {
		runtimes[i] = convertRuntime(r)
	}

	return &genesis.Document{
		Height:  d.Height,
		Time:    d.Time,
		ChainID: d.ChainID,
		Governance: governance.Genesis{
			Proposals:   proposals,
			VoteEntries: voteEntries,
		},
		Registry: registry.Genesis{
			Entities:          d.Registry.Entities,
			Runtimes:          []*registry.Runtime{},
			SuspendedRuntimes: []*registry.Runtime{},
			Nodes:             d.Registry.Nodes,
		},
		Staking: staking.Genesis{
			CommonPool:           d.Staking.CommonPool,
			LastBlockFees:        d.Staking.LastBlockFees,
			GovernanceDeposits:   d.Staking.GovernanceDeposits,
			Ledger:               ledger,
			Delegations:          delegations,
			DebondingDelegations: debondingDelegations,
		},
	}
}

func convertTxResult(r txResultsCobalt.Result) nodeapi.Result {
	events := make([]nodeapi.Event, len(r.Events))
	for i, e := range r.Events {
		switch {
		case e.Staking != nil:
			switch {
			case e.Staking.Transfer != nil:
				events[i] = nodeapi.Event{
					StakingTransfer: (*nodeapi.TransferEvent)(e.Staking.Transfer),
					Raw:             e.Staking.Transfer,
					Type:            apiTypes.ConsensusEventTypeStakingTransfer,
				}
			case e.Staking.Burn != nil:
				events[i] = nodeapi.Event{
					StakingBurn: (*nodeapi.BurnEvent)(e.Staking.Burn),
					Raw:         e.Staking.Burn,
					Type:        apiTypes.ConsensusEventTypeStakingBurn,
				}
			case e.Staking.Escrow != nil:
				switch {
				case e.Staking.Escrow.Add != nil:
					events[i] = nodeapi.Event{
						StakingAddEscrow: &nodeapi.AddEscrowEvent{
							Owner:     e.Staking.Escrow.Add.Owner,
							Escrow:    e.Staking.Escrow.Add.Escrow,
							Amount:    e.Staking.Escrow.Add.Amount,
							NewShares: quantity.Quantity{}, // NOTE: not available in the Cobalt API
						},
						Raw:  e.Staking.Escrow.Add,
						Type: apiTypes.ConsensusEventTypeStakingEscrowAdd,
					}
				case e.Staking.Escrow.Take != nil:
					events[i] = nodeapi.Event{
						StakingTakeEscrow: (*nodeapi.TakeEscrowEvent)(e.Staking.Escrow.Take),
						Raw:               e.Staking.Escrow.Take,
						Type:              apiTypes.ConsensusEventTypeStakingEscrowTake,
					}
				case e.Staking.Escrow.Reclaim != nil:
					events[i] = nodeapi.Event{
						StakingReclaimEscrow: &nodeapi.ReclaimEscrowEvent{
							Owner:  e.Staking.Escrow.Reclaim.Owner,
							Escrow: e.Staking.Escrow.Reclaim.Escrow,
							Amount: e.Staking.Escrow.Reclaim.Amount,
							Shares: quantity.Quantity{}, // NOTE: not available in the Cobalt API
						},
						Raw:  e.Staking.Escrow.Reclaim,
						Type: apiTypes.ConsensusEventTypeStakingEscrowReclaim,
					}
					// NOTE: There is no Staking.Escrow.DebondingStart event in Cobalt.
				}
			case e.Staking.AllowanceChange != nil:
				events[i] = nodeapi.Event{
					StakingAllowanceChange: (*nodeapi.AllowanceChangeEvent)(e.Staking.AllowanceChange),
					Raw:                    e.Staking.AllowanceChange,
					Type:                   apiTypes.ConsensusEventTypeStakingAllowanceChange,
				}
			}
			events[i].Height = e.Staking.Height
			events[i].TxHash = e.Staking.TxHash
			// End Staking.
		case e.Registry != nil:
			switch {
			case e.Registry.RuntimeEvent != nil && e.Registry.RuntimeEvent.Runtime != nil:
				events[i] = nodeapi.Event{
					RegistryRuntime: &nodeapi.RuntimeEvent{
						ID:       e.Registry.RuntimeEvent.Runtime.ID,
						EntityID: e.Registry.RuntimeEvent.Runtime.EntityID,
					},
					Raw:  e.Registry.RuntimeEvent,
					Type: apiTypes.ConsensusEventTypeRegistryRuntime,
				}
			case e.Registry.EntityEvent != nil:
				events[i] = nodeapi.Event{
					RegistryEntity: (*nodeapi.EntityEvent)(e.Registry.EntityEvent),
					Raw:            e.Registry.EntityEvent,
					Type:           apiTypes.ConsensusEventTypeRegistryEntity,
				}
			case e.Registry.NodeEvent != nil:
				runtimeIDs := make([]common.Namespace, len(e.Registry.NodeEvent.Node.Runtimes))
				for i, r := range e.Registry.NodeEvent.Node.Runtimes {
					runtimeIDs[i] = r.ID
				}
				events[i] = nodeapi.Event{
					RegistryNode: &nodeapi.NodeEvent{
						NodeID:         e.Registry.NodeEvent.Node.EntityID,
						EntityID:       e.Registry.NodeEvent.Node.EntityID,
						RuntimeIDs:     runtimeIDs,
						IsRegistration: e.Registry.NodeEvent.IsRegistration,
					},
					Raw:  e.Registry.NodeEvent,
					Type: apiTypes.ConsensusEventTypeRegistryNode,
				}
			case e.Registry.NodeUnfrozenEvent != nil:
				events[i] = nodeapi.Event{
					RegistryNodeUnfrozen: (*nodeapi.NodeUnfrozenEvent)(e.Registry.NodeUnfrozenEvent),
					Raw:                  e.Registry.NodeUnfrozenEvent,
					Type:                 apiTypes.ConsensusEventTypeRegistryNodeUnfrozen,
				}
			}
			events[i].Height = e.Registry.Height
			events[i].TxHash = e.Registry.TxHash
			// End Registry.
		case e.RootHash != nil:
			switch {
			case e.RootHash.ExecutorCommitted != nil:
				events[i] = nodeapi.Event{
					RoothashExecutorCommitted: &nodeapi.ExecutorCommittedEvent{
						NodeID: nil, // Not available in Cobalt.
					},
					Raw:  e.RootHash.ExecutorCommitted,
					Type: apiTypes.ConsensusEventTypeRoothashExecutorCommitted,
				}
			case e.RootHash.ExecutionDiscrepancyDetected != nil:
				events[i] = nodeapi.Event{
					Raw:  e.RootHash.ExecutionDiscrepancyDetected,
					Type: apiTypes.ConsensusEventTypeRoothashExecutionDiscrepancy,
				}
			case e.RootHash.Finalized != nil:
				events[i] = nodeapi.Event{
					Raw:  e.RootHash.Finalized,
					Type: apiTypes.ConsensusEventTypeRoothashFinalized,
				}
			}
			events[i].Height = e.RootHash.Height
			events[i].TxHash = e.RootHash.TxHash
			// End RootHash.
		case e.Governance != nil:
			switch {
			case e.Governance.ProposalSubmitted != nil:
				events[i] = nodeapi.Event{
					GovernanceProposalSubmitted: &nodeapi.ProposalSubmittedEvent{
						Submitter: e.Governance.ProposalSubmitted.Submitter,
					},
					Raw:  e.Governance.ProposalSubmitted,
					Type: apiTypes.ConsensusEventTypeGovernanceProposalSubmitted,
				}
			case e.Governance.ProposalExecuted != nil:
				events[i] = nodeapi.Event{
					Raw:  e.Governance.ProposalExecuted,
					Type: apiTypes.ConsensusEventTypeGovernanceProposalExecuted,
				}
			case e.Governance.ProposalFinalized != nil:
				events[i] = nodeapi.Event{
					Raw:  e.Governance.ProposalFinalized,
					Type: apiTypes.ConsensusEventTypeGovernanceProposalFinalized,
				}
			case e.Governance.Vote != nil:
				events[i] = nodeapi.Event{
					GovernanceVote: &nodeapi.VoteEvent{
						Submitter: e.Governance.Vote.Submitter,
					},
					Raw:  e.Governance.Vote,
					Type: apiTypes.ConsensusEventTypeGovernanceVote,
				}
			}
			events[i].Height = e.Governance.Height
			events[i].TxHash = e.Governance.TxHash
			// End Governance.
		}
	}

	return nodeapi.Result{
		Error:  consensusTxResults.Error(r.Error),
		Events: events,
	}
}
