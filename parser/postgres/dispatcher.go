package postgres

import (
	"errors"

	"github.com/charlesvdv/entitygen/schema"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
)

func dispatchStatementToBuilder(stmt tree.Statement, schemaBuilder *schema.DefinitionBuilder) error {
	dispatcher := newStatementDispatcher(schemaBuilder)
	return dispatcher.run(stmt)
}

type statementDispatcher struct {
	schemaBuilder *schema.DefinitionBuilder
}

func newStatementDispatcher(schemaBuilder *schema.DefinitionBuilder) statementDispatcher {
	return statementDispatcher{
		schemaBuilder: schemaBuilder,
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
	schemaName := stmt.Table.SchemaName.Normalize()
	if schemaName == "" {
		schemaName = d.schemaBuilder.DefaultSchema()
	}

	tableBuilder, err := d.schemaBuilder.CreateTable(schemaName, stmt.Table.ObjectName.Normalize())
	if errors.Is(err, schema.ErrDuplicate) && stmt.IfNotExists {
		// table already exists, so we skip it
		return nil
	}
	if err != nil {
		return err
	}

	for _, tableDef := range stmt.Defs {
		switch concreteTableDef := tableDef.(type) {
		case *tree.ColumnTableDef:
			tableBuilder.AddColumn(concreteTableDef.Name.Normalize(), concreteTableDef.Type.SQLString())
		default:
			// do nothing for now
		}
	}

	return nil
}

func (d *statementDispatcher) dispatchCreateSchema(stmt *tree.CreateSchema) error {
	err := d.schemaBuilder.CreateSchema(stmt.Schema)
	if errors.Is(err, schema.ErrDuplicate) && stmt.IfNotExists {
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}
