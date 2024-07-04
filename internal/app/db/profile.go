package db

import (
	"time"
)

type Profile struct {
	CurrentTeam string    `db:"name"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"created_at"`
}

type ProfileTeams struct {
	JoinedAt time.Time  `db:"joined_at"`
	LeftAt   *time.Time `db:"left_at"`
	Name     string     `db:"name"`
}

type ProfileWithHistory struct {
	Profile Profile
	History []ProfileTeams
}
