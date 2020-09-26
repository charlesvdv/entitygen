package schema

func newDefinition() Definition {
	return Definition{
		schemas: map[string]Schema{},
	}
}

type Definition struct {
	defaultSchema string
	schemas       map[string]Schema
}

func (def Definition) DefaultSchema() Schema {
	return def.schemas[def.defaultSchema]
}

func (def Definition) Schema(name string) *Schema {
	if schema, ok := def.schemas[name]; ok {
		return &schema
	}
	return nil
}

func newSchema(name string) Schema {
	return Schema{
		name: name,
	}
}

type Schema struct {
	name string
}

func (schema Schema) Name() string {
	return schema.name
}
