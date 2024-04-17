package api

import (
	"context"
	"ctf01d/lib/models"
	"ctf01d/lib/repository"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type RequestGame struct {
	StartTime   time.Time `json"start_time"`  // "2000-01-23T04:56:07.000Z",
	EndTime     time.Time `json"end_time"`    // "2000-01-23T04:56:07.000Z",
	Description string    `json"description"` // "description",
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var game RequestGame
	err := json.NewDecoder(r.Body).Decode(&game)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println(game)
	/// move it
	db, err := sql.Open("mysql", "service2_go:service2_go@tcp(localhost:3306)/service2_go")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	///
	gameRepo := repository.NewGameRepository(db)
	newGame := &models.Game{
		StartTime:   game.StartTime,
		EndTime:     game.EndTime,
		Description: game.Description,
	}
	ctx := context.Background()
	err = gameRepo.Create(ctx, newGame)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetGameById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func ListGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UpdateGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
