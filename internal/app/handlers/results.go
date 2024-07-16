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

func (h *Handlers) CreateResult(w http.ResponseWriter, r *http.Request, gameId openapi_types.UUID) {
	var result server.ResultRequest
	var err error
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
	if err = repo.Create(r.Context(), newResult); err != nil {
		slog.Warn(err.Error(), "handler", "CreateResultHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create result"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewResultFromModel(newResult))
}

func (h *Handlers) GetResult(w http.ResponseWriter, r *http.Request, gameId openapi_types.UUID) {
	repo := repository.NewResultRepository(h.DB)
	result, err := repo.GetById(r.Context(), gameId)
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetResult")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Failed to fetch result"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewResultFromModel(result))
}

func (h *Handlers) UpdateResult(w http.ResponseWriter, r *http.Request, teamId openapi_types.UUID, userId openapi_types.UUID) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}
