package postgres

import (
	"errors"
	"fmt"

	"github.com/charlesvdv/entitygen/database"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
)

func dispatchStatementToBuilder(stmt tree.Statement, schemaDefinition *database.Definition) error {
	dispatcher := newStatementDispatcher(schemaDefinition)
	return dispatcher.run(stmt)
}

type statementDispatcher struct {
	schemaDefinition *database.Definition
}

func newStatementDispatcher(schemaDefinition *database.Definition) statementDispatcher {
	return statementDispatcher{
		schemaDefinition: schemaDefinition,
	}
}

func (d *statementDispatcher) run(stmt tree.Statement) error {
	var err error = nil

	switch concreteStmt := stmt.(type) {
	case *tree.CreateTable:
		err = d.dispatchCreateTable(concreteStmt)
	case *tree.CreateSchema:
		err = d.dispatchCreateSchema(concreteStmt)
	default:
		return errors.New("unknown statement")
	}

	return err
}

func (d *statementDispatcher) dispatchCreateTable(stmt *tree.CreateTable) error {
	schemaDef := d.schemaDefinition.DefaultSchema()
	schemaName := stmt.Table.SchemaName.Normalize()
	if schemaName != "" {
		schemaDef = d.schemaDefinition.Schema(schemaName)
	}

	table, err := schemaDef.CreateTable(stmt.Table.ObjectName.Normalize())
	if errors.Is(err, database.ErrAlreadyExists) && stmt.IfNotExists {
		// table already exists, so we skip it
		return nil
	}
	if err != nil {
		return fmt.Errorf("'%s': %w", stmt.Table.Table(), err)
	}

	for _, tableDef := range stmt.Defs {
		switch concreteTableDef := tableDef.(type) {
		case *tree.ColumnTableDef:
			column, err := database.NewColumn(concreteTableDef.Name.Normalize(), convertType(concreteTableDef.Type))
			if err != nil {
				return fmt.Errorf("'%s': fail to add column: %w", stmt.Table.Table(), err)
			}
			table.AddColumn(column)
		default:
			// do nothing for now
		}
	}

	return nil
}

func (d *statementDispatcher) dispatchCreateSchema(stmt *tree.CreateSchema) error {
	_, err := d.schemaDefinition.CreateSchema(stmt.Schema)
	if errors.Is(err, database.ErrAlreadyExists) && stmt.IfNotExists {
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}
