package model

import (
	openapi_types "github.com/oapi-codegen/runtime/types"

	"ctf01d/internal/server"
)

type Result struct {
	Id     openapi_types.UUID `db:"id"      json:"id"`
	TeamId openapi_types.UUID `db:"team_id" json:"team_id"`
	GameId openapi_types.UUID `db:"game_id" json:"game_id"`
	Score  float64            `db:"score"   json:"score"`
}

func (s *Result) ToResponse(rank int) *server.ResultResponse {
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
		results = append(results, r.ToResponse(i+1))
	}
	return results
}
