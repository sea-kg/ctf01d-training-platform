package view

type Service struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	LogoUrl     string `json:"logo_url,omitempty"`
	Description string `json:"description,omitempty"`
	IsPublic    bool   `json:"is_public"`
}
