package api

import (
	"ctf01d/internal/app/models"
	"ctf01d/internal/app/repository"
	"ctf01d/internal/app/view"
	"database/sql"
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
		respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	respondWithJSON(w, http.StatusOK, view.NewUniversitiesFromModels(universities))

}
