package db

import openapi_types "github.com/oapi-codegen/runtime/types"

type University struct {
	Id   openapi_types.UUID `db:"id"`
	Name string             `db:"name"`
}
