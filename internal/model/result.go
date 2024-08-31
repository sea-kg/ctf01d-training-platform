package model

import (
	"ctf01d/internal/server"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type Result struct {
	Id     openapi_types.UUID `db:"id"      json:"id"`
	TeamId openapi_types.UUID `db:"team_id" json:"team_id"`
	GameId openapi_types.UUID `db:"game_id" json:"game_id"`
	Score  float64            `db:"score"   json:"score"`
}

func NewResultFromModel(s *Result, rank int) *server.ResultResponse {
	return &server.ResultResponse{
		Id:     s.Id,
		GameId: s.GameId,
		Rank:   rank,
		Score:  s.Score,
		TeamId: s.TeamId,
	}
}

func NewScoreboardFromResults(ms []*Result) []*server.ResultResponse {
	var results []*server.ResultResponse
	for i, r := range ms {
		results = append(results, NewResultFromModel(r, i+1))
	}
	return results
}
