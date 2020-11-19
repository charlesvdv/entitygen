package postgres_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/charlesvdv/entitygen/internal/database"
	"github.com/charlesvdv/entitygen/internal/database/parser/postgres"
)

func parseOneSQL(t *testing.T, sql string) database.Definition {
	parser := postgres.NewDDLParser()
	err := parser.ParseSQL(sql)
	require.NoError(t, err)
	return parser.GetResultingSchemaDefinition()
}

func parseOneSQLError(t *testing.T, sql string) error {
	parser := postgres.NewDDLParser()
	err := parser.ParseSQL(sql)
	require.Error(t, err)
	return err
}

func requireBuiltinType(t *testing.T, expected string, _type database.Type) {
	builtinType, ok := _type.(*database.BuiltinType)
	require.Truef(t, ok, "should be a builtin type")
	require.Equal(t, builtinType.Name(), expected)
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
	test1Schema := definition.Schema("test1")
	require.NotNil(t, test1Schema)

	table := test1Schema.Table("test")
	require.NotNil(t, table)
	require.Equal(t, 1, len(test1Schema.Tables()))

	require.Equal(t, 2, len(table.Columns()))
	idColumn := table.Column("id")
	require.NotNil(t, idColumn)
	require.Equal(t, "id", idColumn.Name())
	requireBuiltinType(t, database.TypeBigInt, idColumn.Type())

	testColumn := table.Column("test")
	require.NotNil(t, testColumn)
	require.Equal(t, "test", testColumn.Name())
	requireBuiltinType(t, database.TypeString, testColumn.Type())
}

func TestCreateTableOnDefaultSchema(t *testing.T) {
	const sql = `
		CREATE TABLE test (
			id SERIAL PRIMARY KEY
		)
	`

	definition := parseOneSQL(t, sql)
	require.Equal(t, 1, len(definition.Schemas()))

	defaultSchema := definition.DefaultSchema()
	require.NotNil(t, defaultSchema)
	require.Equal(t, 1, len(defaultSchema.Tables()))
	require.Equal(t, "public", defaultSchema.Name())

	testTable := defaultSchema.Table("test")
	require.NotNil(t, testTable)
	require.Equal(t, "test", testTable.Name())
}

func TestColumnBuiltinDataType(t *testing.T) {
	const sql = `
		CREATE TABLE test (
			serial SERIAL,
			bigint BIGINT,
			int8 INT8,
			bool bool,
			boolean boolean,
			float8 float8,
			int4 int4,
			int int,
			integer integer,
			float8 float8,
			doublepres double precision,
			float4 float4,
			real real,
			time time,
			timestamp timestamp,
			date date
		)
	`

	table := parseOneSQL(t, sql).DefaultSchema().Table("test")
	require.NotNil(t, table)
	requireBuiltinType(t, database.TypeBigInt, table.Column("serial").Type())
	requireBuiltinType(t, database.TypeBigInt, table.Column("bigint").Type())
	requireBuiltinType(t, database.TypeBigInt, table.Column("int8").Type())
	requireBuiltinType(t, database.TypeBoolean, table.Column("bool").Type())
	requireBuiltinType(t, database.TypeBoolean, table.Column("boolean").Type())
	requireBuiltinType(t, database.TypeInteger, table.Column("int4").Type())
	// TODO: on https://www.postgresql.org/docs/current/datatype.html, `int` should be aliased to `integer`.
	// but here, it returns `BIGINT` why??
	requireBuiltinType(t, database.TypeBigInt, table.Column("int").Type())
	requireBuiltinType(t, database.TypeBigInt, table.Column("integer").Type())
	requireBuiltinType(t, database.TypeDecimal, table.Column("float8").Type())
	requireBuiltinType(t, database.TypeDecimal, table.Column("doublepres").Type())
	requireBuiltinType(t, database.TypeFloat, table.Column("float4").Type())
	requireBuiltinType(t, database.TypeFloat, table.Column("real").Type())
	requireBuiltinType(t, database.TypeTime, table.Column("time").Type())
	requireBuiltinType(t, database.TypeTimestamp, table.Column("timestamp").Type())
	requireBuiltinType(t, database.TypeDate, table.Column("date").Type())
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
