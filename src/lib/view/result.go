package view

type Result struct {
	Id     string `json:"id"`
	TeamId string `json:"team_id"`
	GameId string `json:"game_id"`
	Score  int32  `json:"score"`
	Rank   int32  `json:"rank"`
}
