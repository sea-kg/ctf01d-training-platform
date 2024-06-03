package view

import (
	apimodels "ctf01d/internal/app/apimodels"
	"ctf01d/internal/app/db"
)

func NewServiceFromModel(s *db.Service) *apimodels.ServiceResponse {
	return &apimodels.ServiceResponse{
		Id:          s.Id,
		Name:        s.Name,
		Author:      s.Author,
		LogoUrl:     &s.LogoUrl,
		Description: &s.Description,
		IsPublic:    s.IsPublic,
	}
}

func NewServiceFromModels(ms []*db.Service) []*apimodels.ServiceResponse {
	var services []*apimodels.ServiceResponse
	for _, s := range ms {
		services = append(services, NewServiceFromModel(s))
	}
	return services
}
