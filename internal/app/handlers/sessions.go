package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"sync"

	"ctf01d/internal/app/repository"
	"ctf01d/internal/app/server"
	api_helpers "ctf01d/internal/app/utils"

	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type SessionCache struct {
	cache sync.Map
}

func NewSessionCache() *SessionCache {
	return &SessionCache{}
}

func (sc *SessionCache) GetSession(sessionID string) (openapi_types.UUID, bool) {
	val, ok := sc.cache.Load(sessionID)
	if !ok {
		return uuid.Nil, false
	}
	return val.(openapi_types.UUID), true
}

func (sc *SessionCache) SetSession(sessionID string, userID uuid.UUID) {
	sc.cache.Store(sessionID, userID)
}

func (sc *SessionCache) DeleteSession(sessionID string) {
	sc.cache.Delete(sessionID)
}

type SessionHandler struct {
	*Handlers
	SessionCache *SessionCache
}

func NewSessionHandler(handlers *Handlers) *SessionHandler {
	return &SessionHandler{
		Handlers:     handlers,
		SessionCache: NewSessionCache(),
	}
}

func (h *SessionHandler) PostApiV1AuthSignIn(w http.ResponseWriter, r *http.Request) {
	var req server.PostApiV1AuthSignInJSONBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Warn(err.Error(), "handler", "PostApiV1AuthSignIn")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	userRepo := repository.NewUserRepository(h.DB)
	user, err := userRepo.GetByUserName(r.Context(), *req.UserName)
	if err != nil || !api_helpers.CheckPasswordHash(*req.Password, user.PasswordHash) {
		slog.Warn(err.Error(), "handler", "PostApiV1AuthSignIn")
		api_helpers.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "Invalid password or user"})
		return
	}

	repo := repository.NewSessionRepository(h.DB)
	slog.Debug("user.Id " + openapi_types.UUID(user.Id).String())

	sessionId, err := repo.StoreSessionInDB(r.Context(), user.Id)
	if err != nil {
		slog.Warn(err.Error(), "handler", "PostApiV1AuthSignIn")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to store session"})
		return
	}

	// Добавляем сессию в кэш
	h.SessionCache.SetSession(sessionId, user.Id)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		HttpOnly: true,
		Value:    sessionId,
		Path:     "/",
		MaxAge:   96 * 3600, // fixme, брать из db
	})

	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "User logged in"})
}

func (h *Handler) PostApiV1AuthSignOut(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		slog.Warn(err.Error(), "handler", "PostApiV1AuthSignOut")
		api_helpers.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "No session found"})
		return
	}
	repo := repository.NewSessionRepository(h.DB)
	err = repo.DeleteSessionInDB(r.Context(), cookie.Value)
	if err != nil {
		slog.Warn(err.Error(), "handler", "PostApiV1AuthSignOut")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete session"})
		return
	}

	// Удаляем сессию из кэша
	h.SessionCache.DeleteSession(cookie.Value)

	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1, // Удаление куки
	})
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "User logout successful"})
}

func (h *SessionHandler) ValidateSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		slog.Warn(err.Error(), "handler", "ValidateSession")
		api_helpers.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "No session found"})
		return
	}

	if userId, ok := h.SessionCache.GetSession(cookie.Value); ok {
		slog.Debug("ValidateSession user.Id " + openapi_types.UUID(userId).String())
		h.respondWithUserDetails(w, r, userId)
		return
	}

	slog.Debug("cookie.Value, " + cookie.Value)
	repo := repository.NewSessionRepository(h.DB)
	var userId openapi_types.UUID
	userId, err = repo.GetSessionFromDB(r.Context(), cookie.Value)
	if err != nil {
		slog.Warn(err.Error(), "handler", "ValidateSession")
		api_helpers.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "No user or session found"})
		return
	}

	h.SessionCache.SetSession(cookie.Value, userId)
	slog.Debug("ValidateSession user.Id " + openapi_types.UUID(userId).String())
	h.respondWithUserDetails(w, r, userId)
}

func (h *SessionHandler) respondWithUserDetails(w http.ResponseWriter, r *http.Request, userId openapi_types.UUID) {
	userRepo := repository.NewUserRepository(h.DB)
	user, err := userRepo.GetById(r.Context(), userId)
	if err != nil {
		slog.Warn(err.Error(), "handler", "respondWithUserDetails")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Could not find user by user id"})
		return
	}
	res := make(map[string]string)
	res["name"] = user.DisplayName
	res["role"] = api_helpers.ConvertUserRequestRoleToString(user.Role)

	api_helpers.RespondWithJSON(w, http.StatusOK, res)
}
