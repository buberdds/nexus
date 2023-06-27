package statecheck

import (
	"context"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	oasisConfig "github.com/oasisprotocol/oasis-sdk/client-sdk/go/config"
	"github.com/oasisprotocol/oasis-sdk/client-sdk/go/connection"
	"github.com/stretchr/testify/require"

	"github.com/oasisprotocol/nexus/common"
	"github.com/oasisprotocol/nexus/log"
	"github.com/oasisprotocol/nexus/storage"
	"github.com/oasisprotocol/nexus/storage/postgres"
)

const (
	ChainName = common.ChainNameMainnet
)

func newTargetClient(t *testing.T) (*postgres.Client, error) {
	connString := os.Getenv("HEALTHCHECK_TEST_CONN_STRING")
	logger, err := log.NewLogger("db-test", io.Discard, log.FmtJSON, log.LevelInfo)
	require.Nil(t, err)

	return postgres.NewClient(connString, logger)
}

func newSdkConnection(ctx context.Context) (connection.Connection, error) {
	net := &oasisConfig.Network{
		ChainContext: os.Getenv("HEALTHCHECK_TEST_CHAIN_CONTEXT"),
		RPC:          os.Getenv("HEALTHCHECK_TEST_NODE_RPC"),
	}
	return connection.ConnectNoVerify(ctx, net)
}

func snapshotBackends(target *postgres.Client, analyzer string, tables []string) (int64, error) {
	ctx := context.Background()

	batch := &storage.QueryBatch{}
	batch.Queue(`CREATE SCHEMA IF NOT EXISTS snapshot;`)
	for _, t := range tables {
		batch.Queue(fmt.Sprintf(`
			DROP TABLE IF EXISTS snapshot.%s CASCADE;
		`, t))
		batch.Queue(fmt.Sprintf(`
			CREATE TABLE snapshot.%[1]s AS TABLE chain.%[1]s;
		`, t))
	}
	batch.Queue(`
		INSERT INTO snapshot.snapshotted_heights (analyzer, height)
			SELECT analyzer, height FROM analysis.processed_blocks WHERE analyzer=$1 ORDER BY height DESC, processed_time DESC LIMIT 1
			ON CONFLICT DO NOTHING;
	`, analyzer)

	// Create the snapshot using a high level of isolation; we don't want another
	// tx to be able to modify the tables while this is running, creating a snapshot that
	// represents indexer state at two (or more) blockchain heights.
	if err := target.SendBatchWithOptions(ctx, batch, pgx.TxOptions{IsoLevel: pgx.Serializable}); err != nil {
		return 0, err
	}

	var snapshotHeight int64
	if err := target.QueryRow(ctx, `
		SELECT height from snapshot.snapshotted_heights
			WHERE analyzer=$1
			ORDER BY height DESC LIMIT 1;
	`, analyzer).Scan(&snapshotHeight); err != nil {
		return 0, err
	}

	return snapshotHeight, nil
}
