package postgres

import (
	"github.com/cockroachdb/cockroach/pkg/sql/parser"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"

	"github.com/charlesvdv/entitygen/internal/database"
)

const (
	defaultSchema = "public"
)

func NewDDLParser() DDLParser {
	schemaDefinition := database.NewDefinition(defaultSchema)

	return DDLParser{
		schemaDefinition: &schemaDefinition,
	}
}

type DDLParser struct {
	schemaDefinition *database.Definition
}

func (p DDLParser) GetResultingSchemaDefinition() database.Definition {
	return *p.schemaDefinition
}

func (p DDLParser) ParseSQL(rawSQL string) error {
	statements, err := parser.Parse(rawSQL)
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
