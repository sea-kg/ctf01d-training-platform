package db

import openapi_types "github.com/oapi-codegen/runtime/types"

type Service struct {
	Id          openapi_types.UUID `db:"id"`
	Name        string             `db:"name"`
	Author      string             `db:"author"`
	LogoUrl     string             `db:"logo_url"`
	Description string             `db:"description"`
	IsPublic    bool               `db:"is_public"`
}
