package gogen

func NewEntity(name string) Entity {
	return Entity{
		name:   name,
		fields: []EntityField{},
	}
}

type Entity struct {
	name   string
	fields []EntityField
}

func NewEntityField(name, columnName string, _type FieldType) EntityField {
	return EntityField{
		name:       name,
		columnName: columnName,
		_type:      _type,
	}
}

type EntityField struct {
	name       string
	columnName string
	_type      FieldType
}

func NewBuiltinType(name string) FieldType {
	return FieldType{
		name:     name,
		_package: "",
	}
}

func NewTypeFromPackage(_package, name string) FieldType {
	return FieldType{
		_package: _package,
		name:     name,
	}
}

type FieldType struct {
	name     string
	_package string
}
