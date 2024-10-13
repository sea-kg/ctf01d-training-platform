package handler

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"time"

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
	e, err := repo.GetById(r.Context(), id)
	if err != nil {
		slog.Warn(err.Error(), "handler", "UploadService")
		helper.RespondWithJSON(w, http.StatusNotFound, map[string]string{"error": "Unable to fetch service"})
		return
	}

	if e.IsServiceValid {
		slog.Warn("Service is valid already. Skip uploading.", "handler", "UploadService")
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Already available"})
		return
	}

	var req httpserver.UploadServiceMultipartRequestBody
	boundedReader := http.MaxBytesReader(w, r.Body, 100<<20) // 100Mb todo externalize to props
	if err := json.NewDecoder(boundedReader).Decode(&req); err != nil {
		slog.Warn(err.Error(), "handler", "UploadService")
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Bad request payload"})
		return
	}

	f := req.File
	reader, err := f.Reader()
	if err != nil {
		slog.Warn(err.Error(), "handler", "UploadService")
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Unable to read bytes"})
		return
	}

	reader, err = validateSignature(reader)
	if err != nil {
		slog.Warn("Unsupported archive type: "+err.Error(), "handler", "UploadService")
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Bad request"})
		return
	}

	tempFile, err := storeTemp(f.Filename(), id, reader)
	if err != nil {
		slog.Warn(err.Error(), "handler", "UploadService")
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal error"})
		return
	}

	// isolation write commited
	tx, err := h.DB.BeginTx(r.Context(), &sql.TxOptions{Isolation: sql.IsolationLevel(3), ReadOnly: false})
	if err != nil {
		slog.Warn("Unable to begin transaction: "+err.Error(), "handler", "UploadService")
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal error"})
		return
	}

	if err = handleUploadService(tx, e, tempFile, f.Filename()); err != nil {
		tx.Rollback()
		os.Remove(tempFile.Name())
		slog.Warn(err.Error(), "handler", "UploadService")
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Cannot store"})
		return
	}

	if err = tx.Commit(); err != nil {
		slog.Error(fmt.Sprintf("Unable to commit: %s", err.Error()), "handler", "UploadService")
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal error"})
		return
	}

	slog.Info(fmt.Sprintf("File uploaded successfully: %s", f.Filename()), "handler", "UploadService")
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "Service uploaded successfully"})
}

func storeTemp(fileName string, id openapi_types.UUID, reader io.Reader) (*os.File, error) {
	storagePath := filepath.Join("srv", "ctf-platform", "temp")
	if err := os.MkdirAll(storagePath, os.ModePerm); err != nil {
		msg := fmt.Sprintf("Unable to create temp storage path: %s: %s", storagePath, err.Error())
		return nil, errors.New(msg)
	}

	compositeName := fmt.Sprintf("%s-%s", fileName, id)
	fileAbs := filepath.Join(storagePath, compositeName)
	file, err := os.Create(fileAbs)
	if err != nil {
		msg := fmt.Sprintf("Unable to create temp file: %s: %s", fileAbs, err.Error())
		return nil, errors.New(msg)
	}

	bytesWritten, err := io.Copy(file, reader)
	if err != nil {
		msg := fmt.Sprintf("Unable to write bytes to temp file: %s: %s", fileAbs, err.Error())
		return nil, errors.New(msg)
	} else {
		msg := fmt.Sprintf("%s written with size: %d under temp storage", fileAbs, bytesWritten)
		slog.Info(msg)
	}

	return file, nil
}

func handleUploadService(tx *sql.Tx, e *model.Service, content *os.File, fileName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := tx.ExecContext(ctx, "UPDATE services SET is_service_valid = $1 WHERE id = $2", true, e.Id)
	if err != nil {
		return err
	}

	storeDir := filepath.Join("srv", "ctf-platform", "files", helper.IdToPath(e.Id))
	if err := os.MkdirAll(storeDir, os.ModePerm); err != nil {
		return errors.New(fmt.Sprintf("Unable to create storage path: %s: ", storeDir) + err.Error())
	}

	newAbsPath := filepath.Join(storeDir, fileName)
	if oldAbsPath, err := filepath.Abs(content.Name()); err == nil {
		if err = os.Rename(oldAbsPath, newAbsPath); err != nil {
			return err
		}
	} else {
		return err
	}

	return nil
}

func validateSignature(reader io.Reader) (io.ReadCloser, error) {
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
