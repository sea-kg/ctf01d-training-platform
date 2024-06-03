package view

import (
	apimodels "ctf01d/internal/app/apimodels"
	"ctf01d/internal/app/db"
)

func NewUniversityFromModel(u *db.University) *apimodels.UniversityResponse {
	return &apimodels.UniversityResponse{
		Id:   u.Id,
		Name: u.Name,
	}
}

func NewUniversitiesFromModels(ms []*db.University) []*apimodels.UniversityResponse {
	var universities []*apimodels.UniversityResponse = []*apimodels.UniversityResponse{}
	for _, m := range ms {
		universities = append(universities, NewUniversityFromModel(m))
	}
	return universities
}
