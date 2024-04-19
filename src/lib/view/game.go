package view

import (
	"ctf01d/lib/models"
	"time"
)

type Game struct {
	Id          string    `json:"id"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Description string    `json:"description,omitempty"`
}

func NewGameFromModel(m *models.Game) *Game {
	return &Game{
		Id:          m.Id,
		StartTime:   m.StartTime,
		EndTime:     m.EndTime,
		Description: m.Description,
	}
}

func NewGamesFromModels(ms []*models.Game) []*Game {
	var games []*Game
	for _, m := range ms {
		games = append(games, NewGameFromModel(m))
	}
	return games
}
