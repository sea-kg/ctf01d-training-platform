package view

import "ctf01d/internal/app/models"

type Team struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	University  string `json:"university,omitempty"`
	SocialLinks string `json:"social_links,omitempty"`
	AvatarUrl   string `json:"avatar_url,omitempty"`
}

func NewTeamFromModel(t *models.Team) *Team {
	return &Team{
		Id:          t.Id,
		Name:        t.Name,
		Description: t.Description,
		University:  t.University,
		SocialLinks: t.SocialLinks,
		AvatarUrl:   t.AvatarUrl,
	}
}

func NewTeamsFromModels(ts []*models.Team) []*Team {
	var teams []*Team
	for _, t := range ts {
		teams = append(teams, NewTeamFromModel(t))
	}
	return teams
}
