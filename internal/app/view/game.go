package view

import (
	"ctf01d/internal/app/models"
	"time"
)

type Game struct {
	Id          int       `json:"id"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Description string    `json:"description,omitempty"`
}

type GameDetails struct {
	Id          int           `json:"id"`
	StartTime   time.Time     `json:"start_time"`
	EndTime     time.Time     `json:"end_time"`
	Description string        `json:"description,omitempty"`
	Teams       []TeamDetails `json:"team_details,omitempty"`
}

type TeamDetails struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Members     []Member `json:"members"`
}

type Member struct {
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
}

func NewGameFromModel(m *models.Game) *Game {
	return &Game{
		Id:          m.Id,
		StartTime:   m.StartTime,
		EndTime:     m.EndTime,
		Description: m.Description,
	}
}

func NewGameDetailsFromModel(m *models.GameDetails) *GameDetails {
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
func NewGamesFromModels(ms []*models.Game) []*Game {
	var games []*Game
	for _, m := range ms {
		games = append(games, NewGameFromModel(m))
	}
	return games
}
