package model

import (
	"time"

	"ctf01d/internal/server"

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

func NewGameFromModel(m *Game) *server.GameResponse {
	return &server.GameResponse{
		Id:          m.Id,
		StartTime:   m.StartTime,
		EndTime:     m.EndTime,
		Description: &m.Description,
	}
}

func NewGameDetailsFromModel(m *GameDetails) *server.GameResponse {
	teams := make([]server.TeamResponse, 0, len(m.Teams))
	for _, t := range m.Teams {
		teams = append(teams, server.TeamResponse{
			Id:          t.Id,
			Name:        t.Name,
			Description: &t.Description,
		})
	}

	return &server.GameResponse{
		Id:          m.Id,
		StartTime:   m.StartTime,
		EndTime:     m.EndTime,
		Description: &m.Description,
		Teams:       &teams,
	}
}

func NewGamesFromModels(ms []*Game) []*server.GameResponse {
	var games []*server.GameResponse
	for _, m := range ms {
		games = append(games, NewGameFromModel(m))
	}
	return games
}
