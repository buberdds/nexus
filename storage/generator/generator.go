// Package generator generates migrations for the Oasis Indexer
// from the genesis file at a particular height.
package generator

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/oasisprotocol/oasis-core/go/common/entity"
	"github.com/oasisprotocol/oasis-core/go/common/node"
	genesis "github.com/oasisprotocol/oasis-core/go/genesis/api"
	registry "github.com/oasisprotocol/oasis-core/go/registry/api"
	staking "github.com/oasisprotocol/oasis-core/go/staking/api"

	"github.com/oasislabs/oasis-indexer/log"
)

const bulkInsertBatchSize = 1000

// MigrationGenerator generates migrations for the Oasis Indexer
// target storage.
type MigrationGenerator struct {
	logger *log.Logger
}

// NewMigrationGenerator creates a new migration generator.
func NewMigrationGenerator(logger *log.Logger) *MigrationGenerator {
	return &MigrationGenerator{logger}
}

// WriteGenesisDocumentMigrationOasis3 creates a new migration that re-initializes all
// height-dependent state as per the provided genesis document.
func (mg *MigrationGenerator) WriteGenesisDocumentMigrationOasis3(w io.Writer, document *genesis.Document) error {
	if _, err := io.WriteString(w, `-- DO NOT MODIFY
-- This file was autogenerated by the oasis-indexer migration generator.
`); err != nil {
		return err
	}

	for _, f := range []func(io.Writer, *genesis.Document) error{
		mg.addRegistryBackendMigrations,
		mg.addStakingBackendMigrations,
		mg.addGovernanceBackendMigrations,
	} {
		if err := f(w, document); err != nil {
			return err
		}
	}

	return nil
}

