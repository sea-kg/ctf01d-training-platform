package view

import (
	"ctf01d/internal/model"
	"ctf01d/internal/server"

	helpers "ctf01d/internal/utils"
)

func NewSessionFromModel(u *model.User) *server.SessionResponse {
	userRole := helpers.ConvertUserRequestRoleToString(u.Role)
	return &server.SessionResponse{
		Id:   &u.Id,
		Name: &u.Username,
		Role: &userRole,
	}
}
