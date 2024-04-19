package models

type Team struct {
	Id          string `db:"id"`
	TeamName    string `db:"team_name"`
	Description string `db:"description"`
	University  string `db:"university"`
	SocialLinks string `db:"social_links"`
	AvatarUrl   string `db:"avatar_url"`
}
