package model

import (
	"ctf01d/internal/server"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type University struct {
	Id   openapi_types.UUID `db:"id"   json:"id"`
	Name string             `db:"name" json:"name"`
}

func NewUniversityFromModel(u *University) *server.UniversityResponse {
	return &server.UniversityResponse{
		Id:   u.Id,
		Name: u.Name,
	}
}

func NewUniversitiesFromModels(ms []*University) []*server.UniversityResponse {
	var universities []*server.UniversityResponse = []*server.UniversityResponse{}
	for _, m := range ms {
		universities = append(universities, NewUniversityFromModel(m))
	}
	return universities
}
