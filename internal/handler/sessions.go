package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"ctf01d/internal/helper"
	"ctf01d/internal/repository"
	"ctf01d/internal/server"
	"ctf01d/internal/view"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Handler) SignInUser(w http.ResponseWriter, r *http.Request) {
	var req server.SignInUserJSONBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Warn(err.Error(), "handler", "SignInUser")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	userRepo := repository.NewUserRepository(h.DB)
	user, err := userRepo.GetByUserName(r.Context(), *req.UserName)
	if err != nil || !helper.CheckPasswordHash(*req.Password, user.PasswordHash) {
		slog.Warn(err.Error(), "handler", "SignInUser")
		helper.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "Invalid password or user"})
		return
	}

	repo := repository.NewSessionRepository(h.DB)
	slog.Debug("user.Id " + openapi_types.UUID(user.Id).String())

	sessionId, err := repo.StoreSessionInDB(r.Context(), user.Id)
	if err != nil {
		slog.Warn(err.Error(), "handler", "SignInUser")
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to store session"})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		HttpOnly: true,
		Value:    sessionId,
		Path:     "/",
		MaxAge:   96 * 3600, // fixme, брать из db
	})

	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "User logged in"})
}

func (h *Handler) SignOutUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		slog.Warn(err.Error(), "handler", "SignOutUser")
		helper.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "No session found"})
		return
	}
	repo := repository.NewSessionRepository(h.DB)
	err = repo.DeleteSessionInDB(r.Context(), cookie.Value)
	if err != nil {
		slog.Warn(err.Error(), "handler", "SignOutUser")
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete session"})
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1, // Удаление куки
	})
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "User logout successful"})
}

func (h *Handler) ValidateSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		slog.Warn(err.Error(), "handler", "ValidateSession")
		helper.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "No session found"})
		return
	}
	slog.Debug("cookie.Value, " + cookie.Value)
	repo := repository.NewSessionRepository(h.DB)
	var userId openapi_types.UUID
	userId, err = repo.GetSessionFromDB(r.Context(), cookie.Value)
	if err != nil {
		slog.Warn(err.Error(), "handler", "ValidateSession")
		helper.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "No user or session found"})
		return
	}
	slog.Debug("ValidateSession user.Id " + openapi_types.UUID(userId).String())

	userRepo := repository.NewUserRepository(h.DB)
	user, err := userRepo.GetById(r.Context(), userId)
	if err != nil {
		slog.Warn(err.Error(), "handler", "ValidateSession")
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Could not find user by user id"})
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, view.NewSessionFromModel(user))
}
