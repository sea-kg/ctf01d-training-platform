package db

import openapi_types "github.com/oapi-codegen/runtime/types"

type Result struct {
	Id     openapi_types.UUID `db:"id"`
	TeamId string             `db:"team_id"`
	GameId string             `db:"game_id"`
	Score  int                `db:"score"`
	Rank   int                `db:"rank"`
}
