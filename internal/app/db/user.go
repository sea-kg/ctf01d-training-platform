package db

import "ctf01d/internal/app/server"

type User struct {
	Id           int                    `db:"id"`
	DisplayName  string                 `db:"display_name"`
	Username     string                 `db:"user_name"`
	Role         server.UserRequestRole `db:"role"`
	AvatarUrl    string                 `db:"avatar_url"`
	Status       string                 `db:"status"`
	PasswordHash string                 `db:"password_hash"`
}
