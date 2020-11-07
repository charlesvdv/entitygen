package postgres

import (
	"github.com/charlesvdv/entitygen/database"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
)

func convertType(rawType tree.ResolvableTypeReference) database.Type {
	switch rawType.SQLString() {
	case "INT4":
		return database.NewTypeInteger()
	case "INT8":
		return database.NewTypeBigInt()
	case "STRING":
		return database.NewTypeString()
	case "BOOL":
		return database.NewTypeBoolean()
	case "FLOAT8":
		return database.NewTypeDecimal()
	case "FLOAT4":
		return database.NewTypeFloat()
	case "TIME":
		return database.NewTypeTime()
	case "TIMESTAMP":
		return database.NewTypeTimestamp()
	case "DATE":
		return database.NewTypeDate()
	default:
		return nil
	}
}
