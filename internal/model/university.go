package model

import (
	"ctf01d/internal/server"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type University struct {
	Id   openapi_types.UUID `db:"id"   json:"id"`
	Name string             `db:"name" json:"name"`
}

func (u *University) ToResponse() *server.UniversityResponse {
	return &server.UniversityResponse{
		Id:   u.Id,
		Name: u.Name,
	}
}

func NewUniversitiesFromModels(us []*University) []*server.UniversityResponse {
	var universities []*server.UniversityResponse = []*server.UniversityResponse{}
	for _, u := range us {
		universities = append(universities, u.ToResponse())
	}
	return universities
}
