package view

import (
	"ctf01d/internal/app/server"

	"ctf01d/internal/app/db"
	helpers "ctf01d/internal/app/utils"
)

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"user_name"`
	Role      string `json:"role,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	Status    string `json:"status,omitempty"`
}

func NewUserFromModel(u *db.User) *server.UserResponse {
	userRole := helpers.ConvertUserRequestRoleToUserResponseRole(u.Role)
	return &server.UserResponse{
		Id:        &u.Id,
		UserName:  &u.Username,
		Role:      &userRole,
		AvatarUrl: &u.AvatarUrl,
		Status:    &u.Status,
	}
}

func NewUsersFromModels(ms []*db.User) []*server.UserResponse {
	var users []*server.UserResponse
	for _, m := range ms {
		users = append(users, NewUserFromModel(m))
	}
	return users
}
