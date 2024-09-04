package model

import (
	"ctf01d/internal/httpserver"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type University struct {
	Id   openapi_types.UUID `db:"id"   json:"id"`
	Name string             `db:"name" json:"name"`
}

func (u *University) ToResponse() *httpserver.UniversityResponse {
	return &httpserver.UniversityResponse{
		Id:   u.Id,
		Name: u.Name,
	}
}

func NewUniversitiesFromModels(us []*University) []*httpserver.UniversityResponse {
	var universities []*httpserver.UniversityResponse = []*httpserver.UniversityResponse{}
	for _, u := range us {
		universities = append(universities, u.ToResponse())
	}
	return universities
}
