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

func CreateResultHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var result apimodels.ResultRequest
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		slog.Warn(err.Error(), "handler", "CreateResultHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	repo := repository.NewResultRepository(db)
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

func GetResultByIdHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetGameByIdHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Bad request"})
		return
	}
	repo := repository.NewResultRepository(db)
	result, err := repo.GetById(r.Context(), id)
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetGameByIdHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch result"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewResultFromModel(result))
}

func ListResultsHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	repo := repository.NewResultRepository(db)
	results, err := repo.List(r.Context())
	if err != nil {
		slog.Warn(err.Error(), "handler", "ListGamesHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Failed to fetch results"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewResultFromModels(results))
}
