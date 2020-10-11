package schema

const (
	TypeBoolean   = "BOOLEAN"
	TypeInteger   = "INTEGER"
	TypeBigInt    = "BIGINT"
	TypeFloat     = "FLOAT"
	TypeDecimal   = "DECIMAL"
	TypeDate      = "DATE"
	TypeTime      = "TIME"
	TypeTimestamp = "TIMESTAMP"
	TypeString    = "STRING"
)

type Type interface {
	_type()
}

func (t *BuiltinType) _type() {}

func NewTypeBoolean() *BuiltinType {
	return newBuiltinType(TypeBoolean)
}

func NewTypeInteger() *BuiltinType {
	return newBuiltinType(TypeInteger)
}

func NewTypeBigInt() *BuiltinType {
	return newBuiltinType(TypeBigInt)
}

func NewTypeFloat() *BuiltinType {
	return newBuiltinType(TypeFloat)
}

func NewTypeDecimal() *BuiltinType {
	return newBuiltinType(TypeDecimal)
}

func NewTypeDate() *BuiltinType {
	return newBuiltinType(TypeDate)
}

func NewTypeTime() *BuiltinType {
	return newBuiltinType(TypeDate)
}

func NewTypeTimestamp() *BuiltinType {
	return newBuiltinType(TypeTimestamp)
}

func NewTypeString() *BuiltinType {
	return newBuiltinType(TypeString)
}

func newBuiltinType(name string) *BuiltinType {
	return &BuiltinType{
		name: name,
	}
}

type BuiltinType struct {
	name string
}

func (t BuiltinType) Name() string {
	return t.name
}
