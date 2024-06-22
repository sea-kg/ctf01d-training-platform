package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	dbmodel "ctf01d/internal/app/db"
	"ctf01d/internal/app/repository"
	"ctf01d/internal/app/server"
	api_helpers "ctf01d/internal/app/utils"
	"ctf01d/internal/app/view"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Handlers) CreateResult(w http.ResponseWriter, r *http.Request) {
	var result server.ResultRequest
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		slog.Warn(err.Error(), "handler", "CreateResultHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	repo := repository.NewResultRepository(h.DB)
	newResult := &dbmodel.Result{
		TeamId: result.TeamId,
		GameId: result.GameId,
		Score:  result.Score,
		Rank:   result.Rank,
	}
	if err := repo.Create(r.Context(), newResult); err != nil {
		slog.Warn(err.Error(), "handler", "CreateGameHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create result"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Game created successfully"})
}

func (h *Handlers) GetResultById(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	repo := repository.NewResultRepository(h.DB)
	result, err := repo.GetById(r.Context(), id)
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetGameByIdHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch result"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewResultFromModel(result))
}

func (h *Handlers) ListResults(w http.ResponseWriter, r *http.Request) {
	repo := repository.NewResultRepository(h.DB)
	results, err := repo.List(r.Context())
	if err != nil {
		slog.Warn(err.Error(), "handler", "ListGamesHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Failed to fetch results"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewResultFromModels(results))
}