func (mg *MigrationGenerator) addRegistryBackendMigrations(w io.Writer, document *genesis.Document) error {
	chainID := strcase.ToSnake(document.ChainID)

	// Populate entities.
	if _, err := io.WriteString(w, fmt.Sprintf(`
-- Registry Backend Data
TRUNCATE %s.entities CASCADE;`, chainID)); err != nil {
		return err
	}
	if _, err := io.WriteString(w, fmt.Sprintf(`
INSERT INTO %s.entities (id, address)
VALUES
`, chainID)); err != nil {
		return err
	}
	for i, signedEntity := range document.Registry.Entities {
		var entity entity.Entity
		if err := signedEntity.Open(registry.RegisterEntitySignatureContext, &entity); err != nil {
			return err
		}

		if _, err := io.WriteString(w, fmt.Sprintf(
			"\t('%s', '%s')",
			entity.ID.String(),
			staking.NewAddress(entity.ID).String(),
		)); err != nil {
			return err
		}

		if i != len(document.Registry.Entities)-1 {
			if _, err := io.WriteString(w, ",\n"); err != nil {
				return err
			}
		}
	}
	if _, err := io.WriteString(w, ";\n"); err != nil {
		return err
	}

	// Populate nodes.
	if _, err := io.WriteString(w, fmt.Sprintf(`
TRUNCATE %s.nodes CASCADE;`, chainID)); err != nil {
		return err
	}
	if _, err := io.WriteString(w, fmt.Sprintf(`
INSERT INTO %s.nodes (id, entity_id, expiration, tls_pubkey, tls_next_pubkey, p2p_pubkey, consensus_pubkey, roles)
VALUES
`, chainID)); err != nil {
		return err
	}
	for i, signedNode := range document.Registry.Nodes {
		var node node.Node
		if err := signedNode.Open(registry.RegisterNodeSignatureContext, &node); err != nil {
			return err
		}

		if _, err := io.WriteString(w, fmt.Sprintf(
			"\t('%s', '%s', %d, '%s', '%s', '%s', '%s', '%s')",
			node.ID.String(),
			node.EntityID.String(),
			node.Expiration,
			node.TLS.PubKey.String(),
			node.TLS.NextPubKey.String(),
			node.P2P.ID.String(),
			node.Consensus.ID.String(),
			node.Roles.String(),
		)); err != nil {
			return err
		}

		if i != len(document.Registry.Nodes)-1 {
			if _, err := io.WriteString(w, ",\n"); err != nil {
				return err
			}
		}
	}
	if _, err := io.WriteString(w, ";\n"); err != nil {
		return err
	}

	// Populate runtimes.
	if _, err := io.WriteString(w, fmt.Sprintf(`
TRUNCATE %s.runtimes CASCADE;`, chainID)); err != nil {
		return err
	}

	if len(document.Registry.Runtimes) > 0 {
		if _, err := io.WriteString(w, fmt.Sprintf(`
INSERT INTO %s.runtimes (id, suspended, kind, tee_hardware, key_manager)
VALUES
`, chainID)); err != nil {
			return err
		}
		for i, runtime := range document.Registry.Runtimes {
			keyManager := "none"
			if runtime.KeyManager != nil {
				keyManager = runtime.KeyManager.String()
			}
			if _, err := io.WriteString(w, fmt.Sprintf(
				"\t('%s', %t, '%s', '%s', '%s')",
				runtime.ID.String(),
				false,
				runtime.Kind.String(),
				runtime.TEEHardware.String(),
				keyManager,

				// TODO(ennsharma): Add extra_data.
			)); err != nil {
				return err
			}

			if i != len(document.Registry.Runtimes)-1 {
				if _, err := io.WriteString(w, ",\n"); err != nil {
					return err
				}
			}
		}
		if _, err := io.WriteString(w, ";\n"); err != nil {
			return err
		}
	}

	if len(document.Registry.SuspendedRuntimes) > 0 {
		if _, err := io.WriteString(w, fmt.Sprintf(`
INSERT INTO %s.runtimes (id, suspended, kind, tee_hardware, key_manager)
VALUES
`, chainID)); err != nil {
			return err
		}

		for i, runtime := range document.Registry.SuspendedRuntimes {
			keyManager := "none"
			if runtime.KeyManager != nil {
				keyManager = runtime.KeyManager.Hex()
			}
			if _, err := io.WriteString(w, fmt.Sprintf(
				"\t('%s', %t, '%s', '%s', '%s')",
				runtime.ID.String(),
				true,
				runtime.Kind.String(),
				runtime.TEEHardware.String(),
				keyManager,

				// TODO(ennsharma): Add extra_data.
			)); err != nil {
				return err
			}

			if i != len(document.Registry.SuspendedRuntimes)-1 {
				if _, err := io.WriteString(w, ",\n"); err != nil {
					return err
				}
			}
		}
		if _, err := io.WriteString(w, ";\n"); err != nil {
			return err
		}
	}

	return nil
}

