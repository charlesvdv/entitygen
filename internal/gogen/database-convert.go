package gogen

import (
	"errors"

	"github.com/charlesvdv/entitygen/internal/database"
)

func ConvertTableToEntity(table *database.Table) (Entity, error) {
	entity := NewEntity(goifyIdentifier(table.Name()))

	for i := range table.Columns() {
		field, err := convertColumnToField(table.Columns()[i])
		if err != nil {
			return entity, err
		}
		entity.fields = append(entity.fields, field)
	}

	return entity, errors.New("not implemented")
}

func convertColumnToField(column *database.Column) (EntityField, error) {
	goType, err := convertDatabaseTypeToGoType(column.Type())
	if err != nil {
		return EntityField{}, err
	}
	field := NewEntityField(goifyIdentifier(column.Name()), column.Name(), goType)
	return field, nil
}

func convertDatabaseTypeToGoType(dbType database.Type) (FieldType, error) {
	switch concreteType := dbType.(type) {
	case *database.BuiltinType:
		return convertDatabaseBuiltinTypeToGoType(concreteType)
	default:
		return FieldType{}, errors.New("unknown type")
	}
}

func convertDatabaseBuiltinTypeToGoType(dbBuiltinType *database.BuiltinType) (FieldType, error) {
	var goType FieldType

	switch dbBuiltinType.Name() {
	case database.TypeBoolean:
		goType = NewBuiltinType("bool")
	case database.TypeDate:
		goType = NewTypeFromPackage("time", "Date")
	case database.TypeInteger:
		goType = NewBuiltinType("int")
	case database.TypeString:
		goType = NewBuiltinType("string")
	case database.TypeTime:
		goType = NewTypeFromPackage("time", "Time")
	case database.TypeTimestamp:
		// TODO
		goType = NewBuiltinType("uint64")
	case database.TypeBigInt:
		goType = NewBuiltinType("int64")
	default:
		return FieldType{}, errors.New("unknown database built-in type")
	}

	return goType, nil
}
