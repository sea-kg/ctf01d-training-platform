package models

type User struct {
	Id           int    `db:"id"`
	Username     string `db:"user_name"`
	Role         string `db:"role"`
	AvatarUrl    string `db:"avatar_url"`
	Status       string `db:"status"`
	PasswordHash string `db:"password_hash"`
}
