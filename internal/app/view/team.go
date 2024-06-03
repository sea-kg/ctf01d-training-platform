package view

import (
	apimodels "ctf01d/internal/app/apimodels"
	"ctf01d/internal/app/db"
)

func NewTeamFromModel(t *db.Team) *apimodels.TeamResponse {
	return &apimodels.TeamResponse{
		Id:          t.Id,
		Name:        t.Name,
		Description: &t.Description,
		University:  &t.University,
		SocialLinks: &t.SocialLinks,
		AvatarUrl:   &t.AvatarUrl,
	}
}

func NewTeamsFromModels(ts []*db.Team) []*apimodels.TeamResponse {
	var teams []*apimodels.TeamResponse
	for _, t := range ts {
		teams = append(teams, NewTeamFromModel(t))
	}
	return teams
}
