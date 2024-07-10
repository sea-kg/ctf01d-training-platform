package view

import (
	"ctf01d/internal/app/server"

	"ctf01d/internal/app/db"
)

func NewProfileFromModel(p *db.ProfileWithHistory) *server.ProfileResponse {
	return &server.ProfileResponse{
		Id:          p.Profile.Id,
		CreatedAt:   p.Profile.CreatedAt,
		UpdatedAt:   &p.Profile.UpdatedAt,
		TeamName:    p.Profile.CurrentTeam,
		TeamRole:    server.ProfileResponseTeamRole(p.Profile.Role),
		TeamHistory: makeTeamHistory(p.History),
	}
}

func makeTeamHistory(tms []db.ProfileTeams) *[]server.TeamHistory {
	out := []server.TeamHistory{}
	for _, tm := range tms {
		out = append(out, server.TeamHistory{
			Join: tm.JoinedAt,
			Left: tm.LeftAt,
			Name: tm.Name,
			Role: server.TeamHistoryRole(tm.Role),
		})
	}
	return &out
}
