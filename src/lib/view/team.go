package view

type Team struct {
	// Unique identifier for the team
	Id string `json:"id"`
	// Name of the team
	TeamName string `json:"team_name"`
	// A brief description of the team
	Description string `json:"description,omitempty"`
	// University or institution the team is associated with
	University string `json:"university,omitempty"`
	// JSON string containing social media links of the team
	SocialLinks string `json:"social_links,omitempty"`
	// URL to the team's avatar
	AvatarUrl string `json:"avatar_url,omitempty"`
}
