package types

func (c ConsensusEventType) IsValid() bool {
	switch c {
	case ConsensusEventTypeStakingTransfer,
		ConsensusEventTypeStakingBurn,
		ConsensusEventTypeStakingEscrowAdd,
		ConsensusEventTypeStakingEscrowTake,
		ConsensusEventTypeStakingEscrowDebondingStart,
		ConsensusEventTypeStakingEscrowReclaim,
		ConsensusEventTypeStakingAllowanceChange,
		ConsensusEventTypeRegistryRuntime,
		ConsensusEventTypeRegistryEntity,
		ConsensusEventTypeRegistryNode,
		ConsensusEventTypeRegistryNodeUnfrozen,
		ConsensusEventTypeRoothashExecutorCommitted,
		ConsensusEventTypeRoothashExecutionDiscrepancy,
		ConsensusEventTypeRoothashFinalized,
		ConsensusEventTypeGovernanceProposalSubmitted,
		ConsensusEventTypeGovernanceProposalExecuted,
		ConsensusEventTypeGovernanceProposalFinalized,
		ConsensusEventTypeGovernanceVote:
		return true
	default:
		return false
	}
}