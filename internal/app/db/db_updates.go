package db

import (
	"time"
)

type DatabaseUpdate struct {
	Id          int       `db:"id"`
	StartTime   time.Time `db:"start_time"`
	UpdateId    string    `db:"update_id"`
	Description string    `db:"description"`
}