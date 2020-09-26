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

type Table struct {
	name string
}
