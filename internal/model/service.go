package model

import (
	"database/sql"

	"ctf01d/internal/helper"
	"ctf01d/internal/httpserver"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type Service struct {
	Id          openapi_types.UUID `db:"id"          json:"id"`
	Name        string             `db:"name"        json:"name"`
	Author      string             `db:"author"      json:"author"`
	LogoUrl     sql.NullString     `db:"logo_url"    json:"logo_url"`
	Description string             `db:"description" json:"description"`
	IsPublic    bool               `db:"is_public"   json:"is_public"`
}

func (s *Service) ToResponse() *httpserver.ServiceResponse {
	var logo string
	if s.LogoUrl.Valid {
		logo = s.LogoUrl.String
	} else {
		logo = helper.WithDefault(s.Name)
	}
	return &httpserver.ServiceResponse{
		Id:          s.Id,
		Name:        s.Name,
		Author:      s.Author,
		LogoUrl:     &logo,
		Description: &s.Description,
		IsPublic:    s.IsPublic,
	}
}

func NewServiceFromModels(ms []*Service) []*httpserver.ServiceResponse {
	var services []*httpserver.ServiceResponse
	for _, s := range ms {
		services = append(services, s.ToResponse())
	}
	return services
}
