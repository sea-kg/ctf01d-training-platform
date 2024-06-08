package api

import (
	apimodels "ctf01d/internal/app/apimodels"
	dbmodels "ctf01d/internal/app/db"
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

func CreateServiceHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var service apimodels.ServiceRequest
	if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
		slog.Warn(err.Error(), "handler", "CreateServiceHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	repo := repository.NewServiceRepository(db)
	newService := &dbmodels.Service{
		Name:        service.Name,
		Author:      service.Author,
		LogoUrl:     api_helpers.PrepareImage(*service.LogoUrl),
		Description: *service.Description,
		IsPublic:    service.IsPublic,
	}
	if err := repo.Create(r.Context(), newService); err != nil {
		slog.Warn(err.Error(), "handler", "CreateServiceHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create service"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Service created successfully"})
}

func DeleteServiceHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		slog.Warn(err.Error(), "handler", "DeleteServiceHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Bad request"})
		return
	}
	repo := repository.NewServiceRepository(db)
	if err := repo.Delete(r.Context(), id); err != nil {
		slog.Warn(err.Error(), "handler", "DeleteServiceHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete service"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Service deleted successfully"})
}

func GetServiceByIdHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetServiceByIdHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Bad request"})
		return
	}
	repo := repository.NewServiceRepository(db)
	service, err := repo.GetById(r.Context(), id)
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetServiceByIdHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch service"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewServiceFromModel(service))
}

func ListServicesHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	repo := repository.NewServiceRepository(db)
	services, err := repo.List(r.Context())
	if err != nil {
		slog.Warn(err.Error(), "handler", "ListServicesHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch services"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewServiceFromModels(services))
}

// fixme implement
func UpdateServiceHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
