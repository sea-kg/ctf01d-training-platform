package view

import "ctf01d/internal/app/models"

type University struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewUniversityFromModel(u *models.University) *University {
	return &University{
		Id:   u.Id,
		Name: u.Name,
	}
}

func NewUniversitiesFromModels(ms []*models.University) []*University {
	var universitys []*University
	for _, m := range ms {
		universitys = append(universitys, NewUniversityFromModel(m))
	}
	return universitys
}
