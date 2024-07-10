package view

import (
	"ctf01d/internal/app/server"

	"ctf01d/internal/app/db"
	helpers "ctf01d/internal/app/utils"
)

func NewSessionFromModel(u *db.User) *server.SessionResponse {
	userRole := helpers.ConvertUserRequestRoleToString(u.Role)
	return &server.SessionResponse{
		Id:   &u.Id,
		Name: &u.Username,
		Role: &userRole,
	}
}
