package schema

import (
	"errors"
	"fmt"
)

var (
	ErrDuplicate      = errors.New("duplicate")
	ErrNotFound       = errors.New("not found")
	ErrSchemaNotFound = fmt.Errorf("schema %w", ErrNotFound)
)

func NewDefinitionBuilder() *DefinitionBuilder {
	return &DefinitionBuilder{
		defaultSchema: "",
		schemas:       map[string]schemaBuilder{},
	}
}

type DefinitionBuilder struct {
	schemas       map[string]schemaBuilder
	defaultSchema string
}

type schemaBuilder = []*TableBuilder

func (builder *DefinitionBuilder) Build() Definition {
	definition := Definition{
		defaultSchema: builder.defaultSchema,
		schemas:       []Schema{},
	}

	for schemaName, schemaBuilder := range builder.schemas {
		schema := Schema{
			name:   schemaName,
			tables: []Table{},
		}

		for _, tableBuilder := range schemaBuilder {
			table := Table{
				name:    tableBuilder.name,
				columns: tableBuilder.columns,
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

func (builder *DefinitionBuilder) DefaultSchema() string {
	return builder.defaultSchema
}

func (builder *DefinitionBuilder) CreateSchema(name string) error {
	if _, ok := builder.schemas[name]; ok {
		return fmt.Errorf("schema '%s' is %w", name, ErrDuplicate)
	}

	builder.schemas[name] = schemaBuilder{}
	return nil
}

func (builder *DefinitionBuilder) CreateTable(schemaName string, tableName string) (*TableBuilder, error) {
	schema, ok := builder.schemas[schemaName]
	if !ok {
		return nil, fmt.Errorf("'%s': ", ErrSchemaNotFound)
	}

	for _, tableBuilder := range schema {
		if tableBuilder.name == tableName {
			return nil, fmt.Errorf("table '%s' is already created: %w", tableName, ErrDuplicate)
		}
	}

	tableBuilder := newTableBuilder(tableName)
	builder.schemas[schemaName] = append(builder.schemas[schemaName], tableBuilder)
	return tableBuilder, nil
}

func newTableBuilder(name string) *TableBuilder {
	return &TableBuilder{
		name:    name,
		columns: []Column{},
	}
}

type TableBuilder struct {
	name    string
	columns []Column
}

func newColumn(name string, _type string) Column {
	return Column{
		name:  name,
		_type: _type,
	}
}

func (builder *TableBuilder) AddColumn(name string, _type string) error {
	for _, column := range builder.columns {
		if column.Name() == name {
			return fmt.Errorf("column '%s' is %w", name, ErrDuplicate)
		}
	}
	builder.columns = append(builder.columns, newColumn(name, _type))
	return nil
}
