package db

import openapi_types "github.com/oapi-codegen/runtime/types"

type Result struct {
	Id     openapi_types.UUID `db:"id"`
	TeamId openapi_types.UUID `db:"team_id"`
	GameId openapi_types.UUID `db:"game_id"`
	Score  float64            `db:"score"`
}