func (mg *MigrationGenerator) addStakingBackendMigrations(w io.Writer, document *genesis.Document) error {
	chainID := strcase.ToSnake(document.ChainID)

	// Populate accounts.
	if _, err := io.WriteString(w, fmt.Sprintf(`
-- Staking Backend Data
TRUNCATE %s.accounts CASCADE;`, chainID)); err != nil {
		return err
	}

	// Populate special accounts with reserved addresses.
	if _, err := io.WriteString(w, fmt.Sprintf(`
-- Reserved addresses
INSERT INTO %s.accounts (address, general_balance, nonce, escrow_balance_active, escrow_total_shares_active, escrow_balance_debonding, escrow_total_shares_debonding)
VALUES
`, chainID)); err != nil {
		return err
	}

	reservedAccounts := make(map[staking.Address]*staking.Account)

	commonPoolAccount := staking.Account{
		General: staking.GeneralAccount{
			Balance: document.Staking.CommonPool,
		},
	}
	feeAccumulatorAccount := staking.Account{
		General: staking.GeneralAccount{
			Balance: document.Staking.LastBlockFees,
		},
	}
	governanceDepositsAccount := staking.Account{
		General: staking.GeneralAccount{
			Balance: document.Staking.GovernanceDeposits,
		},
	}

	reservedAccounts[staking.CommonPoolAddress] = &commonPoolAccount
	reservedAccounts[staking.FeeAccumulatorAddress] = &feeAccumulatorAccount
	reservedAccounts[staking.GovernanceDepositsAddress] = &governanceDepositsAccount

	i := 0
	for address, account := range reservedAccounts {
		if _, err := io.WriteString(w, fmt.Sprintf(
			"\t('%s', %d, %d, %d, %d, %d, %d)",
			address.String(),
			account.General.Balance.ToBigInt(),
			account.General.Nonce,
			account.Escrow.Active.Balance.ToBigInt(),
			account.Escrow.Active.TotalShares.ToBigInt(),
			account.Escrow.Debonding.Balance.ToBigInt(),
			account.Escrow.Debonding.TotalShares.ToBigInt(),
		)); err != nil {
			return err
		}

		if i != len(reservedAccounts)-1 {
			if _, err := io.WriteString(w, ",\n"); err != nil {
				return err
			}
		} else {
			if _, err := io.WriteString(w, ";\n"); err != nil {
				return err
			}
		}
		i++
	}

	if _, err := io.WriteString(w, fmt.Sprintf(`
INSERT INTO %s.accounts (address, general_balance, nonce, escrow_balance_active, escrow_total_shares_active, escrow_balance_debonding, escrow_total_shares_debonding)
VALUES
`, chainID)); err != nil {
		return err
	}

	i = 0
	for address, account := range document.Staking.Ledger {
		if _, err := io.WriteString(w, fmt.Sprintf(
			"\t('%s', %d, %d, %d, %d, %d, %d)",
			address.String(),
			account.General.Balance.ToBigInt(),
			account.General.Nonce,
			account.Escrow.Active.Balance.ToBigInt(),
			account.Escrow.Active.TotalShares.ToBigInt(),
			account.Escrow.Debonding.Balance.ToBigInt(),
			account.Escrow.Debonding.TotalShares.ToBigInt(),
		)); err != nil {
			return err
		}
		i++

		if i%bulkInsertBatchSize == 0 {
			if _, err := io.WriteString(w, ";\n"); err != nil {
				return err
			}
			if _, err := io.WriteString(w, fmt.Sprintf(`
INSERT INTO %s.accounts (address, general_balance, nonce, escrow_balance_active, escrow_total_shares_active, escrow_balance_debonding, escrow_total_shares_debonding)
VALUES
`, chainID)); err != nil {
				return err
			}
		} else if i != len(document.Staking.Ledger) {
			if _, err := io.WriteString(w, ",\n"); err != nil {
				return err
			}
		}
	}
	if _, err := io.WriteString(w, ";\n"); err != nil {
		return err
	}

	// Populate commissions.
	// This likely won't overflow batch limit.
	if _, err := io.WriteString(w, fmt.Sprintf(`
TRUNCATE %s.commissions CASCADE;`, chainID)); err != nil {
		return err
	}

	if _, err := io.WriteString(w, fmt.Sprintf(`
INSERT INTO %s.commissions (address, schedule) VALUES
`, chainID)); err != nil {
		return err
	}

	commissions := make([]string, 0)

	for address, account := range document.Staking.Ledger {
		if len(account.Escrow.CommissionSchedule.Rates) > 0 || len(account.Escrow.CommissionSchedule.Bounds) > 0 {
			schedule, err := json.Marshal(account.Escrow.CommissionSchedule)
			if err != nil {
				return err
			}

			commissions = append(commissions, fmt.Sprintf(
				"\t('%s', '%s')",
				address.String(),
				string(schedule),
			))
		}
	}

	for index, commission := range commissions {
		if _, err := io.WriteString(w, commission); err != nil {
			return err
		}

		if index != len(commissions)-1 {
			if _, err := io.WriteString(w, ",\n"); err != nil {
				return err
			}
		} else {
			if _, err := io.WriteString(w, ";\n"); err != nil {
				return err
			}
		}
	}

	// Populate allowances.
	if _, err := io.WriteString(w, fmt.Sprintf(`
TRUNCATE %s.allowances CASCADE;`, chainID)); err != nil {
		return err
	}

	foundAllowances := false // in case allowances are empty

	i = 0
	for owner, account := range document.Staking.Ledger {
		if len(account.General.Allowances) > 0 && foundAllowances {
			if _, err := io.WriteString(w, ",\n"); err != nil {
				return err
			}
		}

		ownerAllowances := make([]string, len(account.General.Allowances))
		j := 0
		for beneficiary, allowance := range account.General.Allowances {
			ownerAllowances[j] = fmt.Sprintf(
				"\t('%s', '%s', %d)",
				owner.String(),
				beneficiary.String(),
				allowance.ToBigInt(),
			)
			j++
		}
		if len(account.General.Allowances) > 0 && !foundAllowances {
			if _, err := io.WriteString(w, fmt.Sprintf(`
INSERT INTO %s.allowances (owner, beneficiary, allowance)
VALUES
`, chainID)); err != nil {
				return err
			}
			foundAllowances = true
		}

		if _, err := io.WriteString(w, strings.Join(ownerAllowances, ",\n")); err != nil {
			return err
		}
		i++
	}
	if foundAllowances {
		if _, err := io.WriteString(w, ";\n"); err != nil {
			return err
		}
	}

	// Populate delegations.
	if _, err := io.WriteString(w, fmt.Sprintf(`
TRUNCATE %s.delegations CASCADE;`, chainID)); err != nil {
		return err
	}
	if _, err := io.WriteString(w, fmt.Sprintf(`
INSERT INTO %s.delegations (delegatee, delegator, shares)
VALUES
`, chainID)); err != nil {
		return err
	}
	i = 0
	j := 0
	for delegatee, escrows := range document.Staking.Delegations {
		k := 0
		for delegator, delegation := range escrows {
			if _, err := io.WriteString(w, fmt.Sprintf(
				"\t('%s', '%s', %d)",
				delegatee.String(),
				delegator.String(),
				delegation.Shares.ToBigInt(),
			)); err != nil {
				return err
			}
			i++

			if i%bulkInsertBatchSize == 0 {
				if _, err := io.WriteString(w, ";\n"); err != nil {
					return err
				}
				if _, err := io.WriteString(w, fmt.Sprintf(`
INSERT INTO %s.delegations (delegatee, delegator, shares)
VALUES
`, chainID)); err != nil {
					return err
				}
			} else if !(k == len(escrows)-1 && j == len(document.Staking.Delegations)-1) {
				if _, err := io.WriteString(w, ",\n"); err != nil {
					return err
				}
			}
			k++
		}
		j++
	}
	if _, err := io.WriteString(w, ";\n"); err != nil {
		return err
	}

	// Populate debonding delegations.
	if _, err := io.WriteString(w, fmt.Sprintf(`
TRUNCATE %s.debonding_delegations CASCADE;`, chainID)); err != nil {
		return err
	}
	if _, err := io.WriteString(w, fmt.Sprintf(`
INSERT INTO %s.debonding_delegations (delegatee, delegator, shares, debond_end)
VALUES
`, chainID)); err != nil {
		return err
	}
	i = 0
	for delegatee, escrows := range document.Staking.DebondingDelegations {
		delegateeDebondingDelegations := make([]string, 0)
		j := 0
		for delegator, debondingDelegations := range escrows {
			delegatorDebondingDelegations := make([]string, len(debondingDelegations))
			for k, debondingDelegation := range debondingDelegations {
				delegatorDebondingDelegations[k] = fmt.Sprintf(
					"\t('%s', '%s', %d, %d)",
					delegatee.String(),
					delegator.String(),
					debondingDelegation.Shares.ToBigInt(),
					debondingDelegation.DebondEndTime,
				)
			}
			delegateeDebondingDelegations = append(delegateeDebondingDelegations, delegatorDebondingDelegations...)
			j++
		}
		if _, err := io.WriteString(w, strings.Join(delegateeDebondingDelegations, ",\n")); err != nil {
			return err
		}
		i++

		if i != len(document.Staking.DebondingDelegations) && len(escrows) > 0 {
			if _, err := io.WriteString(w, ",\n"); err != nil {
				return err
			}
		}
	}
	if _, err := io.WriteString(w, ";\n"); err != nil {
		return err
	}

	return nil
}

