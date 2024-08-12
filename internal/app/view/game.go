package view

import (
	"ctf01d/internal/app/db"
	"ctf01d/internal/app/server"
)

func NewGameFromModel(m *db.Game) *server.GameResponse {
	return &server.GameResponse{
		Id:          m.Id,
		StartTime:   m.StartTime,
		EndTime:     m.EndTime,
		Description: &m.Description,
	}
}

func NewGameDetailsFromModel(m *db.GameDetails) *server.GameResponse {
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

func NewGamesFromModels(ms []*db.Game) []*server.GameResponse {
	var games []*server.GameResponse
	for _, m := range ms {
		games = append(games, NewGameFromModel(m))
	}
	return games
}
