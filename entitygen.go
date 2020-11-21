package entitygen

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/charlesvdv/entitygen/internal/database"
	"github.com/charlesvdv/entitygen/internal/database/parser/postgres"
	"github.com/charlesvdv/entitygen/internal/gogen"
)

const (
	SQL_VARIANT_POSTGRES = "POSTGRES"
)

type databaseParser interface {
	ParseSQL(rawSQL string) error
	GetResultingSchemaDefinition() database.Definition
}

func NewGeneratorFromVariant(variant string, source MigrationSource, destinationPath string) (Generator, error) {
	parser, err := getDatabaseParser(variant)
	if err != nil {
		return Generator{}, err
	}

	destinationDir := filepath.Dir(destinationPath)
	if err = os.MkdirAll(destinationDir, 0755); err != nil {
		return Generator{}, fmt.Errorf("Failed to create directory '%s': %w", destinationDir, err)
	}

	return Generator{
		parser:          parser,
		source:          source,
		destinationPath: destinationPath,
	}, nil
}

type Generator struct {
	parser          databaseParser
	source          MigrationSource
	destinationPath string
	schemasFilter   []string
}

func (gen *Generator) WithSchemaFilter(schemas ...string) *Generator {
	gen.schemasFilter = append(gen.schemasFilter, schemas...)
	return gen
}

func (gen Generator) Run() error {
	dbdefinition, err := gen.getDatabaseDefinition()
	if err != nil {
		return err
	}

	// TODO: add proper template
	gogenerator, err := gogen.NewGoGenerator(dbdefinition, []string{})
	if err != nil {
		return err
	}
	gogenerator.WithSchemaFilter(gen.schemasFilter...)

	goFileContent, err := gogenerator.Run()
	if err != nil {
		return err
	}
	defer goFileContent.Close()

	err = gen.writeGoFile(goFileContent)
	if err != nil {
		return err
	}

	return nil
}

func (gen Generator) writeGoFile(content io.Reader) error {
	file, err := os.Create(gen.destinationPath)
	if err != nil {
		return fmt.Errorf("failed to create file '%s': %w", gen.destinationPath, err)
	}
	defer file.Close()

	_, err = io.Copy(file, content)
	if err != nil {
		return fmt.Errorf("failed to write go file content: %w", err)
	}

	return nil
}

func (gen *Generator) getDatabaseDefinition() (database.Definition, error) {
	for {
		reader, err := gen.source.Next()
		if errors.Is(err, ErrEndOfIteration) {
			break
		}
		defer reader.Close()

		rawSQL, err := ioutil.ReadAll(reader)
		if err != nil {
			return database.Definition{}, fmt.Errorf("failed to read SQL: %w", err)
		}

		err = gen.parser.ParseSQL(string(rawSQL))
		if err != nil {
			return database.Definition{}, fmt.Errorf("failed to parse SQL: %w", err)
		}
	}

	return gen.parser.GetResultingSchemaDefinition(), nil
}

func getDatabaseParser(variant string) (databaseParser, error) {
	if variant == SQL_VARIANT_POSTGRES {
		return postgres.NewDDLParser(), nil
	}
	return nil, fmt.Errorf("database variant '%s' is unknown", variant)
}
