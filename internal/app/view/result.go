package view

import (
	"ctf01d/internal/app/db"
	"ctf01d/internal/app/server"
)

func NewResultFromModel(s *db.Result) *server.ResultResponse {
	return &server.ResultResponse{
		Id:     s.Id,
		GameId: s.GameId,
		Rank:   s.Rank,
		Score:  s.Score,
		TeamId: s.TeamId,
	}
}
