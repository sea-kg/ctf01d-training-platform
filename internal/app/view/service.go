package view

import "ctf01d/internal/app/models"

type Service struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	LogoUrl     string `json:"logo_url,omitempty"`
	Description string `json:"description,omitempty"`
	IsPublic    bool   `json:"is_public"`
}

func NewServiceFromModel(s *models.Service) *Service {
	return &Service{
		Id:          s.Id,
		Name:        s.Name,
		Author:      s.Author,
		LogoUrl:     s.LogoUrl,
		Description: s.Description,
		IsPublic:    s.IsPublic,
	}
}

func NewServiceFromModels(ms []*models.Service) []*Service {
	var services []*Service
	for _, s := range ms {
		services = append(services, NewServiceFromModel(s))
	}
	return services
}
