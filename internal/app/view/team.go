package view

type Team struct {
	Id          string `json:"id"`
	TeamName    string `json:"team_name"`
	Description string `json:"description,omitempty"`
	University  string `json:"university,omitempty"`
	SocialLinks string `json:"social_links,omitempty"`
	AvatarUrl   string `json:"avatar_url,omitempty"`
}
