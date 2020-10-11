package postgres

import (
	"github.com/charlesvdv/entitygen/schema"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
)

func convertType(rawType tree.ResolvableTypeReference) schema.Type {
	switch rawType.SQLString() {
	case "INT8":
		return schema.NewTypeBigInt()
	case "STRING":
		return schema.NewTypeString()
	default:
		return nil
	}
}
