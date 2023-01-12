// Package types provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package types

import (
	"time"

	common "github.com/oasisprotocol/oasis-indexer/common"
)

// Defines values for AddressDerivationContext.
const (
	OasisCoreaddressStaking            AddressDerivationContext = "oasis-core/address: staking"
	OasisRuntimeSdkaddressModule       AddressDerivationContext = "oasis-runtime-sdk/address: module"
	OasisRuntimeSdkaddressMultisig     AddressDerivationContext = "oasis-runtime-sdk/address: multisig"
	OasisRuntimeSdkaddressRuntime      AddressDerivationContext = "oasis-runtime-sdk/address: runtime"
	OasisRuntimeSdkaddressSecp256k1eth AddressDerivationContext = "oasis-runtime-sdk/address: secp256k1eth"
	OasisRuntimeSdkaddressSr25519      AddressDerivationContext = "oasis-runtime-sdk/address: sr25519"
)

// Defines values for ConsensusEventType.
const (
	ConsensusEventTypeGovernanceProposalExecuted   ConsensusEventType = "governance.proposal_executed"
	ConsensusEventTypeGovernanceProposalFinalized  ConsensusEventType = "governance.proposal_finalized"
	ConsensusEventTypeGovernanceProposalSubmitted  ConsensusEventType = "governance.proposal_submitted"
	ConsensusEventTypeGovernanceVote               ConsensusEventType = "governance.vote"
	ConsensusEventTypeRegistryEntity               ConsensusEventType = "registry.entity"
	ConsensusEventTypeRegistryNode                 ConsensusEventType = "registry.node"
	ConsensusEventTypeRegistryNodeUnfrozen         ConsensusEventType = "registry.node_unfrozen"
	ConsensusEventTypeRegistryRuntime              ConsensusEventType = "registry.runtime"
	ConsensusEventTypeRoothashExecutionDiscrepancy ConsensusEventType = "roothash.execution_discrepancy"
	ConsensusEventTypeRoothashExecutorCommitted    ConsensusEventType = "roothash.executor_committed"
	ConsensusEventTypeRoothashFinalized            ConsensusEventType = "roothash.finalized"
	ConsensusEventTypeStakingAllowanceChange       ConsensusEventType = "staking.allowance_change"
	ConsensusEventTypeStakingBurn                  ConsensusEventType = "staking.burn"
	ConsensusEventTypeStakingEscrowAdd             ConsensusEventType = "staking.escrow.add"
	ConsensusEventTypeStakingEscrowDebondingStart  ConsensusEventType = "staking.escrow.debonding_start"
	ConsensusEventTypeStakingEscrowReclaim         ConsensusEventType = "staking.escrow.reclaim"
	ConsensusEventTypeStakingEscrowTake            ConsensusEventType = "staking.escrow.take"
	ConsensusEventTypeStakingTransfer              ConsensusEventType = "staking.transfer"
)

// Defines values for ConsensusTxMethod.
const (
	ConsensusTxMethodBeaconPVSSCommit                ConsensusTxMethod = "beacon.PVSSCommit"
	ConsensusTxMethodBeaconPVSSReveal                ConsensusTxMethod = "beacon.PVSSReveal"
	ConsensusTxMethodBeaconVRFProve                  ConsensusTxMethod = "beacon.VRFProve"
	ConsensusTxMethodGovernanceCastVote              ConsensusTxMethod = "governance.CastVote"
	ConsensusTxMethodGovernanceSubmitProposal        ConsensusTxMethod = "governance.SubmitProposal"
	ConsensusTxMethodRegistryRegisterEntity          ConsensusTxMethod = "registry.RegisterEntity"
	ConsensusTxMethodRegistryRegisterNode            ConsensusTxMethod = "registry.RegisterNode"
	ConsensusTxMethodRegistryRegisterRuntime         ConsensusTxMethod = "registry.RegisterRuntime"
	ConsensusTxMethodRoothashExecutorCommit          ConsensusTxMethod = "roothash.ExecutorCommit"
	ConsensusTxMethodRoothashExecutorProposerTimeout ConsensusTxMethod = "roothash.ExecutorProposerTimeout"
	ConsensusTxMethodStakingAddEscrow                ConsensusTxMethod = "staking.AddEscrow"
	ConsensusTxMethodStakingAllow                    ConsensusTxMethod = "staking.Allow"
	ConsensusTxMethodStakingAmendCommissionSchedule  ConsensusTxMethod = "staking.AmendCommissionSchedule"
	ConsensusTxMethodStakingReclaimEscrow            ConsensusTxMethod = "staking.ReclaimEscrow"
	ConsensusTxMethodStakingTransfer                 ConsensusTxMethod = "staking.Transfer"
	ConsensusTxMethodStakingWithdraw                 ConsensusTxMethod = "staking.Withdraw"
)

