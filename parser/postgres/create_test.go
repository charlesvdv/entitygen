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
		CREATE SCHEMA test1;
		CREATE TABLE test1.test (
			id SERIAL PRIMARY KEY,
			test TEXT 
		);
	`

	definition := parseOneSQL(t, sql)
	schema := definition.Schema("test1")
	require.NotNil(t, schema)

	table := schema.Table("test")
	require.NotNil(t, table)
	require.Equal(t, 1, len(schema.Tables()))

	require.Equal(t, 2, len(table.Columns()))
	idColumn := table.Column("id")
	require.NotNil(t, idColumn)
	require.Equal(t, "id", idColumn.Name())
	require.Equal(t, "INT8", idColumn.Type())

	testColumn := table.Column("test")
	require.NotNil(t, testColumn)
	require.Equal(t, "test", testColumn.Name())
	require.Equal(t, "STRING", testColumn.Type())
}

func TestCreateSchema(t *testing.T) {
	definition := parseOneSQL(t, "CREATE SCHEMA test;")
	require.NotNil(t, definition.Schema("test"))
	require.Equal(t, definition.Schema("test").Name(), "test")
	// default schema + "test" schema
	require.Equal(t, 2, len(definition.Schemas()))
}

func TestCreateSchemaDuplicate(t *testing.T) {
	parseOneSQLError(t, "CREATE SCHEMA test; CREATE SCHEMA test;")
}

func TestCreateSchemaIfNotExists(t *testing.T) {
	definition := parseOneSQL(t, "CREATE SCHEMA test; CREATE SCHEMA IF NOT EXISTS test;")
	require.NotNil(t, definition.Schema("test"))
	require.Equal(t, definition.Schema("test").Name(), "test")
}
