package models

type Team struct {
	Id           int    `db:"id"`
	Name         string `db:"name"`
	Description  string `db:"description"`
	UniversityId string `db:"university_id"`
	University   string
	SocialLinks  string `db:"social_links"`
	AvatarUrl    string `db:"avatar_url"`
}

type TeamDetails struct {
	Team
	Members []*User
}