// Defines values for EvmTokenType.
const (
	ERC1155 EvmTokenType = "ERC1155"
	ERC20   EvmTokenType = "ERC20"
	ERC721  EvmTokenType = "ERC721"
)

// Defines values for RuntimeName.
const (
	Emerald RuntimeName = "emerald"
)

// Account A consensus layer account.
type Account struct {
	// Address The staking address for this account.
	Address string `json:"address"`

	// AddressPreimage The data from which a consensus-style address (`oasis1...`)
	// was derived. Notably, for EVM runtimes like Sapphire,
	// this links the oasis address and the Ethereum address.
	//
	// Oasis addresses are derived from a piece of data, such as an ed25519
	// public key or an Ethereum address. For example, [this](https://github.com/oasisprotocol/oasis-sdk/blob/b37e6da699df331f5a2ac62793f8be099c68469c/client-sdk/go/helpers/address.go#L90-L91)
	// is how an Ethereum is converted to an oasis address. The type of underlying data usually also
	// determines how the signatuers for this address are verified.
	//
	// Consensus supports only "staking addresses" (`context="oasis-core/address: staking"`
	// below; always ed25519-backed).
	// Runtimes support all types. This means that every consensus address is also
	// valid in every runtime. For example, in EVM runtimes, you can use staking
	// addresses, but only with oasis tools (e.g. a wallet); EVM contracts such as
	// ERC20 tokens or tools such as Metamask cannot interact with staking addresses.
	AddressPreimage AddressPreimage `json:"address_preimage"`

	// Allowances The allowances made by this account.
	Allowances []Allowance `json:"allowances"`

	// Available The available balance, in base units.
	Available common.BigInt `json:"available"`

	// Debonding The debonding escrow balance, in base units.
	Debonding common.BigInt `json:"debonding"`

	// DebondingDelegationsBalance The debonding delegations balance, in base units.
	DebondingDelegationsBalance common.BigInt `json:"debonding_delegations_balance"`

	// DelegationsBalance The delegations balance, in base units.
	DelegationsBalance common.BigInt `json:"delegations_balance"`

	// Escrow The active escrow balance, in base units.
	Escrow common.BigInt `json:"escrow"`

	// Nonce A nonce used to prevent replay.
	Nonce           int64            `json:"nonce"`
	RuntimeBalances []RuntimeBalance `json:"runtime_balances"`
}

// AccountList A list of consensus layer accounts.
type AccountList struct {
	Accounts []Account `json:"accounts"`
}

// AddressDerivationContext defines model for AddressDerivationContext.
type AddressDerivationContext string

// AddressPreimage The data from which a consensus-style address (`oasis1...`)
// was derived. Notably, for EVM runtimes like Sapphire,
// this links the oasis address and the Ethereum address.
//
// Oasis addresses are derived from a piece of data, such as an ed25519
// public key or an Ethereum address. For example, [this](https://github.com/oasisprotocol/oasis-sdk/blob/b37e6da699df331f5a2ac62793f8be099c68469c/client-sdk/go/helpers/address.go#L90-L91)
// is how an Ethereum is converted to an oasis address. The type of underlying data usually also
// determines how the signatuers for this address are verified.
//
// Consensus supports only "staking addresses" (`context="oasis-core/address: staking"`
// below; always ed25519-backed).
// Runtimes support all types. This means that every consensus address is also
// valid in every runtime. For example, in EVM runtimes, you can use staking
// addresses, but only with oasis tools (e.g. a wallet); EVM contracts such as
// ERC20 tokens or tools such as Metamask cannot interact with staking addresses.
type AddressPreimage struct {
	// AddressData The hex-encoded data from which the oasis address was derived.
	// When `context = "oasis-runtime-sdk/address: secp256k1eth"`, this
	// is the Ethereum address (without the leading `0x`). All-lowercase.
	AddressData string                   `json:"address_data"`
	Context     AddressDerivationContext `json:"context"`

	// ContextVersion Version of the `context`.
	ContextVersion *int `json:"context_version"`
}

// Allowance defines model for Allowance.
type Allowance struct {
	// Address The allowed account.
	Address string `json:"address"`

	// Amount The amount allowed for the allowed account.
	Amount common.BigInt `json:"amount"`
}

// ApiError defines model for ApiError.
type ApiError struct {
	// Msg An error message.
	Msg *string `json:"msg,omitempty"`
}

