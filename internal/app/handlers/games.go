package handlers

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"

	dbmodel "ctf01d/internal/app/db"
	"ctf01d/internal/app/repository"
	"ctf01d/internal/app/server"
	api_helpers "ctf01d/internal/app/utils"
	"ctf01d/internal/app/view"
)

type Handlers struct {
	DB *sql.DB
}

func (h *Handlers) CreateGame(w http.ResponseWriter, r *http.Request) {
	var game server.GameRequest
	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		slog.Warn(err.Error(), "handler", "CreateGame")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	if game.EndTime.Before(game.StartTime) {
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "EndTime must be after StartTime"})
		return
	}
	repo := repository.NewGameRepository(h.DB)
	newGame := &dbmodel.Game{
		StartTime:   game.StartTime,
		EndTime:     game.EndTime,
		Description: *game.Description,
	}
	if err := repo.Create(r.Context(), newGame); err != nil {
		slog.Warn(err.Error(), "handler", "CreateGame")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create game"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Game created successfully"})
}

func (h *Handlers) DeleteGame(w http.ResponseWriter, r *http.Request, id int) {
	repo := repository.NewGameRepository(h.DB)
	if err := repo.Delete(r.Context(), id); err != nil {
		slog.Warn(err.Error(), "handler", "DeleteGame")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete game"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Game deleted successfully"})
}

func (h *Handlers) GetGameById(w http.ResponseWriter, r *http.Request, id int) {
	repo := repository.NewGameRepository(h.DB)
	game, err := repo.GetGameDetails(r.Context(), id) // короткий ответ, если нужен см. GetById
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetGameById")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch game"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewGameDetailsFromModel(game))
}

func (h *Handlers) ListGames(w http.ResponseWriter, r *http.Request) {
	repo := repository.NewGameRepository(h.DB)
	games, err := repo.List(r.Context())
	if err != nil {
		slog.Warn(err.Error(), "handler", "ListGames")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Failed to fetch games"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewGamesFromModels(games))
}

func (h *Handlers) UpdateGame(w http.ResponseWriter, r *http.Request, id int) {
	// fixme update не проверяет есть ли запись в бд
	var game server.GameRequest
	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateGame")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	repo := repository.NewGameRepository(h.DB)
	updateGame := &dbmodel.Game{
		StartTime:   game.StartTime,
		EndTime:     game.EndTime,
		Description: *game.Description,
	}
	updateGame.Id = id
	err := repo.Update(r.Context(), updateGame)
	if err != nil {
		slog.Warn(err.Error(), "handler", "UpdateGame")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Invalid request payload"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Game updated successfully"})
}
