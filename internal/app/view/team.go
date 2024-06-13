package view

import (
	"ctf01d/internal/app/db"
	"ctf01d/internal/app/server"
)

func NewTeamFromModel(t *db.Team) *server.TeamResponse {
	return &server.TeamResponse{
		Id:          t.Id,
		Name:        t.Name,
		Description: &t.Description,
		University:  &t.University,
		SocialLinks: &t.SocialLinks,
		AvatarUrl:   &t.AvatarUrl,
	}
}

func NewTeamsFromModels(ts []*db.Team) []*server.TeamResponse {
	var teams []*server.TeamResponse
	for _, t := range ts {
		teams = append(teams, NewTeamFromModel(t))
	}
	return teams
}
