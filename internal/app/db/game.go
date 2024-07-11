package db

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type Game struct {
	Id          openapi_types.UUID `db:"id"`
	StartTime   time.Time          `db:"start_time"`
	EndTime     time.Time          `db:"end_time"`
	Description string             `db:"description"`
}

type GameDetails struct {
	Game
	Teams []*Team
}