// Block A consensus block.
type Block struct {
	// Hash The block header hash.
	Hash string `json:"hash"`

	// Height The block height.
	Height int64 `json:"height"`

	// NumTransactions Number of transactions in the block.
	NumTransactions int32 `json:"num_transactions"`

	// Timestamp The second-granular consensus time.
	Timestamp time.Time `json:"timestamp"`
}

// BlockList A list of consensus blocks.
type BlockList struct {
	Blocks []Block `json:"blocks"`
}

// ConsensusEvent An event emitted by the consensus layer.
type ConsensusEvent struct {
	// Block The block height at which this event was generated.
	Block *int64 `json:"block,omitempty"`

	// Body The event contents. This spec does not encode the many possible types;
	// instead, see [the Go API](https://pkg.go.dev/github.com/oasisprotocol/oasis-core/go/consensus/api/transaction/results#Event) of oasis-core.
	// This object will conform to one of the `*Event` types two levels down
	// the hierarchy, e.g. `TransferEvent` from `Event > staking.Event > TransferEvent`
	Body map[string]interface{} `json:"body"`

	// TxHash Hash of this event's originating transaction.
	// Absent if the event did not originate from a transaction.
	TxHash *string `json:"tx_hash"`

	// TxIndex 0-based index of this event's originating transaction within its block.
	// Absent if the event did not originate from a transaction.
	TxIndex *int32             `json:"tx_index"`
	Type    ConsensusEventType `json:"type"`
}

// ConsensusEventList A list of consensus events.
type ConsensusEventList struct {
	Events []ConsensusEvent `json:"events"`
}

// ConsensusEventType defines model for ConsensusEventType.
type ConsensusEventType string

// ConsensusTxMethod defines model for ConsensusTxMethod.
type ConsensusTxMethod string

// DebondingDelegation A debonding delegation.
type DebondingDelegation struct {
	// Amount The amount of tokens delegated in base units.
	Amount common.BigInt `json:"amount"`

	// DebondEnd The epoch at which the debonding ends.
	DebondEnd int64 `json:"debond_end"`

	// Shares The shares of tokens delegated.
	Shares common.BigInt `json:"shares"`

	// ValidatorAddress The delegatee address.
	ValidatorAddress string `json:"validator_address"`
}

// DebondingDelegationList A list of debonding delegations.
type DebondingDelegationList struct {
	DebondingDelegations []DebondingDelegation `json:"debonding_delegations"`
}

// Delegation A delegation.
type Delegation struct {
	// Amount The amount of tokens delegated in base units.
	Amount common.BigInt `json:"amount"`

	// Shares The shares of tokens delegated.
	Shares common.BigInt `json:"shares"`

	// ValidatorAddress The delegatee address.
	ValidatorAddress string `json:"validator_address"`
}

// DelegationList A list of delegations.
type DelegationList struct {
	Delegations []Delegation `json:"delegations"`
}

// Entity An entity registered at the consensus layer.
type Entity struct {
	// Address The staking address belonging to this entity; derived from the entity's public key.
	Address string `json:"address"`

	// Id The public key identifying this entity.
	ID string `json:"id"`

	// Nodes The vector of nodes owned by this entity.
	Nodes []string `json:"nodes"`
}

// EntityList A list of entities registered at the consensus layer.
type EntityList struct {
	Entities []Entity `json:"entities"`
}

// Epoch A consensus epoch.
type Epoch struct {
	// EndHeight The (inclusive) height at which this epoch ended. Omitted if the epoch is still active.
	EndHeight *uint64 `json:"end_height,omitempty"`

	// Id The epoch number.
	ID int64 `json:"id"`

	// StartHeight The (inclusive) height at which this epoch started.
	StartHeight uint64 `json:"start_height"`
}

// EpochList A list of consensus epochs.
type EpochList struct {
	Epochs []Epoch `json:"epochs"`
}

// EvmTokenType defines model for EvmTokenType.
type EvmTokenType string

// Node A node registered at the consensus layer.
type Node struct {
	// ConsensusPubkey The unique identifier of this node as a consensus member
	ConsensusPubkey string `json:"consensus_pubkey"`

	// EntityId The public key identifying the entity controlling this node.
	EntityID string `json:"entity_id"`

	// Expiration The epoch in which this node's commitment expires.
	Expiration int64 `json:"expiration"`

	// Id The public key identifying this node.
	ID string `json:"id"`

	// P2pPubkey The unique identifier of this node on the P2P transport.
	P2PPubkey string `json:"p2p_pubkey"`

	// Roles A bitmask representing this node's roles.
	Roles string `json:"roles"`

	// TlsNextPubkey The public key that will be used for establishing TLS connections
	// upon rotation.
	TLSNextPubkey string `json:"tls_next_pubkey"`

	// TlsPubkey The public key used for establishing TLS connections.
	TLSPubkey string `json:"tls_pubkey"`
}

