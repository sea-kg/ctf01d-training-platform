package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	openapi_types "github.com/oapi-codegen/runtime/types"

	"ctf01d/internal/helper"
	"ctf01d/internal/model"
	"ctf01d/internal/repository"
	"ctf01d/internal/server"
)

func (h *Handler) CreateGame(w http.ResponseWriter, r *http.Request) {
	var game server.GameRequest
	var err error
	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		slog.Warn(err.Error(), "handler", "CreateGame")
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	if game.EndTime.Before(game.StartTime) {
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "EndTime must be after StartTime"})
		return
	}
	repo := repository.NewGameRepository(h.DB)
	newGame := &model.Game{
		StartTime:   game.StartTime,
		EndTime:     game.EndTime,
		Description: *game.Description,
	}

	err = repo.Create(r.Context(), newGame)
	if err != nil {
		slog.Warn(err.Error(), "handler", "CreateGame")
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create game"})
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, newGame.ToResponse())
}

func (h *Handler) DeleteGame(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	repo := repository.NewGameRepository(h.DB)
	if err := repo.Delete(r.Context(), id); err != nil {
		slog.Warn(err.Error(), "handler", "DeleteGame")
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete game"})
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Game deleted successfully"})
}

func (h *Handler) GetGameById(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	repo := repository.NewGameRepository(h.DB)
	game, err := repo.GetGameDetails(r.Context(), id) // короткий ответ, если нужен см. GetById
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetGameById")
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch game"})
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, game.ToResponseGameDetails())
}

func (h *Handler) ListGames(w http.ResponseWriter, r *http.Request) {
	repo := repository.NewGameRepository(h.DB)
	games, err := repo.ListGamesDetails(r.Context())
	if err != nil {
		slog.Warn(err.Error(), "handler", "ListGames")
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Failed to fetch games"})
		return
	}
	gameResponses := make([]*server.GameResponse, 0, len(games))
	for _, game := range games {
		gameResponses = append(gameResponses, game.ToResponseGameDetails())
	}

	helper.RespondWithJSON(w, http.StatusOK, gameResponses)
}

func (h *Handler) UpdateGame(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// fixme update не проверяет есть ли запись в бд
	var game server.GameRequest
	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateGame")
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	repo := repository.NewGameRepository(h.DB)
	updateGame := &model.Game{
		StartTime:   game.StartTime,
		EndTime:     game.EndTime,
		Description: *game.Description,
	}
	updateGame.Id = id
	err := repo.Update(r.Context(), updateGame)
	if err != nil {
		slog.Warn(err.Error(), "handler", "UpdateGame")
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Invalid request payload"})
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Game updated successfully"})
}
