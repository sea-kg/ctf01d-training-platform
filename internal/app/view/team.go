package view

import (
	"ctf01d/internal/app/db"
	"ctf01d/internal/app/server"
	helpers "ctf01d/internal/app/utils"
)

func NewTeamFromModel(t *db.Team) *server.TeamResponse {
	var avatarUrl string
	if t.AvatarUrl.Valid {
		avatarUrl = t.AvatarUrl.String
	} else {
		avatarUrl = helpers.WithDefault(t.Name)
	}
	return &server.TeamResponse{
		Id:          t.Id,
		Name:        t.Name,
		Description: &t.Description,
		University:  t.University,
		SocialLinks: &t.SocialLinks,
		AvatarUrl:   &avatarUrl,
	}
}

func NewTeamsFromModels(ts []*db.Team) []*server.TeamResponse {
	var teams []*server.TeamResponse
	for _, t := range ts {
		teams = append(teams, NewTeamFromModel(t))
	}
	return teams
}
