package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	dbmodels "ctf01d/internal/app/db"
	"ctf01d/internal/app/repository"
	"ctf01d/internal/app/server"
	api_helpers "ctf01d/internal/app/utils"
	"ctf01d/internal/app/view"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {
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
	newUser := &dbmodels.User{
		Username:     user.UserName,
		DisplayName:  *user.DisplayName,
		Role:         user.Role,
		Status:       user.Status,
		PasswordHash: passwordHash,
		AvatarUrl:    api_helpers.PrepareImage(*user.AvatarUrl),
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
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewUserFromModel(newUser))

}

func (h *Handlers) DeleteUser(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	repo := repository.NewUserRepository(h.DB)
	if err := repo.Delete(r.Context(), id); err != nil {
		slog.Warn(err.Error(), "handler", "DeleteUserHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "User deleted successfully"})
}

func (h *Handlers) GetUserById(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	repo := repository.NewUserRepository(h.DB)
	user, err := repo.GetById(r.Context(), id)
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetUserByIdHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch user"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewUserFromModel(user))
}

func (h *Handlers) GetProfileById(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	repo := repository.NewUserRepository(h.DB)
	userProfile, err := repo.GetProfileWithHistory(r.Context(), id)
	if err != nil {
		slog.Info(err.Error(), "handler", "GetProfileByIdHandler")
		api_helpers.RespondWithJSON(w, http.StatusNotFound, map[string]string{"data": "User have not profile"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewProfileFromModel(userProfile))
}

func (h *Handlers) ListUsers(w http.ResponseWriter, r *http.Request) {
	repo := repository.NewUserRepository(h.DB)
	users, err := repo.List(r.Context())
	if err != nil {
		slog.Warn(err.Error(), "handler", "ListUsersHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewUsersFromModels(users))
}

func (h *Handlers) UpdateUser(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
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
	updateUser := &dbmodels.User{
		Username:     user.UserName,
		DisplayName:  *user.DisplayName,
		Role:         user.Role,
		Status:       user.Status,
		PasswordHash: passwordHash,
		AvatarUrl:    api_helpers.PrepareImage(*user.AvatarUrl),
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
