package models

type Team struct {
	Id           string `db:"id"`
	Name         string `db:"name"`
	Description  string `db:"description"`
	UniversityId string `db:"university_id"`
	SocialLinks  string `db:"social_links"`
	AvatarUrl    string `db:"avatar_url"`
}