// NodeList A list of nodes registered at the consensus layer.
type NodeList struct {
	EntityID string `json:"entity_id"`
	Nodes    []Node `json:"nodes"`
}

// Proposal A governance proposal.
type Proposal struct {
	// Cancels The proposal to cancel, if this proposal proposes
	// cancelling an existing proposal.
	Cancels int64 `json:"cancels"`

	// ClosesAt The epoch at which voting for this proposal will close.
	ClosesAt int64 `json:"closes_at"`

	// CreatedAt The epoch at which this proposal was created.
	CreatedAt int64 `json:"created_at"`

	// Deposit The deposit attached to this proposal.
	Deposit common.BigInt `json:"deposit"`

	// Epoch The epoch at which the proposed upgrade will happen.
	Epoch *uint64 `json:"epoch,omitempty"`

	// Handler The name of the upgrade handler.
	Handler *string `json:"handler,omitempty"`

	// Id The unique identifier of the proposal.
	ID uint64 `json:"id"`

	// InvalidVotes The number of invalid votes for this proposal, after tallying.
	InvalidVotes common.BigInt `json:"invalid_votes"`

	// State The state of the proposal.
	State string `json:"state"`

	// Submitter The staking address of the proposal submitter.
	Submitter string `json:"submitter"`

	// Target The target propotocol versions for this upgrade proposal.
	Target ProposalTarget `json:"target"`
}

// ProposalList A list of governance proposals.
type ProposalList struct {
	Proposals []Proposal `json:"proposals"`
}

// ProposalTarget The target propotocol versions for this upgrade proposal.
type ProposalTarget struct {
	ConsensusProtocol        *string `json:"consensus_protocol,omitempty"`
	RuntimeCommitteeProtocol *string `json:"runtime_committee_protocol,omitempty"`
	RuntimeHostProtocol      *string `json:"runtime_host_protocol,omitempty"`
}

// ProposalVote defines model for ProposalVote.
type ProposalVote struct {
	// Address The staking address casting this vote.
	Address string `json:"address"`

	// Vote The vote cast.
	Vote string `json:"vote"`
}

// ProposalVotes A list of votes for a governance proposal.
type ProposalVotes struct {
	// ProposalId The unique identifier of the proposal.
	ProposalID uint64 `json:"proposal_id"`

	// Votes The list of votes for the proposal.
	Votes []ProposalVote `json:"votes"`
}

// RuntimeBalance Balance of an account in a runtime.
type RuntimeBalance struct {
	// Amount Number of base units held; as a string.
	Amount common.BigInt `json:"amount"`

	// Runtime The name of a runtime. This is a human-readable identifier, and should
	// stay stable across runtime upgrades/versions.
	Runtime RuntimeName `json:"runtime"`

	// TokenId Unique identifier for the token. For EVM tokens, this is their eth address.
	TokenID string `json:"token_id"`

	// TokenSymbol The token ticker symbol. Not guaranteed to be unique across distinct tokens.
	TokenSymbol string `json:"token_symbol"`
}

// RuntimeBlock A ParaTime block.
type RuntimeBlock struct {
	// GasUsed The total gas used by all transactions in the block.
	GasUsed int64 `json:"gas_used"`

	// Hash The block header hash.
	Hash string `json:"hash"`

	// NumTransactions The number of transactions in the block.
	NumTransactions int32 `json:"num_transactions"`

	// Round The block round.
	Round int64 `json:"round"`

	// Size The total byte size of all transactions in the block.
	Size int32 `json:"size"`

	// Timestamp The second-granular consensus time.
	Timestamp time.Time `json:"timestamp"`
}

// RuntimeBlockList A list of consensus blocks.
type RuntimeBlockList struct {
	Blocks []RuntimeBlock `json:"blocks"`
}

// RuntimeName The name of a runtime. This is a human-readable identifier, and should
// stay stable across runtime upgrades/versions.
type RuntimeName string

