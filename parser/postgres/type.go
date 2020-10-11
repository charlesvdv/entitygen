package postgres

import (
	"github.com/charlesvdv/entitygen/schema"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
)

func convertType(rawType tree.ResolvableTypeReference) schema.Type {
	switch rawType.SQLString() {
	case "INT4":
		return schema.NewTypeInteger()
	case "INT8":
		return schema.NewTypeBigInt()
	case "STRING":
		return schema.NewTypeString()
	case "BOOL":
		return schema.NewTypeBoolean()
	case "FLOAT8":
		return schema.NewTypeDecimal()
	case "FLOAT4":
		return schema.NewTypeFloat()
	case "TIME":
		return schema.NewTypeTime()
	case "TIMESTAMP":
		return schema.NewTypeTimestamp()
	case "DATE":
		return schema.NewTypeDate()
	default:
		return nil
	}
}
