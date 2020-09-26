package schema

import (
	"errors"
	"fmt"
)

var (
	ErrDuplicate = errors.New("duplicate")
)

func NewDefinitionBuilder() *DefinitionBuilder {
	return &DefinitionBuilder{
		schemas: map[string]SchemaBuilder{},
	}
}

type DefinitionBuilder struct {
	schemas       map[string]SchemaBuilder
	defaultSchema string
}

func (builder *DefinitionBuilder) Build() Definition {
	definition := Definition{
		defaultSchema: builder.defaultSchema,
		schemas:       []Schema{},
	}

	for schemaName, schemaVal := range builder.schemas {
		schema := Schema{
			name:   schemaName,
			tables: []Table{},
		}
		for tableName, _ := range schemaVal.tables {
			table := Table{
				name: tableName,
			}
			schema.tables = append(schema.tables, table)
		}
		definition.schemas = append(definition.schemas, schema)
	}

	return definition
}

func (builder *DefinitionBuilder) WithDefaultSchema(name string) *DefinitionBuilder {
	builder.defaultSchema = name
	err := builder.CreateSchema(builder.defaultSchema)
	if err != nil && !errors.Is(err, ErrDuplicate) {
		panic("failed to set default schema")
	}
	return builder
}

func (builder *DefinitionBuilder) CreateSchema(name string) error {
	if _, ok := builder.schemas[name]; ok {
		return fmt.Errorf("schema '%s' is %w", name, ErrDuplicate)
	}

	builder.schemas[name] = newSchemaBuilder(name)
	return nil
}

func newSchemaBuilder(name string) SchemaBuilder {
	return SchemaBuilder{
		name:   name,
		tables: map[string]TableBuilder{},
	}
}

type SchemaBuilder struct {
	name   string
	tables map[string]TableBuilder
}

type TableBuilder struct {
	name    string
	columns map[string]column
}

type column struct {
	name  string
	_type string
}
