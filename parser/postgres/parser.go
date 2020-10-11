package postgres

import (
	"github.com/cockroachdb/cockroach/pkg/sql/parser"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"

	"github.com/charlesvdv/entitygen/schema"
)

const (
	defaultSchema = "public"
)

func NewDDLParser() DDLParser {
	schemaDefinition := schema.NewDefinition(defaultSchema)

	return DDLParser{
		schemaDefinition: &schemaDefinition,
	}
}

type DDLParser struct {
	schemaDefinition *schema.Definition
}

func (p DDLParser) GetResultingSchemaDefinition() schema.Definition {
	return *p.schemaDefinition
}

func (p DDLParser) ParseStatement(rawSql string) error {
	statements, err := parser.Parse(rawSql)
	if err != nil {
		return err
	}

	for _, stmt := range statements {
		err := p.applyStatement(stmt)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p DDLParser) applyStatement(stmt parser.Statement) error {
	if !tree.CanModifySchema(stmt.AST) {
		// Skip non DDL statements
		return nil
	}

	return dispatchStatementToBuilder(stmt.AST, p.schemaDefinition)
}
