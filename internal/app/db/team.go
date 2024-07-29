package db

import (
	"database/sql"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type Team struct {
	Id           openapi_types.UUID `db:"id"`
	Name         string             `db:"name"`
	Description  string             `db:"description"`
	UniversityId openapi_types.UUID `db:"university_id"`
	SocialLinks  sql.NullString     `db:"social_links"`
	AvatarUrl    sql.NullString     `db:"avatar_url"`
	University   *string
}
