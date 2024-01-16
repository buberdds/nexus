package statecheck

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	sdkConfig "github.com/oasisprotocol/oasis-sdk/client-sdk/go/config"
	sdkTypes "github.com/oasisprotocol/oasis-sdk/client-sdk/go/types"

	common "github.com/oasisprotocol/nexus/common"
)

var RuntimeTables = []string{"runtime_sdk_balances"}

type TestRuntimeAccount struct {
	Address string
	Balance common.BigInt
	Symbol  string
}

func TestEmeraldAccounts(t *testing.T) {
	testRuntimeAccounts(t, common.RuntimeEmerald)
}

func TestSapphireAccounts(t *testing.T) {
	testRuntimeAccounts(t, common.RuntimeSapphire)
}

func testRuntimeAccounts(t *testing.T, runtime common.Runtime) {
	if _, ok := os.LookupEnv("OASIS_INDEXER_HEALTHCHECK"); !ok {
		t.Skip("skipping test since healthcheck tests are not enabled")
	}

	t.Log("Initializing data stores...")

	ctx := context.Background()

	sdkNet := sdkConfig.DefaultNetworks.All[string(ChainName)]
	sdkPT := sdkNet.ParaTimes.All[string(runtime)]
	t.Log("Runtime ID determined", "runtime", runtime, "runtime_id", sdkPT.ID)

	conn, err := newSdkConnection(ctx)
	require.NoError(t, err)
	oasisRuntimeClient := conn.Runtime(sdkPT)

	postgresClient, err := newTargetClient(t)
	require.NoError(t, err)

	t.Log("Creating snapshot for runtime tables...")
	height, err := snapshotBackends(postgresClient, string(runtime), RuntimeTables)
	require.NoError(t, err)

	t.Logf("Fetching accounts information at height %d...", height)
	addresses, err := oasisRuntimeClient.Accounts.Addresses(ctx, uint64(height), sdkTypes.NativeDenomination)
	require.NoError(t, err)
	expectedAccts := make(map[sdkTypes.Address]bool)
	for _, addr := range addresses {
		expectedAccts[addr] = true
	}
	t.Logf("Fetched %d account addresses", len(addresses))

	acctRows, err := postgresClient.Query(ctx,
		`SELECT account_address, balance, symbol
			FROM snapshot.runtime_sdk_balances
			WHERE runtime=$1`, runtime,
	)
	require.NoError(t, err)
	actualAccts := make(map[string]bool)
	var allBalances uint64
	var balanceDiscrepancies uint64
	for acctRows.Next() {
		var a TestRuntimeAccount
		err = acctRows.Scan(
			&a.Address,
			&a.Balance,
			&a.Symbol,
		)
		assert.NoError(t, err)

		// Check that the account exists.
		var actualAddr sdkTypes.Address
		err = actualAddr.UnmarshalText([]byte(a.Address))
		assert.NoError(t, err)

		_, ok := expectedAccts[actualAddr]
		if !ok {
			t.Logf("address %s found, but not expected", a.Address)
			t.Fail()
			continue
		}

		// Check that the account balance is accurate.
		balances, err := oasisRuntimeClient.Accounts.Balances(ctx, uint64(height), actualAddr)
		assert.NoError(t, err)
		for denom, amount := range balances.Balances {
			if stringifyDenomination(denom, sdkPT) == a.Symbol {
				allBalances++
				if !assert.Equal(t, *amount.ToBigInt(), a.Balance.Int, "address: %s", a.Address) {
					balanceDiscrepancies++
				}
			}
		}
		actualAccts[a.Address] = true
	}
	t.Logf("Number of discrepancies in account balances: %d (out of: %d)", balanceDiscrepancies, allBalances)

	for _, addr := range addresses {
		_, ok := actualAccts[addr.String()]
		if !ok {
			t.Logf("address %s expected, but not found", addr.String())
			t.Fail()
		}
	}
}

func nativeTokenSymbol(sdkPT *sdkConfig.ParaTime) string {
	return sdkPT.Denominations[sdkConfig.NativeDenominationKey].Symbol
}

// StringifyDenomination returns a string representation denomination `d`
// in the context of `runtimeID`. The context matters for the native denomination.
func stringifyDenomination(d sdkTypes.Denomination, sdkPT *sdkConfig.ParaTime) string {
	if d.IsNative() {
		return nativeTokenSymbol(sdkPT)
	}

	return d.String()
}
