package view

type User struct {
	Id        string `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Role      string `json:"role,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	Status    string `json:"status,omitempty"`
}
