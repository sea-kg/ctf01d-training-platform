package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"ctf01d/internal/helper"
	"ctf01d/internal/model"
	"ctf01d/internal/repository"
	"ctf01d/internal/server"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Handler) CreateResult(w http.ResponseWriter, r *http.Request, gameId openapi_types.UUID) {
	var result server.ResultRequest
	var err error
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		slog.Warn(err.Error(), "handler", "CreateResultHandler")
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	repo := repository.NewResultRepository(h.DB)
	newResult := &model.Result{
		GameId: gameId,
		TeamId: result.TeamId,
		Score:  result.Score,
	}
	if err = repo.Create(r.Context(), newResult); err != nil {
		slog.Warn(err.Error(), "handler", "CreateResultHandler")
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create result"})
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, newResult.ToResponse(0))
}

func (h *Handler) GetResult(w http.ResponseWriter, r *http.Request, gameId openapi_types.UUID, resultId openapi_types.UUID) {
	repo := repository.NewResultRepository(h.DB)
	result, err := repo.GetById(r.Context(), gameId)
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetResult")
		// todo - empty result ?
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Failed to fetch result"})
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, result.ToResponse(0))
}

func (h *Handler) UpdateResult(w http.ResponseWriter, r *http.Request, gameId openapi_types.UUID, resultId openapi_types.UUID) {
	var resultRequest server.ResultRequest
	if err := json.NewDecoder(r.Body).Decode(&resultRequest); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateResult")
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}

	repo := repository.NewResultRepository(h.DB)
	result := &model.Result{
		Id:     resultId,
		GameId: gameId,
		TeamId: resultRequest.TeamId,
		Score:  resultRequest.Score,
	}

	if err := repo.Update(r.Context(), result); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateResult")
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to update result"})
		return
	}

	helper.RespondWithJSON(w, http.StatusOK, result.ToResponse(0))
}

// GetScoreboard retrieves the scoreboard for a given game ID
func (h *Handler) GetScoreboard(w http.ResponseWriter, r *http.Request, gameId openapi_types.UUID) {
	repo := repository.NewResultRepository(h.DB)
	results, err := repo.List(r.Context(), gameId)
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetScoreboard")
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch scoreboard"})
		return
	}

	scoreboard := model.NewScoreboardFromResults(results)
	helper.RespondWithJSON(w, http.StatusOK, scoreboard)
}
