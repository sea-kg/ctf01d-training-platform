package db

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type Profile struct {
	Id          openapi_types.UUID `db:"id"`
	CurrentTeam string             `db:"name"`
	CreatedAt   time.Time          `db:"created_at"`
	UpdatedAt   time.Time          `db:"created_at"`
	Role        string             `db:"role"`
}

type ProfileTeams struct {
	JoinedAt time.Time  `db:"joined_at"`
	LeftAt   *time.Time `db:"left_at"`
	Role     string     `db:"role"`
	Name     string     `db:"name"`
}

type ProfileWithHistory struct {
	Profile Profile
	History []ProfileTeams
}
