package db

import openapi_types "github.com/oapi-codegen/runtime/types"

type Team struct {
	Id           openapi_types.UUID `db:"id"`
	Name         string             `db:"name"`
	Description  string             `db:"description"`
	UniversityId int                `db:"university_id"`
	University   string
	SocialLinks  string `db:"social_links"`
	AvatarUrl    string `db:"avatar_url"`
}

type TeamDetails struct {
	Team
	Members []*User
}
