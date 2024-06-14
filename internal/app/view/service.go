package view

import (
	"ctf01d/internal/app/db"
	"ctf01d/internal/app/server"
)

func NewServiceFromModel(s *db.Service) *server.ServiceResponse {
	return &server.ServiceResponse{
		Id:          s.Id,
		Name:        s.Name,
		Author:      s.Author,
		LogoUrl:     &s.LogoUrl,
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
