package model

import (
	"database/sql"

	"ctf01d/internal/server"
	helpers "ctf01d/internal/utils"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type Team struct {
	Id           openapi_types.UUID `db:"id"            json:"id"`
	Name         string             `db:"name"          json:"name"`
	Description  string             `db:"description"   json:"description"`
	UniversityId openapi_types.UUID `db:"university_id" json:"university_id"`
	SocialLinks  sql.NullString     `db:"social_links"  json:"social_links"`
	AvatarUrl    sql.NullString     `db:"avatar_url"    json:"avatar_url"`
	University   *string
}

func NewTeamFromModel(t *Team) *server.TeamResponse {
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
		SocialLinks: &t.SocialLinks.String,
		AvatarUrl:   &avatarUrl,
	}
}

func NewTeamsFromModels(ts []*Team) []*server.TeamResponse {
	var teams []*server.TeamResponse
	for _, t := range ts {
		teams = append(teams, NewTeamFromModel(t))
	}
	return teams
}
