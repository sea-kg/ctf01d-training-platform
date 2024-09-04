package model

import (
	"database/sql"

	"ctf01d/internal/helper"
	"ctf01d/internal/server"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type User struct {
	Id           openapi_types.UUID     `db:"id"            json:"id"`
	DisplayName  sql.NullString         `db:"display_name"  json:"display_name"`
	Username     string                 `db:"user_name"     json:"user_name"`
	Role         server.UserRequestRole `db:"role"          json:"role"`
	AvatarUrl    sql.NullString         `db:"avatar_url"    json:"avatar_url"`
	Status       string                 `db:"status"        json:"status"`
	PasswordHash string                 `db:"password_hash" json:"password_hash"`
}

func (u *User) ToResponse() *server.UserResponse {
	userRole := helper.ConvertUserRequestRoleToUserResponseRole(u.Role)
	var avatarUrl string
	if u.AvatarUrl.Valid {
		avatarUrl = u.AvatarUrl.String
	} else {
		avatarUrl = helper.WithDefault(u.Username)
	}
	return &server.UserResponse{
		Id:          &u.Id,
		UserName:    &u.Username,
		DisplayName: &u.DisplayName.String,
		Role:        &userRole,
		AvatarUrl:   &avatarUrl,
		Status:      &u.Status,
	}
}

func NewUsersFromModels(us []*User) []*server.UserResponse {
	var users []*server.UserResponse
	for _, u := range us {
		users = append(users, u.ToResponse())
	}
	return users
}
