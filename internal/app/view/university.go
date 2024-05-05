package view

import "ctf01d/internal/app/models"

type University struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func NewUniversityFromModel(u *models.University) *University {
	return &University{
		Id:   u.Id,
		Name: u.Name,
	}
}

func NewUniversitiesFromModels(ms []*models.University) []*University {
	var universities []*University = []*University{}
	for _, m := range ms {
		universities = append(universities, NewUniversityFromModel(m))
	}
	return universities
}
