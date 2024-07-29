package view

import (
	"ctf01d/internal/app/db"
	"ctf01d/internal/app/server"
)

func NewResultFromModel(s *db.Result, rank int) *server.ResultResponse {
	return &server.ResultResponse{
		Id:     s.Id,
		GameId: s.GameId,
		Rank:   rank,
		Score:  s.Score,
		TeamId: s.TeamId,
	}
}

func NewScoreboardFromResults(ms []*db.Result) []*server.ResultResponse {
	var results []*server.ResultResponse
	for i, r := range ms {
		results = append(results, NewResultFromModel(r, i+1))
	}
	return results
}
