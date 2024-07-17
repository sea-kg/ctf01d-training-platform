package view

import (
	"ctf01d/internal/app/server"

	"ctf01d/internal/app/db"
	helpers "ctf01d/internal/app/utils"
)

func NewUserFromModel(u *db.User) *server.UserResponse {
	userRole := helpers.ConvertUserRequestRoleToUserResponseRole(u.Role)
	var avatarUrl string
	if u.AvatarUrl.Valid {
		avatarUrl = u.AvatarUrl.String
	} else {
		avatarUrl = helpers.WithDefault(u.Username)
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

func NewUsersFromModels(ms []*db.User) []*server.UserResponse {
	var users []*server.UserResponse
	for _, m := range ms {
		users = append(users, NewUserFromModel(m))
	}
	return users
}
