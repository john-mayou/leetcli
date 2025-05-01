package testutil

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/john-mayou/leetcli/config"
	"github.com/john-mayou/leetcli/db"
	"github.com/stretchr/testify/require"

	_ "github.com/lib/pq" // postgres driver
)

func SetupTestClient(t *testing.T) *db.Client {
	config, err := config.LoadConfig()
	require.NoError(t, err)

	database := SetupDB(t, config)
	ResetDB(t, database)

	return db.NewClient(database)
}

func SetupDB(t *testing.T, cfg *config.Config) *sqlx.DB {
	db, err := sqlx.Connect("postgres", cfg.DatabaseURL)
	require.NoError(t, err)

	return db
}

func ResetDB(t *testing.T, db *sqlx.DB) {
	query := `
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = 'public'
			AND table_type = 'BASE TABLE'
			AND table_name != 'schema_migrations'
	`
	var tables []struct {
		Name string `db:"table_name"`
	}
	err := db.Select(&tables, query)
	require.NoError(t, err)

	var sb strings.Builder
	for _, table := range tables {
		sb.WriteString(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE;\n", table.Name))
	}

	_, err = db.Exec(sb.String())
	require.NoError(t, err)
}
