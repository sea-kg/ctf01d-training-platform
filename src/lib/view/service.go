package view

type Service struct {
	// Unique identifier for the service
	Id string `json:"id"`
	// Name of the service
	Name string `json:"name"`
	// Author of the service
	Author string `json:"author"`
	// URL to the logo of the service
	LogoUrl string `json:"logo_url,omitempty"`
	// A brief description of the service
	Description string `json:"description,omitempty"`
	// Boolean indicating if the service is public
	IsPublic bool `json:"is_public"`
}
