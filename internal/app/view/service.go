package view

import (
	"ctf01d/internal/app/db"
	"ctf01d/internal/app/server"
	helpers "ctf01d/internal/app/utils"
)

func NewServiceFromModel(s *db.Service) *server.ServiceResponse {
	var logo string
	if s.LogoUrl.Valid {
		logo = s.LogoUrl.String
	} else {
		logo = helpers.WithDefault(s.Name)
	}
	return &server.ServiceResponse{
		Id:          s.Id,
		Name:        s.Name,
		Author:      s.Author,
		LogoUrl:     &logo,
		Description: &s.Description,
		IsPublic:    s.IsPublic,
	}
}

func NewServiceFromModels(ms []*db.Service) []*server.ServiceResponse {
	var services []*server.ServiceResponse
	for _, s := range ms {
		services = append(services, NewServiceFromModel(s))
	}
	return services
}
