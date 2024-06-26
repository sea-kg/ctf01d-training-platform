package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"ctf01d/internal/app/repository"
	"ctf01d/internal/app/server"
	"ctf01d/internal/app/utils"
	api_helpers "ctf01d/internal/app/utils"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Handlers) PostApiV1AuthSignin(w http.ResponseWriter, r *http.Request) {
	var req server.PostApiV1AuthSigninJSONBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Warn(err.Error(), "handler", "PostApiV1AuthSignin")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	userRepo := repository.NewUserRepository(h.DB)
	user, err := userRepo.GetByUserName(r.Context(), *req.UserName)
	if err != nil || !api_helpers.CheckPasswordHash(*req.Password, user.PasswordHash) {
		slog.Warn(err.Error(), "handler", "PostApiV1AuthSignin")
		api_helpers.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "Invalid password or user"})
		return
	}

	repo := repository.NewSessionRepository(h.DB)
	// slog.Info("user.Id " + openapi_types.UUID(user.Id).String())

	sessionId, err := repo.StoreSessionInDB(r.Context(), user.Id)
	if err != nil {
		slog.Warn(err.Error(), "handler", "PostApiV1AuthSignin")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to store session"})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		HttpOnly: true,
		Value:    sessionId,
		Path:     "/",
		MaxAge:   96 * 3600, // fixme, брать из db
	})

	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "User logged in"})
}

func (h *Handlers) PostApiV1AuthSignout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		slog.Warn(err.Error(), "handler", "PostApiV1AuthSignout")
		api_helpers.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "No session found"})
		return
	}
	repo := repository.NewSessionRepository(h.DB)
	err = repo.DeleteSessionInDB(r.Context(), cookie.Value)
	if err != nil {
		slog.Warn(err.Error(), "handler", "PostApiV1AuthSignout")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete session"})
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1, // Удаление куки
	})
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "User logout successful"})
}

func (h *Handlers) ValidateSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		slog.Warn(err.Error(), "handler", "PostApiV1AuthSession")
		api_helpers.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "No session found"})
		return
	}
	// slog.Info("cookie.Value, " + cookie.Value)
	repo := repository.NewSessionRepository(h.DB)
	var userId openapi_types.UUID
	userId, err = repo.GetSessionFromDB(r.Context(), cookie.Value)
	if err != nil {
		slog.Warn(err.Error(), "handler", "PostApiV1AuthSession")
		api_helpers.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "No user or session found"})
		return
	}
	// slog.Info("ValidateSession user.Id " + openapi_types.UUID(userId).String())

	userRepo := repository.NewUserRepository(h.DB)
	user, err := userRepo.GetById(r.Context(), userId)
	if err != nil {
		slog.Warn(err.Error(), "handler", "PostApiV1AuthSession")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Could not find user by user id"})
		return
	}
	var res = make(map[string]string)
	res["name"] = user.DisplayName
	res["role"] = helpers.ConvertUserRequestRoleToString(user.Role)

	api_helpers.RespondWithJSON(w, http.StatusOK, res)
}
