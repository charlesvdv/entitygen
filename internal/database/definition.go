package database

import (
	"errors"
	"fmt"
)

var (
	ErrAlreadyExists = errors.New("already exists")
)

func NewDefinition(defaultSchema string) Definition {
	def := Definition{
		defaultSchema: defaultSchema,
		schemas:       []*Schema{},
	}
	def.CreateSchema(defaultSchema)
	return def
}

type Definition struct {
	defaultSchema string
	schemas       []*Schema
}

func (def Definition) DefaultSchema() *Schema {
	return def.Schema(def.defaultSchema)
}

func (def Definition) Schemas() []*Schema {
	return def.schemas
}

func (def Definition) Schema(name string) *Schema {
	for i := range def.schemas {
		if def.schemas[i].Name() == name {
			return def.schemas[i]
		}
	}
	return nil
}

func (def *Definition) CreateSchema(name string) (*Schema, error) {
	if def.Schema(name) != nil {
		return nil, fmt.Errorf("schema '%s' %w", name, ErrAlreadyExists)
	}
	newSchema := newSchema(name)
	def.schemas = append(def.schemas, &newSchema)
	return &newSchema, nil
}

func newSchema(name string) Schema {
	return Schema{
		name:   name,
		tables: []*Table{},
	}
}

type Schema struct {
	name   string
	tables []*Table
}

func (schema Schema) Name() string {
	return schema.name
}

func (schema Schema) Tables() []*Table {
	return schema.tables
}

func (schema Schema) Table(name string) *Table {
	for i := range schema.tables {
		if schema.tables[i].Name() == name {
			return schema.tables[i]
		}
	}
	return nil
}

func (schema *Schema) CreateTable(name string) (*Table, error) {
	if schema.Table(name) != nil {
		return nil, fmt.Errorf("table '%s' %w", name, ErrAlreadyExists)
	}
	table := newTable(name)
	schema.tables = append(schema.tables, &table)
	return &table, nil
}

func newTable(name string) Table {
	return Table{
		name:    name,
		columns: []*Column{},
	}
}

type Table struct {
	name    string
	columns []*Column
}

func (table Table) Name() string {
	return table.name
}

func (table Table) Columns() []*Column {
	return table.columns
}

func (table Table) Column(name string) *Column {
	for i := range table.columns {
		if table.columns[i].Name() == name {
			return table.columns[i]
		}
	}

	return nil
}

func (table *Table) AddColumn(column Column) error {
	if table.Column(column.Name()) != nil {
		return fmt.Errorf("column '%s' %w", column.Name(), ErrAlreadyExists)
	}
	table.columns = append(table.columns, &column)
	return nil
}

type Column struct {
	name  string
	_type Type
}

func NewColumn(name string, _type Type) (Column, error) {
	if name == "" {
		return Column{}, errors.New("column name cannot be empty")
	}

	return Column{
		name:  name,
		_type: _type,
	}, nil
}

func (column Column) Name() string {
	return column.name
}

func (column Column) Type() Type {
	return column._type
}
