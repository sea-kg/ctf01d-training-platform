package models

type Result struct {
	Id     int    `db:"id"`
	TeamId string `db:"team_id"`
	GameId string `db:"game_id"`
	Score  int32  `db:"score"`
	Rank   int32  `db:"rank"`
}
