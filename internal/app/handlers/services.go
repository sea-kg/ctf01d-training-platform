package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	dbmodel "ctf01d/internal/app/db"
	dbmodels "ctf01d/internal/app/db"
	"ctf01d/internal/app/repository"
	"ctf01d/internal/app/server"
	api_helpers "ctf01d/internal/app/utils"
	"ctf01d/internal/app/view"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Handlers) CreateService(w http.ResponseWriter, r *http.Request) {
	var service server.ServiceRequest
	var err error
	if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
		slog.Warn(err.Error(), "handler", "CreateServiceHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	repo := repository.NewServiceRepository(h.DB)
	newService := &dbmodels.Service{
		Name:        service.Name,
		Author:      service.Author,
		LogoUrl:     api_helpers.ToNullString(service.LogoUrl),
		Description: *service.Description,
		IsPublic:    service.IsPublic,
	}
	if err = repo.Create(r.Context(), newService); err != nil {
		slog.Warn(err.Error(), "handler", "CreateServiceHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create service"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewServiceFromModel(newService))
}

func (h *Handlers) DeleteService(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	repo := repository.NewServiceRepository(h.DB)
	if err := repo.Delete(r.Context(), id); err != nil {
		slog.Warn(err.Error(), "handler", "DeleteServiceHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete service"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Service deleted successfully"})
}

func (h *Handlers) GetServiceById(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	repo := repository.NewServiceRepository(h.DB)
	service, err := repo.GetById(r.Context(), id)
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetServiceByIdHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch service"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewServiceFromModel(service))
}

func (h *Handlers) ListServices(w http.ResponseWriter, r *http.Request) {
	repo := repository.NewServiceRepository(h.DB)
	services, err := repo.List(r.Context())
	if err != nil {
		slog.Warn(err.Error(), "handler", "ListServicesHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch services"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewServiceFromModels(services))
}

func (h *Handlers) UpdateService(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	var sr server.ServiceRequest
	if err := json.NewDecoder(r.Body).Decode(&sr); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateService")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	repo := repository.NewServiceRepository(h.DB)
	service := &dbmodel.Service{
		Id:          id,
		Name:        sr.Name,
		Author:      sr.Author,
		LogoUrl:     api_helpers.ToNullString(sr.LogoUrl),
		Description: *sr.Description,
		IsPublic:    sr.IsPublic,
	}
	err := repo.Update(r.Context(), service)
	if err != nil {
		slog.Warn(err.Error(), "handler", "UpdateService")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Invalid request payload"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Service updated successfully"})
}

// fixme implement
func (h *Handlers) PostApiV1ServicesUuidChecker(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}

// fixme implement
func (h *Handlers) PostApiV1ServicesUuidService(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}
