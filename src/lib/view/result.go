package view

type Result struct {
	// Unique identifier for the result entry
	Id string `json:"id"`
	// Identifier of the team this result belongs to
	TeamId string `json:"team_id"`
	// Identifier of the game this result is for
	GameId string `json:"game_id"`
	// The score achieved by the team
	Score int32 `json:"score"`
	// The rank achieved by the team in this game
	Rank int32 `json:"rank"`
}
