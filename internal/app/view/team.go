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

func NewTeamFromModel(u *models.Team) *Team {
	return &Team{
		Id:          u.Id,
		Name:        u.Name,
		Description: u.Description,
		University:  u.UniversityId,
		SocialLinks: u.SocialLinks,
		AvatarUrl:   u.AvatarUrl,
	}
}

func NewTeamsFromModels(ms []*models.Team) []*Team {
	var teams []*Team
	for _, m := range ms {
		teams = append(teams, NewTeamFromModel(m))
	}
	return teams
}
