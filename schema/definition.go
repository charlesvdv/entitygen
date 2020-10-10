package schema

type Definition struct {
	defaultSchema string
	schemas       []Schema
}

func (def Definition) DefaultSchema() string {
	return def.defaultSchema
}

func (def Definition) Schemas() []Schema {
	return def.schemas
}

func (def Definition) Schema(name string) *Schema {
	for _, schema := range def.schemas {
		if schema.Name() == name {
			return &schema
		}
	}
	return nil
}

type Schema struct {
	name   string
	tables []Table
}

func (schema Schema) Name() string {
	return schema.name
}

func (schema Schema) Tables() []Table {
	return schema.tables
}

func (schema Schema) Table(name string) *Table {
	for _, table := range schema.tables {
		if table.Name() == name {
			return &table
		}
	}
	return nil
}

type Table struct {
	name    string
	columns []Column
}

func (table Table) Name() string {
	return table.name
}

func (table Table) Columns() []Column {
	return table.columns
}

func (table Table) Column(name string) *Column {
	for _, column := range table.columns {
		if column.Name() == name {
			return &column
		}
	}

	return nil
}

type Column struct {
	name  string
	_type string
}

func (column Column) Name() string {
	return column.name
}

func (column Column) Type() string {
	return column._type
}
