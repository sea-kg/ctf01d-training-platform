package db

import (
	"ctf01d/internal/app/server"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type User struct {
	Id           openapi_types.UUID     `db:"id"`
	DisplayName  string                 `db:"display_name"`
	Username     string                 `db:"user_name"`
	Role         server.UserRequestRole `db:"role"`
	AvatarUrl    string                 `db:"avatar_url"`
	Status       string                 `db:"status"`
	PasswordHash string                 `db:"password_hash"`
}