// RuntimeToken defines model for RuntimeToken.
type RuntimeToken struct {
	// ContractAddr The Oasis address of this token's contract.
	ContractAddr string `json:"contract_addr"`

	// Decimals The number of least significant digits in base units that should be displayed as
	// decimals when displaying tokens. `tokens = base_units / (10**decimals)`.
	// Affects display only. Often equals 18, to match ETH.
	Decimals *int `json:"decimals,omitempty"`

	// Name Name of the token, as provided by token contract's `name()` method.
	Name *string `json:"name,omitempty"`

	// NumHolders The number of addresses that have a nonzero balance of this token,
	// as calculated from Transfer events.
	NumHolders int64 `json:"num_holders"`

	// Symbol Symbol of the token, as provided by token contract's `symbol()` method.
	Symbol *string `json:"symbol,omitempty"`

	// TotalSupply The total number of base units available.
	TotalSupply *string      `json:"total_supply,omitempty"`
	Type        EvmTokenType `json:"type"`
}

// RuntimeTokenList A list of ERC-20 tokens on a runtime.
type RuntimeTokenList struct {
	Tokens []RuntimeToken `json:"tokens"`
}

// RuntimeTransaction A runtime transaction.
type RuntimeTransaction struct {
	// Amount A reasonable "amount" associated with this transaction, if
	// applicable. The meaning varies based on the transaction mehtod.
	// Usually in native denomination, ParaTime units. As a string.
	Amount *string `json:"amount,omitempty"`

	// Body The method call body.
	Body map[string]interface{} `json:"body"`

	// EthHash The Ethereum cryptographic hash of this transaction's encoding.
	// Absent for non-Ethereum-format transactions.
	EthHash *string `json:"eth_hash,omitempty"`

	// Fee The fee that this transaction's sender committed to pay to execute
	// it (total, native denomination, ParaTime base units, as a string).
	Fee string `json:"fee"`

	// GasLimit The maximum gas that this transaction's sender committed to use to
	// execute it.
	GasLimit uint64 `json:"gas_limit"`

	// Hash The Oasis cryptographic hash of this transaction's encoding.
	Hash string `json:"hash"`

	// Method The method that was called.
	Method string `json:"method"`

	// Nonce0 The nonce used with this transaction's 0th signer, to prevent replay.
	Nonce0 uint64 `json:"nonce_0"`

	// Round The block round at which this transaction was executed.
	Round int64 `json:"round"`

	// Sender0 The Oasis address of this transaction's 0th signer.
	// Unlike Ethereum, Oasis natively supports multiple-signature transactions.
	// However, the great majority of transactions only have a single signer in practice.
	// Retrieving the other signers is currently not supported by this API.
	Sender0 string `json:"sender_0"`

	// Success Whether this transaction successfully executed.
	Success bool `json:"success"`

	// Timestamp The second-granular consensus time when this tx's block was proposed.
	Timestamp time.Time `json:"timestamp"`

	// To A reasonable "to" Oasis address associated with this transaction,
	// if applicable. The meaning varies based on the transaction method. Some notable examples:
	//   - For `method = "accounts.Transfer"`, this is the paratime account receiving the funds.
	//   - For `method = "consensus.Deposit"`, this is the paratime account receiving the funds.
	//   - For `method = "consensus.Withdraw"`, this is a consensus (!) account receiving the funds.
	//   - For `method = "evm.Create"`, this is the address of the newly created smart contract.
	//   - For `method = "evm.Call"`, this is the address of the called smart contract
	To *string `json:"to,omitempty"`
}

// RuntimeTransactionList A list of runtime transactions.
type RuntimeTransactionList struct {
	Transactions []RuntimeTransaction `json:"transactions"`
}

// Status defines model for Status.
type Status struct {
	// LatestBlock The height of the most recent indexed block. Query a synced Oasis node for the latest block produced.
	LatestBlock int64 `json:"latest_block"`

	// LatestChainId The most recently indexed chain ID.
	LatestChainID string `json:"latest_chain_id"`

	// LatestUpdate The RFC 3339 formatted time when the Indexer processed the latest block. Compare with current time for approximate indexing progress with the Oasis Network.
	LatestUpdate time.Time `json:"latest_update"`
}

// Transaction A consensus transaction.
type Transaction struct {
	// Block The block height at which this transaction was executed.
	Block int64 `json:"block"`

	// Body The method call body.
	Body []byte `json:"body"`

	// Fee The fee that this transaction's sender committed
	// to pay to execute it.
	Fee common.BigInt `json:"fee"`

	// Hash The cryptographic hash of this transaction's encoding.
	Hash string `json:"hash"`

	// Index 0-based index of this transaction in its block
	Index  int32             `json:"index"`
	Method ConsensusTxMethod `json:"method"`

	// Nonce The nonce used with this transaction, to prevent replay.
	Nonce int64 `json:"nonce"`

	// Sender The address of who sent this transaction.
	Sender string `json:"sender"`

	// Success Whether this transaction successfully executed.
	Success bool `json:"success"`

	// Timestamp The second-granular consensus time this tx's block, i.e. roughly when the
	// [block was proposed](https://github.com/tendermint/tendermint/blob/v0.34.x/spec/core/data_structures.md#header).
	Timestamp time.Time `json:"timestamp"`
}