func (mg *MigrationGenerator) addGovernanceBackendMigrations(w io.Writer, document *genesis.Document) error {
	chainID := strcase.ToSnake(document.ChainID)

	// Populate proposals.
	if _, err := io.WriteString(w, fmt.Sprintf(`
-- Governance Backend Data
TRUNCATE %s.proposals CASCADE;`, chainID)); err != nil {
		return err
	}

	if len(document.Governance.Proposals) > 0 {

		// TODO(ennsharma): Extract `executed` for proposal.
		if _, err := io.WriteString(w, fmt.Sprintf(`
INSERT INTO %s.proposals (id, submitter, state, deposit, handler, cp_target_version, rhp_target_version, rcp_target_version, upgrade_epoch, cancels, created_at, closes_at, invalid_votes)
VALUES
`, chainID)); err != nil {
			return err
		}

		for i, proposal := range document.Governance.Proposals {
			if proposal.Content.Upgrade != nil {
				if _, err := io.WriteString(w, fmt.Sprintf(
					"\t(%d, '%s', '%s', %d, '%s', '%s', '%s', '%s', %d, %s, %d, %d, %d)",
					proposal.ID,
					proposal.Submitter.String(),
					proposal.State.String(),
					proposal.Deposit.ToBigInt(),
					proposal.Content.Upgrade.Handler,
					proposal.Content.Upgrade.Target.ConsensusProtocol.String(),
					proposal.Content.Upgrade.Target.RuntimeHostProtocol.String(),
					proposal.Content.Upgrade.Target.RuntimeCommitteeProtocol.String(),
					proposal.Content.Upgrade.Epoch,
					"null",
					proposal.CreatedAt,
					proposal.ClosesAt,
					proposal.InvalidVotes,
				)); err != nil {
					return err
				}
			} else if proposal.Content.CancelUpgrade != nil {
				if _, err := io.WriteString(w, fmt.Sprintf(
					"\t(%d, '%s', '%s', %d, '%s', '%s', '%s', '%s', '%s', %d, %d, %d, %d)",
					proposal.ID,
					proposal.Submitter.String(),
					proposal.State.String(),
					proposal.Deposit.ToBigInt(),
					"",
					"",
					"",
					"",
					"",
					proposal.Content.CancelUpgrade.ProposalID,
					proposal.CreatedAt,
					proposal.ClosesAt,
					proposal.InvalidVotes,
				)); err != nil {
					return err
				}
			}

			if i != len(document.Governance.Proposals)-1 {
				if _, err := io.WriteString(w, ",\n"); err != nil {
					return err
				}
			}
		}
		if _, err := io.WriteString(w, ";\n"); err != nil {
			return err
		}
	}

	// Populate votes.
	if _, err := io.WriteString(w, fmt.Sprintf(`
TRUNCATE %s.votes CASCADE;`, chainID)); err != nil {
		return err
	}

	foundVotes := false // in case votes are empty

	i := 0
	for proposalID, voteEntries := range document.Governance.VoteEntries {
		if len(voteEntries) > 0 && !foundVotes {
			if _, err := io.WriteString(w, fmt.Sprintf(`
INSERT INTO %s.votes (proposal, voter, vote)
VALUES
`, chainID)); err != nil {
				return err
			}
			foundVotes = true
		}
		votes := make([]string, len(voteEntries))
		for j, voteEntry := range voteEntries {
			votes[j] = fmt.Sprintf(
				"\t(%d, '%s', '%s')",
				proposalID,
				voteEntry.Voter.String(),
				voteEntry.Vote.String(),
			)
		}
		if _, err := io.WriteString(w, strings.Join(votes, ",\n")); err != nil {
			return err
		}
		i++

		if i != len(document.Governance.VoteEntries) && len(voteEntries) > 0 {
			if _, err := io.WriteString(w, ",\n"); err != nil {
				return err
			}
		}
	}
	if foundVotes {
		if _, err := io.WriteString(w, ";\n"); err != nil {
			return err
		}
	}

	return nil
}
