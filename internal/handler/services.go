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

func (h *Handler) CreateService(w http.ResponseWriter, r *http.Request) {
	var service server.ServiceRequest
	var err error
	if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
		slog.Warn(err.Error(), "handler", "CreateServiceHandler")
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	repo := repository.NewServiceRepository(h.DB)
	newService := &model.Service{
		Name:        service.Name,
		Author:      service.Author,
		LogoUrl:     helper.ToNullString(service.LogoUrl),
		Description: *service.Description,
		IsPublic:    service.IsPublic,
	}
	if err = repo.Create(r.Context(), newService); err != nil {
		slog.Warn(err.Error(), "handler", "CreateServiceHandler")
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create service"})
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, newService.ToResponse())
}

func (h *Handler) DeleteService(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	repo := repository.NewServiceRepository(h.DB)
	if err := repo.Delete(r.Context(), id); err != nil {
		slog.Warn(err.Error(), "handler", "DeleteServiceHandler")
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete service"})
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Service deleted successfully"})
}

func (h *Handler) GetServiceById(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	repo := repository.NewServiceRepository(h.DB)
	service, err := repo.GetById(r.Context(), id)
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetServiceByIdHandler")
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch service"})
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, service.ToResponse())
}

func (h *Handler) ListServices(w http.ResponseWriter, r *http.Request) {
	repo := repository.NewServiceRepository(h.DB)
	services, err := repo.List(r.Context())
	if err != nil {
		slog.Warn(err.Error(), "handler", "ListServicesHandler")
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch services"})
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, model.NewServiceFromModels(services))
}

func (h *Handler) UpdateService(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	var sr server.ServiceRequest
	if err := json.NewDecoder(r.Body).Decode(&sr); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateService")
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	repo := repository.NewServiceRepository(h.DB)
	service := &model.Service{
		Id:          id,
		Name:        sr.Name,
		Author:      sr.Author,
		LogoUrl:     helper.ToNullString(sr.LogoUrl),
		Description: *sr.Description,
		IsPublic:    sr.IsPublic,
	}
	err := repo.Update(r.Context(), service)
	if err != nil {
		slog.Warn(err.Error(), "handler", "UpdateService")
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Invalid request payload"})
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Service updated successfully"})
}

// fixme implement
func (h *Handler) UploadChecker(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}

// fixme implement
func (h *Handler) UploadService(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}