// TransactionList A list of consensus transactions.
type TransactionList struct {
	Transactions []Transaction `json:"transactions"`
}

// TxVolume defines model for TxVolume.
type TxVolume struct {
	// BucketStart The date for this daily transaction volume measurement.
	BucketStart time.Time `json:"bucket_start"`

	// TxVolume The transaction volume on this day.
	TxVolume uint64 `json:"tx_volume"`
}

// TxVolumeList A list of daily transaction volumes.
type TxVolumeList struct {
	BucketSizeSeconds uint32 `json:"bucket_size_seconds"`

	// Buckets The list of daily transaction volumes.
	Buckets []TxVolume `json:"buckets"`
}

// Validator An validator registered at the consensus layer.
type Validator struct {
	// Active Whether the entity is part of validator set (top <scheduler.params.max_validators> by stake).
	Active                 bool                     `json:"active"`
	CurrentCommissionBound ValidatorCommissionBound `json:"current_commission_bound"`

	// CurrentRate Commission rate.
	CurrentRate uint64 `json:"current_rate"`

	// EntityAddress The staking address identifying this Validator.
	EntityAddress string `json:"entity_address"`

	// EntityId The public key identifying this Validator.
	EntityID string `json:"entity_id"`

	// Escrow The amount staked.
	Escrow common.BigInt   `json:"escrow"`
	Media  *ValidatorMedia `json:"media,omitempty"`

	// NodeId The public key identifying this Validator's node.
	NodeID string `json:"node_id"`

	// Status Whether the entity has a node that is registered for being a validator, node is up to date, and has successfully registered itself. It may or may not be part of validator set.
	Status bool `json:"status"`
}

// ValidatorCommissionBound defines model for ValidatorCommissionBound.
type ValidatorCommissionBound struct {
	EpochEnd   uint64 `json:"epoch_end"`
	EpochStart uint64 `json:"epoch_start"`
	Lower      uint64 `json:"lower"`
	Upper      uint64 `json:"upper"`
}

// ValidatorList A list of validators registered at the consensus layer.
type ValidatorList struct {
	Validators []Validator `json:"validators"`
}

// ValidatorMedia defines model for ValidatorMedia.
type ValidatorMedia struct {
	// EmailAddress An email address for the validator.
	EmailAddress *string `json:"email_address,omitempty"`

	// Logotype A logo type.
	Logotype *string `json:"logotype,omitempty"`

	// Name The human-readable name of this validator.
	Name *string `json:"name,omitempty"`

	// TgChat An Telegram handle.
	TgChat *string `json:"tg_chat,omitempty"`

	// TwitterAcc A Twitter handle.
	TwitterAcc *string `json:"twitter_acc,omitempty"`

	// WebsiteLink An URL associated with the entity.
	WebsiteLink *string `json:"website_link,omitempty"`
}

// GetConsensusAccountsParams defines parameters for GetConsensusAccounts.
type GetConsensusAccountsParams struct {
	// Limit The maximum numbers of items to return.
	Limit *uint64 `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset The number of items to skip before starting to collect the result set.
	Offset *uint64 `form:"offset,omitempty" json:"offset,omitempty"`

	// MinAvailable A filter on the minimum available account balance.
	MinAvailable *common.BigInt `form:"minAvailable,omitempty" json:"minAvailable,omitempty"`

	// MaxAvailable A filter on the maximum available account balance.
	MaxAvailable *common.BigInt `form:"maxAvailable,omitempty" json:"maxAvailable,omitempty"`

	// MinEscrow A filter on the minimum active escrow account balance.
	MinEscrow *common.BigInt `form:"minEscrow,omitempty" json:"minEscrow,omitempty"`

	// MaxEscrow A filter on the maximum active escrow account balance.
	MaxEscrow *common.BigInt `form:"maxEscrow,omitempty" json:"maxEscrow,omitempty"`

	// MinDebonding A filter on the minimum debonding account balance.
	MinDebonding *common.BigInt `form:"minDebonding,omitempty" json:"minDebonding,omitempty"`

	// MaxDebonding A filter on the maximum debonding account balance.
	MaxDebonding *common.BigInt `form:"maxDebonding,omitempty" json:"maxDebonding,omitempty"`

	// MinTotalBalance A filter on the minimum total account balance.
	MinTotalBalance *common.BigInt `form:"minTotalBalance,omitempty" json:"minTotalBalance,omitempty"`

	// MaxTotalBalance A filter on the maximum total account balance.
	MaxTotalBalance *common.BigInt `form:"maxTotalBalance,omitempty" json:"maxTotalBalance,omitempty"`
}

// GetConsensusBlocksParams defines parameters for GetConsensusBlocks.
type GetConsensusBlocksParams struct {
	// Limit The maximum numbers of items to return.
	Limit *uint64 `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset The number of items to skip before starting to collect the result set.
	Offset *uint64 `form:"offset,omitempty" json:"offset,omitempty"`

	// From A filter on minimum block height, inclusive.
	From *int64 `form:"from,omitempty" json:"from,omitempty"`

	// To A filter on maximum block height, inclusive.
	To *int64 `form:"to,omitempty" json:"to,omitempty"`

	// After A filter on minimum block time, inclusive.
	After *time.Time `form:"after,omitempty" json:"after,omitempty"`

	// Before A filter on maximum block time, inclusive.
	Before *time.Time `form:"before,omitempty" json:"before,omitempty"`
}

