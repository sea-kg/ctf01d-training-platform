package api

import (
	"ctf01d/internal/app/models"
	"ctf01d/internal/app/repository"
	api_helpers "ctf01d/internal/app/utils"
	"ctf01d/internal/app/view"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type RequestService struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	LogoUrl     string `json:"logo_url"`
	Description string `json:"description"`
	IsPublic    bool   `json:"is_public"`
}

func CreateServiceHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var service RequestService
	if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload: " + err.Error()})
		return
	}

	serviceRepo := repository.NewServiceRepository(db)
	// fixme request to model надо вынести и переиспользовать
	newService := &models.Service{
		Name:        service.Name,
		Author:      service.Author,
		LogoUrl:     api_helpers.PrepareImage(service.LogoUrl),
		Description: service.Description,
		IsPublic:    service.IsPublic,
	}
	if err := serviceRepo.Create(r.Context(), newService); err != nil {
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create service: " + err.Error()})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Service created successfully"})
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
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Service deleted successfully"})
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
