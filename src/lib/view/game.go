package view

import (
	"time"
)

type Game struct {
	// Unique identifier for the game
	Id string `json:"id"`
	// The start time of the game
	StartTime time.Time `json:"start_time"`
	// The end time of the game
	EndTime time.Time `json:"end_time"`
	// A brief description of the game
	Description string `json:"description,omitempty"`
}