// GetConsensusEntitiesParams defines parameters for GetConsensusEntities.
type GetConsensusEntitiesParams struct {
	// Limit The maximum numbers of items to return.
	Limit *uint64 `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset The number of items to skip before starting to collect the result set.
	Offset *uint64 `form:"offset,omitempty" json:"offset,omitempty"`
}

// GetConsensusEntitiesEntityIdNodesParams defines parameters for GetConsensusEntitiesEntityIdNodes.
type GetConsensusEntitiesEntityIdNodesParams struct {
	// Limit The maximum numbers of items to return.
	Limit *uint64 `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset The number of items to skip before starting to collect the result set.
	Offset *uint64 `form:"offset,omitempty" json:"offset,omitempty"`
}

// GetConsensusEpochsParams defines parameters for GetConsensusEpochs.
type GetConsensusEpochsParams struct {
	// Limit The maximum numbers of items to return.
	Limit *uint64 `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset The number of items to skip before starting to collect the result set.
	Offset *uint64 `form:"offset,omitempty" json:"offset,omitempty"`
}

// GetConsensusEventsParams defines parameters for GetConsensusEvents.
type GetConsensusEventsParams struct {
	// Limit The maximum numbers of items to return.
	Limit *uint64 `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset The number of items to skip before starting to collect the result set.
	Offset *uint64 `form:"offset,omitempty" json:"offset,omitempty"`

	// Block A filter on block height.
	Block *int64 `form:"block,omitempty" json:"block,omitempty"`

	// TxIndex A filter on transaction index. The returned events all need to originate
	// from a transaction that appeared in `tx_index`-th position in the block.
	// It is invalid to specify this filter without also specifying a `block`.
	// Specifying `tx_index` and `block` is an alternative to specifying `tx_hash`;
	// either works to fetch events from a specific transaction.
	TxIndex *int32 `form:"tx_index,omitempty" json:"tx_index,omitempty"`

	// TxHash A filter on the hash of the transaction that originated the events.
	// Specifying `tx_index` and `block` is an alternative to specifying `tx_hash`;
	// either works to fetch events from a specific transaction.
	TxHash *string `form:"tx_hash,omitempty" json:"tx_hash,omitempty"`

	// Rel A filter on related accounts. Every returned event will refer to
	// this account. For example, for a `Transfer` event, this will be the
	// the sender or the recipient of tokens.
	Rel *string `form:"rel,omitempty" json:"rel,omitempty"`

	// Type A filter on the event type.
	Type *ConsensusEventType `form:"type,omitempty" json:"type,omitempty"`
}

// GetConsensusProposalsParams defines parameters for GetConsensusProposals.
type GetConsensusProposalsParams struct {
	// Limit The maximum numbers of items to return.
	Limit *uint64 `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset The number of items to skip before starting to collect the result set.
	Offset *uint64 `form:"offset,omitempty" json:"offset,omitempty"`

	// Submitter The submitter of the proposal.
	Submitter *string `form:"submitter,omitempty" json:"submitter,omitempty"`

	// State The state of the proposal.
	State *string `form:"state,omitempty" json:"state,omitempty"`
}

// GetConsensusProposalsProposalIdVotesParams defines parameters for GetConsensusProposalsProposalIdVotes.
type GetConsensusProposalsProposalIdVotesParams struct {
	// Limit The maximum numbers of items to return.
	Limit *uint64 `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset The number of items to skip before starting to collect the result set.
	Offset *uint64 `form:"offset,omitempty" json:"offset,omitempty"`
}

