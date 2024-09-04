package view

import (
	"ctf01d/internal/helper"
	"ctf01d/internal/model"
	"ctf01d/internal/server"
)

func NewSessionFromModel(u *model.User) *server.SessionResponse {
	userRole := helper.ConvertUserRequestRoleToString(u.Role)
	return &server.SessionResponse{
		Id:   &u.Id,
		Name: &u.Username,
		Role: &userRole,
	}
}
