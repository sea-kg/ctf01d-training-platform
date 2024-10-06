package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"ctf01d/internal/helper"
	"ctf01d/internal/httpserver"
	"ctf01d/internal/model"
	"ctf01d/internal/repository"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Handler) CreateService(w http.ResponseWriter, r *http.Request) {
	var service httpserver.ServiceRequest
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
	var sr httpserver.ServiceRequest
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

func (h *Handler) UploadService(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	repo := repository.NewServiceRepository(h.DB)
	_, err := repo.GetById(r.Context(), id)
	if err != nil {
		slog.Warn(err.Error(), "handler", "UploadService")
		helper.RespondWithJSON(w, http.StatusNotFound, map[string]string{"error": "Unable to fetch service"})
		return
	}

	var req httpserver.UploadServiceMultipartRequestBody
	boundedReader := http.MaxBytesReader(w, r.Body, 100<<20) // 100Mb todo externalize to props
	if err := json.NewDecoder(boundedReader).Decode(&req); err != nil {
		slog.Warn(err.Error(), "handler", "UploadService")
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}

	f := req.File
	reader, err := f.Reader()
	if err != nil {
		slog.Warn(err.Error(), "handler", "UploadService")
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Unable to read bytes"})
		return
	}

	reader, err = validateUploadService(reader)
	if err != nil {
		slog.Warn(err.Error(), "handler", "UploadService")
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}
}

func validateUploadService(reader io.Reader) (io.ReadCloser, error) {
	footprint := make([]byte, 4)
	if _, err := io.ReadFull(reader, footprint); err != nil {
		slog.Warn("Unable to read file", "handler", "UploadService")
		return nil, errors.New(fmt.Sprintf("Unable to read file: %s", err.Error()))
	}

	if !helper.IsZip(footprint) {
		slog.Warn("Uploaded file is not a zip", "handler", "UploadService")
		return nil, errors.New("Uploaded file is not a zip")
	}

	restored := io.MultiReader(bytes.NewReader(footprint), reader)

	return io.NopCloser(restored), nil
}
