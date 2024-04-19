package models

type User struct {
	Id        string `db:"id"`
	Username  string `db:"username"`
	Role      string `db:"role"`
	AvatarUrl string `db:"avatar_url"`
	Status    string `db:"status"`
}
