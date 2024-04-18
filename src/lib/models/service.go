package models

type Service struct {
	Id          string `db:"id"`
	Name        string `db:"name"`
	Author      string `db:"author"`
	LogoUrl     string `db:"logo_url"`
	Description string `db:"description"`
	IsPublic    bool   `db:"is_public"`
}
