package view

import (
	apimodels "ctf01d/internal/app/apimodels"
	"ctf01d/internal/app/db"
)

func NewResultFromModel(s *db.Result) *apimodels.ResultResponse {
	return &apimodels.ResultResponse{
		Id:     s.Id,
		GameId: s.GameId,
		Rank:   s.Rank,
		Score:  s.Score,
		TeamId: s.TeamId,
	}
}

func NewResultFromModels(rm []*db.Result) []*apimodels.ResultResponse {
	var services []*apimodels.ResultResponse
	for _, s := range rm {
		services = append(services, NewResultFromModel(s))
	}
	return services
}
