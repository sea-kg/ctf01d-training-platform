package models

import (
	"time"
)

type User struct {
	Id        string    `db:"id"`
	Username  time.Time `db:"username"`
	Role      time.Time `db:"role"`
	AvatarUrl string    `db:"avatar_url"`
	Status    string    `db:"status"`
}
