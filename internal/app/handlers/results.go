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
		GameId: gameId,
		TeamId: result.TeamId,
		Score:  result.Score,
	}
	if err = repo.Create(r.Context(), newResult); err != nil {
		slog.Warn(err.Error(), "handler", "CreateResultHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create result"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewResultFromModel(newResult, 0))
}

func (h *Handlers) GetResult(w http.ResponseWriter, r *http.Request, gameId openapi_types.UUID, resultId openapi_types.UUID) {
	repo := repository.NewResultRepository(h.DB)
	result, err := repo.GetById(r.Context(), gameId)
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetResult")
		// todo - empty result ?
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Failed to fetch result"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewResultFromModel(result, 0))
}

func (h *Handlers) UpdateResult(w http.ResponseWriter, r *http.Request, gameId openapi_types.UUID, resultId openapi_types.UUID) {
	var resultRequest server.ResultRequest
	if err := json.NewDecoder(r.Body).Decode(&resultRequest); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateResult")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}

	repo := repository.NewResultRepository(h.DB)
	result := &dbmodel.Result{
		Id:     resultId,
		GameId: gameId,
		TeamId: resultRequest.TeamId,
		Score:  resultRequest.Score,
	}

	if err := repo.Update(r.Context(), result); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateResult")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to update result"})
		return
	}

	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewResultFromModel(result, 0))
}

// GetScoreboard retrieves the scoreboard for a given game ID
func (h *Handlers) GetScoreboard(w http.ResponseWriter, r *http.Request, gameId openapi_types.UUID) {
	repo := repository.NewResultRepository(h.DB)
	results, err := repo.List(r.Context(), gameId)
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetScoreboard")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch scoreboard"})
		return
	}

	scoreboard := view.NewScoreboardFromResults(results)
	api_helpers.RespondWithJSON(w, http.StatusOK, scoreboard)
}
