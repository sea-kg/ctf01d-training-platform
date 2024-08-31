package handler

import (
	"log/slog"
	"net/http"

	"ctf01d/internal/model"
	"ctf01d/internal/repository"
	"ctf01d/internal/server"
	api_helpers "ctf01d/internal/utils"
)

func (h *Handler) ListUniversities(w http.ResponseWriter, r *http.Request, params server.ListUniversitiesParams) {
	queryParam := r.URL.Query().Get("term")

	repo := repository.NewUniversityRepository(h.DB)
	var universities []*model.University
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
	api_helpers.RespondWithJSON(w, http.StatusOK, model.NewUniversitiesFromModels(universities))
}
