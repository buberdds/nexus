package genesis

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/iancoleman/strcase"
	staking "github.com/oasisprotocol/oasis-core/go/staking/api"
	oasisConfig "github.com/oasisprotocol/oasis-sdk/client-sdk/go/config"
	"github.com/stretchr/testify/require"

	"github.com/oasislabs/oasis-indexer/log"
	"github.com/oasislabs/oasis-indexer/storage"
	"github.com/oasislabs/oasis-indexer/storage/oasis"
	"github.com/oasislabs/oasis-indexer/storage/postgres"
)

type TestAccount struct {
	Address   string
	Nonce     uint64
	Available uint64
	Escrow    uint64
	Debonding uint64
}

func newTargetClient(t *testing.T) (*postgres.Client, error) {
	connString := os.Getenv("CI_TEST_CONN_STRING")
	logger, err := log.NewLogger("cockroach-test", ioutil.Discard, log.FmtJSON, log.LevelInfo)
	require.Nil(t, err)

	return postgres.NewClient(connString, logger)
}

func newSourceClient(t *testing.T) (*oasis.Client, error) {
	network := &oasisConfig.Network{
		ChainContext: os.Getenv("CI_TEST_CHAIN_CONTEXT"),
		RPC:          os.Getenv("CI_TEST_NODE_RPC"),
	}
	return oasis.NewClient(context.Background(), network)
}

func checkpointStakingBackend(t *testing.T, source *oasis.Client, target *postgres.Client) (int64, error) {
	ctx := context.Background()

	doc, err := source.GenesisDocument(ctx)
	require.Nil(t, err)
	chainID := strcase.ToSnake(doc.ChainID)

	// Prepare checkpoint queries.
	batch := &storage.QueryBatch{}

	batch.Queue("BEGIN WORK;")

	// Lock staking backend tables.
	batch.Queue(fmt.Sprintf(`
		LOCK TABLE %s.accounts IN ROW EXCLUSIVE MODE;
	`, chainID))
	batch.Queue(fmt.Sprintf(`
		LOCK TABLE %s.allowances IN ROW EXCLUSIVE MODE;
	`, chainID))
	batch.Queue(fmt.Sprintf(`
		LOCK TABLE %s.delegations IN ROW EXCLUSIVE MODE;
	`, chainID))
	batch.Queue(fmt.Sprintf(`
		LOCK TABLE %s.debonding_delegations IN ROW EXCLUSIVE MODE;
	`, chainID))
	batch.Queue(fmt.Sprintf(`
		LOCK TABLE %s.processed_blocks IN ROW EXCLUSIVE MODE;
	`, chainID))

	// Create checkpoints.
	batch.Queue(fmt.Sprintf(`
		TRUNCATE %s.accounts_checkpoint CASCADE;
	`, chainID))
	batch.Queue(fmt.Sprintf(`
		INSERT INTO %s.accounts_checkpoint (address, general_balance, nonce, escrow_balance_active, escrow_total_shares_active, escrow_balance_debonding, escrow_total_shares_debonding, extra_data)
		SELECT address, general_balance, nonce, escrow_balance_active, escrow_total_shares_active, escrow_balance_debonding, escrow_total_shares_debonding, extra_data FROM %s.accounts;
	`, chainID, chainID))

	batch.Queue(fmt.Sprintf(`
		TRUNCATE %s.allowances_checkpoint CASCADE;
	`, chainID))
	batch.Queue(fmt.Sprintf(`
		INSERT INTO %s.allowances_checkpoint (owner, beneficiary, allowance)
		SELECT owner, beneficiary, allowance FROM %s.allowances;
	`, chainID, chainID))

	batch.Queue(fmt.Sprintf(`
		TRUNCATE %s.delegations_checkpoint CASCADE;
	`, chainID))
	batch.Queue(fmt.Sprintf(`
		INSERT INTO %s.delegations_checkpoint (delegatee, delegator, shares)
		SELECT delegatee, delegator, shares FROM %s.delegations;
	`, chainID, chainID))

	batch.Queue(fmt.Sprintf(`
		TRUNCATE %s.debonding_delegations_checkpoint CASCADE;
	`, chainID))
	batch.Queue(fmt.Sprintf(`
		INSERT INTO %s.debonding_delegations_checkpoint (delegatee, delegator, shares, debond_end)
		SELECT delegatee, delegator, shares, debond_end FROM %s.debonding_delegations;
	`, chainID, chainID))

	batch.Queue(fmt.Sprintf(`
		INSERT INTO %s.checkpointed_heights (height)
		SELECT height FROM %s.processed_blocks ORDER BY height DESC, processed_time DESC LIMIT 1;
	`, chainID, chainID))

	// Release locks and commit!
	batch.Queue("COMMIT WORK;")

	if err := target.SendBatch(ctx, batch); err != nil {
		return 0, err
	}

	var checkpointHeight int64
	if err := target.QueryRow(ctx, fmt.Sprintf(`
		SELECT height FROM %s.checkpointed_heights
		ORDER BY height DESC LIMIT 1;
	`, chainID)).Scan(&checkpointHeight); err != nil {
		return 0, nil
	}

	return checkpointHeight, nil
}

func TestStakingGenesis(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}
	if _, ok := os.LookupEnv("CI_TEST_GENESIS"); !ok {
		t.Skip("skipping test since genesis tests are not enabled")
	}

	t.Log("Initializing data stores...")

	ctx := context.Background()

	oasisClient, err := newSourceClient(t)
	require.Nil(t, err)

	postgresClient, err := newTargetClient(t)
	require.Nil(t, err)

	doc, err := oasisClient.GenesisDocument(ctx)
	require.Nil(t, err)
	chainID := strcase.ToSnake(doc.ChainID)

	t.Log("Creating checkpoint...")

	height, err := checkpointStakingBackend(t, oasisClient, postgresClient)
	require.Nil(t, err)

	t.Log("Fetching genesis state...")

	doc, err = oasisClient.StateToGenesis(ctx, height)
	require.Nil(t, err)

	t.Log("Validating...")

	stakingGenesis := doc.Staking
	rows, err := postgresClient.Query(ctx, fmt.Sprintf(
		`SELECT address, nonce, general_balance, escrow_balance_active, escrow_balance_debonding
				FROM %s.accounts`, chainID),
	)
	require.Nil(t, err)
	for rows.Next() {
		var ta TestAccount
		err = rows.Scan(
			&ta.Address,
			&ta.Nonce,
			&ta.Available,
			&ta.Escrow,
			&ta.Debonding,
		)
		require.Nil(t, err)

		var address staking.Address
		err = address.UnmarshalText([]byte(ta.Address))
		require.Nil(t, err)

		a, ok := stakingGenesis.Ledger[address]
		require.True(t, ok)
		require.Equal(t, a.General.Nonce, ta.Nonce)
		require.Equal(t, a.General.Balance.ToBigInt().Uint64(), ta.Available)
	}

	t.Log("Done!")
}
