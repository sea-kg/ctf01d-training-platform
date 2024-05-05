package api

import (
	"ctf01d/internal/app/models"
	"ctf01d/internal/app/repository"
	api_helpers "ctf01d/internal/app/utils"
	"ctf01d/internal/app/view"
	"database/sql"
	"log/slog"
	"net/http"
)

func ListUniversitiesHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query().Get("term")

	universityRepo := repository.NewUniversityRepository(db)
	var universities []*models.University
	var err error

	if queryParam != "" {
		universities, err = universityRepo.Search(r.Context(), queryParam)
	} else {
		universities, err = universityRepo.List(r.Context())
	}

	if err != nil {
		slog.Warn(err.Error(), "handler", "ListUniversitiesHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Failed to fetch universities"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewUniversitiesFromModels(universities))
}
