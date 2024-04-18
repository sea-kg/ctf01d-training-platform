package view

type User struct {
	// The unique identifier for the user
	Id string `json:"id,omitempty"`
	// The name of the user
	Username string `json:"username,omitempty"`
	// The role of the user (admin, player)
	Role string `json:"role,omitempty"`
	// URL to the user's avatar
	AvatarUrl string `json:"avatar_url,omitempty"`
	// Status of the user (active, disabled)
	Status string `json:"status,omitempty"`
}
