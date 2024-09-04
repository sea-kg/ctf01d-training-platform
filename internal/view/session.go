package view

import (
	"ctf01d/internal/helper"
	"ctf01d/internal/httpserver"
	"ctf01d/internal/model"
)

func NewSessionFromModel(u *model.User) *httpserver.SessionResponse {
	userRole := helper.ConvertUserRequestRoleToString(u.Role)
	return &httpserver.SessionResponse{
		Id:   &u.Id,
		Name: &u.Username,
		Role: &userRole,
	}
}
