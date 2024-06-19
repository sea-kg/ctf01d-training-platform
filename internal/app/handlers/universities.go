package handlers

import (
	"log/slog"
	"net/http"

	dbmodels "ctf01d/internal/app/db"
	"ctf01d/internal/app/repository"
	"ctf01d/internal/app/server"
	api_helpers "ctf01d/internal/app/utils"
	"ctf01d/internal/app/view"
)

func (h *Handlers) GetApiV1Universities(w http.ResponseWriter, r *http.Request, params server.GetApiV1UniversitiesParams) {
	queryParam := r.URL.Query().Get("term")

	repo := repository.NewUniversityRepository(h.DB)
	var universities []*dbmodels.University
	var err error

	if queryParam != "" {
		universities, err = repo.Search(r.Context(), queryParam)
	} else {
		universities, err = repo.List(r.Context())
	}

	if err != nil {
		slog.Warn(err.Error(), "handler", "ListUniversitiesHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Failed to fetch universities"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewUniversitiesFromModels(universities))
}
