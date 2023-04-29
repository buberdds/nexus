package postgres

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/oasisprotocol/oasis-indexer/common"
	"github.com/oasisprotocol/oasis-indexer/log"
	"github.com/oasisprotocol/oasis-indexer/storage"
	"github.com/oasisprotocol/oasis-indexer/tests"
)

func newClient(t *testing.T) (*Client, error) {
	connString := os.Getenv("CI_TEST_CONN_STRING")
	logger, err := log.NewLogger("postgres-test", io.Discard, log.FmtJSON, log.LevelInfo)
	require.Nil(t, err)

	return NewClient(connString, logger)
}

func TestConnect(t *testing.T) {
	tests.SkipIfShort(t)

	client, err := newClient(t)
	require.Nil(t, err)
	client.Close()
}

func TestInvalidConnect(t *testing.T) {
	tests.SkipIfShort(t)

	connString := "an invalid connstring"
	logger, err := log.NewLogger("postgres-test", io.Discard, log.FmtJSON, log.LevelInfo)
	require.Nil(t, err)

	_, err = NewClient(connString, logger)
	require.NotNil(t, err)
}

func TestQuery(t *testing.T) {
	tests.SkipIfShort(t)

	client, err := newClient(t)
	require.Nil(t, err)
	defer client.Close()

	rows, err := client.Query(context.Background(), `
		SELECT * FROM ( VALUES (0),(1),(2) ) AS q;
	`)
	require.Nil(t, err)

	i := 0
	for rows.Next() {
		var result int
		err = rows.Scan(&result)
		require.Nil(t, err)
		require.Equal(t, i, result)

		i++
	}
	require.Equal(t, 3, i)
}

func TestInvalidQuery(t *testing.T) {
	tests.SkipIfShort(t)

	client, err := newClient(t)
	require.Nil(t, err)
	defer client.Close()

	_, err = client.Query(context.Background(), `
		an invalid query
	`)
	require.NotNil(t, err)
}

func TestQueryRow(t *testing.T) {
	tests.SkipIfShort(t)

	client, err := newClient(t)
	require.Nil(t, err)
	defer client.Close()

	var result int
	err = client.QueryRow(context.Background(), `
		SELECT 1+1;
	`).Scan(&result)
	require.Nil(t, err)
	require.Equal(t, 2, result)
}

func TestInvalidQueryRow(t *testing.T) {
	tests.SkipIfShort(t)

	client, err := newClient(t)
	require.Nil(t, err)
	defer client.Close()

	var result int
	err = client.QueryRow(context.Background(), `
		an invalid query
	`).Scan(&result)
	require.NotNil(t, err)
}

func TestSendBatch(t *testing.T) {
	tests.SkipIfShort(t)

	client, err := newClient(t)
	require.Nil(t, err)
	defer client.Close()

	defer func() {
		destroy := &storage.QueryBatch{}
		destroy.Queue(`
			DROP TABLE films;
		`)
		err := client.SendBatch(context.Background(), destroy)
		require.Nil(t, err)
	}()

	create := &storage.QueryBatch{}
	create.Queue(`
		CREATE TABLE films (
			fid  INTEGER PRIMARY KEY,
			name TEXT
		);
	`)
	err = client.SendBatch(context.Background(), create)
	require.Nil(t, err)

	insert := &storage.QueryBatch{}
	queueFilms := func(b *storage.QueryBatch, f []string, idOffset int) {
		rows := make([]string, 0, len(f))
		for i, film := range f {
			rows = append(rows, fmt.Sprintf("(%d, '%s')", i+idOffset, film))
		}
		b.Queue(fmt.Sprintf(`
			INSERT INTO films (fid, name)
			VALUES %s;
		`, strings.Join(rows, ", ")))
	}

	films1 := []string{
		"Gone with the Wind",
		"Avatar",
		"Titanic",
	}
	films2 := []string{
		"Star Wars",
		"Avengers: Endgame",
	}
	queueFilms(insert, films1, 0)
	queueFilms(insert, films2, len(films1))
	err = client.SendBatch(context.Background(), insert)
	require.Nil(t, err)

	var wg sync.WaitGroup
	for i, film := range append(films1, films2...) {
		wg.Add(1)
		go func(i int, film string) {
			defer wg.Done()

			var result string
			err := client.QueryRow(context.Background(), `
				SELECT name FROM films WHERE fid = $1;
			`, i).Scan(&result)
			require.Nil(t, err)
			require.Equal(t, film, result)
		}(i, film)
	}

	wg.Wait()
}

func TestInvalidSendBatch(t *testing.T) {
	tests.SkipIfShort(t)

	client, err := newClient(t)
	require.Nil(t, err)
	defer client.Close()

	invalid := &storage.QueryBatch{}
	invalid.Queue(`
		an invalid query
	`)
	err = client.SendBatch(context.Background(), invalid)
	require.NotNil(t, err)
}

func TestNumeric(t *testing.T) {
	client, err := newClient(t)
	require.Nil(t, err)
	defer client.Close()

	// Create custom type, derived from NUMERIC.
	_, err = client.pool.Exec(context.Background(), `CREATE DOMAIN mynumeric NUMERIC(1000,0) CHECK(VALUE >= 0)`)
	if err != nil {
		t.Fatal(err.Error())
	}

	// Test that we can scan both null and non-null values into a *common.BigInt.
	var mynull *common.BigInt
	var my2 *common.BigInt
	err = client.QueryRow(context.Background(), `
		SELECT null::mynumeric, 2::mynumeric;
	`).Scan(&mynull, &my2)
	if err != nil {
		t.Fatal(err.Error())
	}
	require.Nil(t, err)
	require.Nil(t, mynull)
	require.Equal(t, int64(2), my2.Int64())
}
