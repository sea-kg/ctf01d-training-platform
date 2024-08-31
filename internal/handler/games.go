package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"ctf01d/internal/model"
	"ctf01d/internal/repository"
	"ctf01d/internal/server"
	api_helpers "ctf01d/internal/utils"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Handler) CreateGame(w http.ResponseWriter, r *http.Request) {
	var game server.GameRequest
	var err error
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
	newGame := &model.Game{
		StartTime:   game.StartTime,
		EndTime:     game.EndTime,
		Description: *game.Description,
	}

	err = repo.Create(r.Context(), newGame)
	if err != nil {
		slog.Warn(err.Error(), "handler", "CreateGame")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create game"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, newGame.ToResponse())
}

func (h *Handler) DeleteGame(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	repo := repository.NewGameRepository(h.DB)
	if err := repo.Delete(r.Context(), id); err != nil {
		slog.Warn(err.Error(), "handler", "DeleteGame")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete game"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Game deleted successfully"})
}

func (h *Handler) GetGameById(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	repo := repository.NewGameRepository(h.DB)
	game, err := repo.GetGameDetails(r.Context(), id) // короткий ответ, если нужен см. GetById
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetGameById")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch game"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, game.ToResponseGameDetails())
}

func (h *Handler) ListGames(w http.ResponseWriter, r *http.Request) {
	repo := repository.NewGameRepository(h.DB)
	games, err := repo.ListGamesDetails(r.Context())
	if err != nil {
		slog.Warn(err.Error(), "handler", "ListGames")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Failed to fetch games"})
		return
	}
	gameResponses := make([]*server.GameResponse, 0, len(games))
	for _, game := range games {
		gameResponses = append(gameResponses, game.ToResponseGameDetails())
	}

	api_helpers.RespondWithJSON(w, http.StatusOK, gameResponses)
}

func (h *Handler) UpdateGame(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// fixme update не проверяет есть ли запись в бд
	var game server.GameRequest
	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateGame")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
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
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Invalid request payload"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Game updated successfully"})
}
