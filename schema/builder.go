package schema

import (
	"errors"
	"fmt"
)

var (
	ErrAlreadyExist = errors.New("already exist")
)

func NewDefinitionBuilder() *DefinitionBuilder {
	return &DefinitionBuilder{
		Definition: newDefinition(),
	}
}

type DefinitionBuilder struct {
	Definition
}

func (builder *DefinitionBuilder) WithDefaultSchema(name string) *DefinitionBuilder {
	builder.Definition.defaultSchema = name
	return builder
}

func (builder *DefinitionBuilder) Build() Definition {
	return builder.Definition
}

func (builder *DefinitionBuilder) AddNewSchema(name string) error {
	schemas := builder.Definition.schemas
	if _, ok := schemas[name]; ok {
		return fmt.Errorf("schema '%s' %w", name, ErrAlreadyExist)
	}

	schemas[name] = newSchema(name)
	return nil
}
