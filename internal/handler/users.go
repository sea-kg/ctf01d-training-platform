package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"ctf01d/internal/model"
	"ctf01d/internal/repository"
	"ctf01d/internal/server"
	api_helpers "ctf01d/internal/utils"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// fixme обернуть в транзакцию, т.к. две вставки подряд
	var user server.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		slog.Warn(err.Error(), "handler", "CreateUserHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	repo := repository.NewUserRepository(h.DB)
	passwordHash, err := api_helpers.HashPassword(user.Password)
	if err != nil {
		slog.Warn(err.Error(), "handler", "CreateUserHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	newUser := &model.User{
		Username:     user.UserName,
		DisplayName:  api_helpers.ToNullString(user.DisplayName),
		Role:         user.Role,
		Status:       user.Status,
		PasswordHash: passwordHash,
		AvatarUrl:    api_helpers.ToNullString(user.AvatarUrl),
	}
	if err = repo.Create(r.Context(), newUser); err != nil {
		slog.Warn(err.Error(), "handler", "CreateUserHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
		return
	}
	if user.TeamIds != nil && len(*user.TeamIds) > 0 {
		if err := repo.AddUserToTeams(r.Context(), newUser.Id, user.TeamIds); err != nil {
			slog.Warn(err.Error(), "handler", "CreateUserHandler")
			api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to add user to teams"})
			return
		}
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, model.NewUserFromModel(newUser))
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	repo := repository.NewUserRepository(h.DB)
	if err := repo.Delete(r.Context(), id); err != nil {
		slog.Warn(err.Error(), "handler", "DeleteUserHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "User deleted successfully"})
}

func (h *Handler) GetUserById(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	repo := repository.NewUserRepository(h.DB)
	user, err := repo.GetById(r.Context(), id)
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetUserByIdHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch user"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, model.NewUserFromModel(user))
}

func (h *Handler) GetProfileById(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	repo := repository.NewUserRepository(h.DB)
	userProfile, err := repo.GetProfileWithHistory(r.Context(), id)
	if err != nil {
		slog.Info(err.Error(), "handler", "GetProfileByIdHandler")
		api_helpers.RespondWithJSON(w, http.StatusNotFound, map[string]string{"data": "User have not profile"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, model.NewProfileFromModel(userProfile))
}

func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	repo := repository.NewUserRepository(h.DB)
	users, err := repo.List(r.Context())
	if err != nil {
		slog.Warn(err.Error(), "handler", "ListUsersHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, model.NewUsersFromModels(users))
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// fixme update не проверяет есть ли запись в бд
	var user server.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateUserHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	passwordHash, err := api_helpers.HashPassword(user.Password)
	if err != nil {
		slog.Warn(err.Error(), "handler", "UpdateUserHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	repo := repository.NewUserRepository(h.DB)
	updateUser := &model.User{
		Username:     user.UserName,
		DisplayName:  api_helpers.ToNullString(user.DisplayName),
		Role:         user.Role,
		Status:       user.Status,
		PasswordHash: passwordHash,
		AvatarUrl:    api_helpers.ToNullString(user.AvatarUrl),
	}
	updateUser.Id = id
	err = repo.Update(r.Context(), updateUser)
	if err != nil {
		slog.Warn(err.Error(), "handler", "UpdateUserHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "User updated successfully"})
}
