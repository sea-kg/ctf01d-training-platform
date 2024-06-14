package view

import (
	"ctf01d/internal/app/db"
	"ctf01d/internal/app/server"
)

func NewUniversityFromModel(u *db.University) *server.UniversityResponse {
	return &server.UniversityResponse{
		Id:   u.Id,
		Name: u.Name,
	}
}

func NewUniversitiesFromModels(ms []*db.University) []*server.UniversityResponse {
	var universities []*server.UniversityResponse = []*server.UniversityResponse{}
	for _, m := range ms {
		universities = append(universities, NewUniversityFromModel(m))
	}
	return universities
}
