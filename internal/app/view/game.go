package view

import (
	"ctf01d/internal/app/db"
	"time"

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
	Teams       []TeamDetails      `json:"team_details,omitempty"`
}

type TeamDetails struct {
	Id          openapi_types.UUID `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Members     []Member           `json:"members"`
}

type Member struct {
	Id       openapi_types.UUID `json:"id"`
	UserName string             `json:"user_name"`
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
	teams := make([]TeamDetails, 0, len(m.TeamDetails))
	for _, t := range m.TeamDetails {
		members := make([]Member, 0, len(t.Members))
		for _, u := range t.Members {
			members = append(members, Member{
				Id:       u.Id,
				UserName: u.Username,
			})
		}
		teams = append(teams, TeamDetails{
			Id:          t.Id,
			Name:        t.Name,
			Description: t.Description,
			Members:     members,
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
