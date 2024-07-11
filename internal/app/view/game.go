package view

import (
	"time"

	"ctf01d/internal/app/db"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type Game struct {
	Id          openapi_types.UUID `json:"id"`
	StartTime   time.Time          `json:"start_time"`
	EndTime     time.Time          `json:"end_time"`
	Description string             `json:"description,omitempty"`
}

type GameDetails struct {
	Id          openapi_types.UUID `json:"id"`
	StartTime   time.Time          `json:"start_time"`
	EndTime     time.Time          `json:"end_time"`
	Description string             `json:"description,omitempty"`
	Teams       []Teams            `json:"team_details,omitempty"`
}

type Teams struct {
	Id          openapi_types.UUID `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
}

func NewGameFromModel(m *db.Game) *Game {
	return &Game{
		Id:          m.Id,
		StartTime:   m.StartTime,
		EndTime:     m.EndTime,
		Description: m.Description,
	}
}

func NewGameDetailsFromModel(m *db.GameDetails) *GameDetails {
	teams := make([]Teams, 0, len(m.Teams))
	for _, t := range m.Teams {
		teams = append(teams, Teams{
			Id:          t.Id,
			Name:        t.Name,
			Description: t.Description,
		})
	}

	return &GameDetails{
		Id:          m.Id,
		StartTime:   m.StartTime,
		EndTime:     m.EndTime,
		Description: m.Description,
		Teams:       teams,
	}
}

func NewGamesFromModels(ms []*db.Game) []*Game {
	var games []*Game
	for _, m := range ms {
		games = append(games, NewGameFromModel(m))
	}
	return games
}
