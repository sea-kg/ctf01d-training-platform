package view

import "ctf01d/internal/app/db"

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"user_name"`
	Role      string `json:"role,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	Status    string `json:"status,omitempty"`
}

func NewUserFromModel(u *db.User) *User {
	return &User{
		Id:        u.Id,
		Username:  u.Username,
		Role:      u.Role,
		AvatarUrl: u.AvatarUrl,
		Status:    u.Status,
	}
}

func NewUsersFromModels(ms []*db.User) []*User {
	var users []*User
	for _, m := range ms {
		users = append(users, NewUserFromModel(m))
	}
	return users
}
