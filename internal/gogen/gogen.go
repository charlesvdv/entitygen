package gogen

import (
	"errors"
	"io"
	"text/template"

	"github.com/charlesvdv/entitygen/internal/database"
)

func NewGoGenerator(databaseDefinition database.Definition, templateFiles []string) (GoGenerator, error) {
	template, err := template.ParseFiles(templateFiles...)
	if err != nil {
		return GoGenerator{}, err
	}

	return GoGenerator{
		databaseDefinition: databaseDefinition,
		template:           template,
	}, nil
}

type GoGenerator struct {
	databaseDefinition database.Definition
	template           *template.Template
	schemas            []string
}

func (gen *GoGenerator) WithSchemaFilter(schemas ...string) *GoGenerator {
	gen.schemas = append(gen.schemas, schemas...)
	return gen
}

func (gen *GoGenerator) Run() (io.ReadCloser, error) {
	return nil, errors.New("not implemented yet")
}
