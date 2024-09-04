package model

import (
	"time"

	"ctf01d/internal/httpserver"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type Game struct {
	Id          openapi_types.UUID `db:"id"          json:"id"`
	StartTime   time.Time          `db:"start_time"  json:"start_time"`
	EndTime     time.Time          `db:"end_time"    json:"end_time"`
	Description string             `db:"description" json:"description"`
}

type GameDetails struct {
	Game
	Teams []*Team
}

func (m *Game) ToResponse() *httpserver.GameResponse {
	return &httpserver.GameResponse{
		Id:          m.Id,
		StartTime:   m.StartTime,
		EndTime:     m.EndTime,
		Description: &m.Description,
	}
}

func (m *GameDetails) ToResponseGameDetails() *httpserver.GameResponse {
	teams := make([]httpserver.TeamResponse, 0, len(m.Teams))
	for _, t := range m.Teams {
		teams = append(teams, httpserver.TeamResponse{
			Id:          t.Id,
			Name:        t.Name,
			Description: &t.Description,
		})
	}

	return &httpserver.GameResponse{
		Id:          m.Id,
		StartTime:   m.StartTime,
		EndTime:     m.EndTime,
		Description: &m.Description,
		Teams:       &teams,
	}
}

func NewGamesFromModels(ms []*Game) []*httpserver.GameResponse {
	var games []*httpserver.GameResponse
	for _, m := range ms {
		games = append(games, m.ToResponse())
	}
	return games
}
