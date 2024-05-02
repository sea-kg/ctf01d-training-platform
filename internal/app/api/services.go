package api

import (
	"ctf01d/internal/app/repository"
	api_helpers "ctf01d/internal/app/utils"
	"ctf01d/internal/app/view"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// fixme implement
func CreateServiceHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func DeleteServiceHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Bad request"})
		return
	}
	serviceRepo := repository.NewServiceRepository(db)
	if err := serviceRepo.Delete(r.Context(), id); err != nil {
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete service"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Team deleted successfully"})
}

func GetServiceByIdHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Bad request"})
		return
	}
	serviceRepo := repository.NewServiceRepository(db)
	service, err := serviceRepo.GetById(r.Context(), id)
	if err != nil {
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch service"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewServiceFromModel(service))
}

func ListServicesHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	serviceRepo := repository.NewServiceRepository(db)
	services, err := serviceRepo.List(r.Context())
	if err != nil {
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewServiceFromModels(services))
}

// fixme implement
func UpdateServiceHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
