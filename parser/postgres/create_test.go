package postgres_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/charlesvdv/entitygen/parser/postgres"
	"github.com/charlesvdv/entitygen/schema"
)

func parseOneSQL(t *testing.T, sql string) schema.Definition {
	parser := postgres.NewDDLParser()
	err := parser.ParseStatement(sql)
	require.NoError(t, err)
	return parser.GetResultingSchemaDefinition()
}

func parseOneSQLError(t *testing.T, sql string) error {
	parser := postgres.NewDDLParser()
	err := parser.ParseStatement(sql)
	require.Error(t, err)
	return err
}

func TestCreateTable(t *testing.T) {
	const sql = `
		CREATE TABLE test (
			id SERIAL PRIMARY KEY,
			test TEXT 
		);
	`

	parseOneSQL(t, sql)
}

func TestCreateSchema(t *testing.T) {
	definition := parseOneSQL(t, "CREATE SCHEMA test;")
	require.NotNil(t, definition.Schema("test"))
	require.Equal(t, definition.Schema("test").Name(), "test")
}

func TestCreateSchemaDuplicate(t *testing.T) {
	parseOneSQLError(t, "CREATE SCHEMA test; CREATE SCHEMA test;")
}

func TestCreateSchemaIfNotExists(t *testing.T) {
	definition := parseOneSQL(t, "CREATE SCHEMA test; CREATE SCHEMA IF NOT EXISTS test;")
	require.NotNil(t, definition.Schema("test"))
	require.Equal(t, definition.Schema("test").Name(), "test")
}
