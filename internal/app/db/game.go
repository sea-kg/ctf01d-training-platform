package db

import (
	"time"
)

type Game struct {
	Id          int       `db:"id"`
	StartTime   time.Time `db:"start_time"`
	EndTime     time.Time `db:"end_time"`
	Description string    `db:"description"`
}

type GameDetails struct {
	Game
	TeamDetails []*TeamDetails
}
