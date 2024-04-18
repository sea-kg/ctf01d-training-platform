package api

import (
	"ctf01d/lib/models"
	"ctf01d/lib/repository"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type RequestGame struct {
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Description string    `json:"description"`
}

func CreateGameHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var game RequestGame
	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	if game.EndTime.Before(game.StartTime) {
		respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "EndTime must be after StartTime"})
		return
	}
	gameRepo := repository.NewGameRepository(db)
	newGame := &models.Game{
		StartTime:   game.StartTime,
		EndTime:     game.EndTime,
		Description: game.Description,
	}
	if err := gameRepo.Create(r.Context(), newGame); err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create game"})
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"data": "Game created successfully"})
}

func DeleteGameHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	gameRepo := repository.NewGameRepository(db)
	if err := gameRepo.Delete(r.Context(), id); err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete game"})
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"data": "Game deleted successfully"})
}

func GetGameByIdHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	gameRepo := repository.NewGameRepository(db)
	game, err := gameRepo.GetById(r.Context(), id)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch game"})
		return
	}
	gameJSON, err := json.Marshal(game)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	respondWithJSON(w, http.StatusOK, gameJSON)
}

func ListGamesHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	gameRepo := repository.NewGameRepository(db)
	games, err := gameRepo.List(r.Context())
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	gamesJSON, err := json.Marshal(games)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	respondWithJSON(w, http.StatusOK, gamesJSON)
}

func UpdateGameHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var game RequestGame
	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	gameRepo := repository.NewGameRepository(db)
	updateGame := &models.Game{
		StartTime:   game.StartTime,
		EndTime:     game.EndTime,
		Description: game.Description,
	}
	vars := mux.Vars(r)
	id := vars["id"]
	updateGame.Id = id
	err := gameRepo.Update(r.Context(), updateGame)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"data": "Game updated successfully"})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(`{"error": "Error marshaling the response object"}`)); err != nil {
			log.Printf("Error writing error response: %v", err)
		}
		return
	}
	w.WriteHeader(code)
	if _, err := w.Write(response); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
