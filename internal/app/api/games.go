package api

import (
	apimodels "ctf01d/internal/app/apimodels"
	dbmodel "ctf01d/internal/app/db"
	"ctf01d/internal/app/repository"
	api_helpers "ctf01d/internal/app/utils"
	"ctf01d/internal/app/view"
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateGameHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var game apimodels.GameRequest
	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		slog.Warn(err.Error(), "handler", "CreateGameHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	if game.EndTime.Before(game.StartTime) {
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "EndTime must be after StartTime"})
		return
	}
	repo := repository.NewGameRepository(db)
	newGame := &dbmodel.Game{
		StartTime:   game.StartTime,
		EndTime:     game.EndTime,
		Description: *game.Description,
	}
	if err := repo.Create(r.Context(), newGame); err != nil {
		slog.Warn(err.Error(), "handler", "CreateGameHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create game"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Game created successfully"})
}

func DeleteGameHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		slog.Warn(err.Error(), "handler", "DeleteGameHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Bad request"})
		return
	}
	repo := repository.NewGameRepository(db)
	if err := repo.Delete(r.Context(), id); err != nil {
		slog.Warn(err.Error(), "handler", "DeleteGameHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete game"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Game deleted successfully"})
}

func GetGameByIdHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetGameByIdHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Bad request"})
		return
	}
	repo := repository.NewGameRepository(db)
	game, err := repo.GetGameDetails(r.Context(), id) // короткий ответ, если нужен см. GetById
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetGameByIdHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch game"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewGameDetailsFromModel(game))
}

func ListGamesHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	repo := repository.NewGameRepository(db)
	games, err := repo.List(r.Context())
	if err != nil {
		slog.Warn(err.Error(), "handler", "ListGamesHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Failed to fetch games"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewGamesFromModels(games))
}

func UpdateGameHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// fixme update не проверяет есть ли запись в бд
	var game apimodels.GameRequest
	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateGameHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	repo := repository.NewGameRepository(db)
	updateGame := &dbmodel.Game{
		StartTime:   game.StartTime,
		EndTime:     game.EndTime,
		Description: *game.Description,
	}
	vars := mux.Vars(r)
	id, err2 := strconv.Atoi(vars["id"])
	if err2 != nil {
		slog.Warn(err2.Error(), "handler", "UpdateGameHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Invalid request payload"})
		return
	}
	updateGame.Id = id
	err := repo.Update(r.Context(), updateGame)
	if err != nil {
		slog.Warn(err.Error(), "handler", "UpdateGameHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Invalid request payload"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Game updated successfully"})
}