// GetConsensusStatsTxVolumeParams defines parameters for GetConsensusStatsTxVolume.
type GetConsensusStatsTxVolumeParams struct {
	// Limit The maximum numbers of items to return.
	Limit *uint64 `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset The number of items to skip before starting to collect the result set.
	Offset *uint64 `form:"offset,omitempty" json:"offset,omitempty"`

	// BucketSizeSeconds The size of buckets into which the statistic is grouped, in seconds.
	// The backend supports a limited number of bucket sizes: 300 (5 minutes) and
	// 3600 (1 hour). Requests with other values may be rejected.
	BucketSizeSeconds *int32 `form:"bucket_size_seconds,omitempty" json:"bucket_size_seconds,omitempty"`
}

// GetConsensusTransactionsParams defines parameters for GetConsensusTransactions.
type GetConsensusTransactionsParams struct {
	// Limit The maximum numbers of items to return.
	Limit *uint64 `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset The number of items to skip before starting to collect the result set.
	Offset *uint64 `form:"offset,omitempty" json:"offset,omitempty"`

	// Block A filter on block height.
	Block *int64 `form:"block,omitempty" json:"block,omitempty"`

	// Method A filter on transaction method.
	Method *ConsensusTxMethod `form:"method,omitempty" json:"method,omitempty"`

	// Sender A filter on transaction sender.
	Sender *string `form:"sender,omitempty" json:"sender,omitempty"`

	// Rel A filter on related accounts.
	Rel *string `form:"rel,omitempty" json:"rel,omitempty"`

	// MinFee A filter on minimum transaction fee, inclusive.
	MinFee *int64 `form:"minFee,omitempty" json:"minFee,omitempty"`

	// MaxFee A filter on maximum transaction fee, inclusive.
	MaxFee *int64 `form:"maxFee,omitempty" json:"maxFee,omitempty"`

	// Code A filter on transaction status code.
	Code *int `form:"code,omitempty" json:"code,omitempty"`
}

// GetConsensusValidatorsParams defines parameters for GetConsensusValidators.
type GetConsensusValidatorsParams struct {
	// Limit The maximum numbers of items to return.
	Limit *uint64 `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset The number of items to skip before starting to collect the result set.
	Offset *uint64 `form:"offset,omitempty" json:"offset,omitempty"`
}

// GetEmeraldBlocksParams defines parameters for GetEmeraldBlocks.
type GetEmeraldBlocksParams struct {
	// Limit The maximum numbers of items to return.
	Limit *uint64 `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset The number of items to skip before starting to collect the result set.
	Offset *uint64 `form:"offset,omitempty" json:"offset,omitempty"`

	// From A filter on minimum block height, inclusive.
	From *int64 `form:"from,omitempty" json:"from,omitempty"`

	// To A filter on maximum block height, inclusive.
	To *int64 `form:"to,omitempty" json:"to,omitempty"`

	// After A filter on minimum block time, inclusive.
	After *time.Time `form:"after,omitempty" json:"after,omitempty"`

	// Before A filter on maximum block time, inclusive.
	Before *time.Time `form:"before,omitempty" json:"before,omitempty"`
}

// GetEmeraldTokensParams defines parameters for GetEmeraldTokens.
type GetEmeraldTokensParams struct {
	// Limit The maximum numbers of items to return.
	Limit *uint64 `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset The number of items to skip before starting to collect the result set.
	Offset *uint64 `form:"offset,omitempty" json:"offset,omitempty"`
}

// GetEmeraldTransactionsParams defines parameters for GetEmeraldTransactions.
type GetEmeraldTransactionsParams struct {
	// Limit The maximum numbers of items to return.
	Limit *uint64 `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset The number of items to skip before starting to collect the result set.
	Offset *uint64 `form:"offset,omitempty" json:"offset,omitempty"`

	// Block A filter on block round.
	Block *int64 `form:"block,omitempty" json:"block,omitempty"`

	// Rel A filter on related accounts. Every returned transaction will refer to
	// this account in a way. For example, for an `accounts.Transfer` tx, this will be
	// the sender or the recipient of tokens.
	// The indexer detects related accounts inside EVM transactions and events on a
	// best-effort basis. For example, it inspects ERC20 methods inside `evm.Call` txs.
	// However, you must provide the oasis-style derived address here, not the Eth address.
	// See `AddressPreimage` for more info on oasis-style vs Eth addresses.
	Rel *string `form:"rel,omitempty" json:"rel,omitempty"`
}
