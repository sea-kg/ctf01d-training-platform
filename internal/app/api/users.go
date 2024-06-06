package api

import (
	apimodels "ctf01d/internal/app/apimodels"
	dbmodels "ctf01d/internal/app/db"
	"ctf01d/internal/app/repository"
	api_helpers "ctf01d/internal/app/utils"
	"ctf01d/internal/app/view"
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateUserHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// fixme обернуть в транзакцию, т.к. две вставки подряд
	var user apimodels.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		slog.Warn(err.Error(), "handler", "CreateUserHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	repo := repository.NewUserRepository(db)
	passwordHash, err := api_helpers.HashPassword(*user.Password)
	slog.Info("user.password " + passwordHash)
	if err != nil {
		slog.Warn(err.Error(), "handler", "CreateUserHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	newUser := &dbmodels.User{
		Username:     *user.UserName,
		Role:         string(*user.Role),
		Status:       *user.Status,
		PasswordHash: passwordHash,
		AvatarUrl:    api_helpers.PrepareImage(*user.AvatarUrl),
	}
	if err := repo.Create(r.Context(), newUser); err != nil {
		slog.Warn(err.Error(), "handler", "CreateUserHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
		return
	}
	if len(*user.TeamIds) > 0 {
		if err := repo.AddUserToTeams(r.Context(), newUser.Id, user.TeamIds); err != nil {
			slog.Warn(err.Error(), "handler", "CreateUserHandler")
			api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to add user to teams"})
			return
		}
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "User created successfully"})
}

func DeleteUserHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		slog.Warn(err.Error(), "handler", "DeleteUserHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Bad request"})
		return
	}
	repo := repository.NewUserRepository(db)
	if err := repo.Delete(r.Context(), id); err != nil {
		slog.Warn(err.Error(), "handler", "DeleteUserHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "User deleted successfully"})
}

func GetUserByIdHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetUserByIdHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Bad request"})
		return
	}
	repo := repository.NewUserRepository(db)
	user, err := repo.GetById(r.Context(), id)
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetUserByIdHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch user"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewUserFromModel(user))
}

func ListUsersHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	repo := repository.NewUserRepository(db)
	users, err := repo.List(r.Context())
	if err != nil {
		slog.Warn(err.Error(), "handler", "ListUsersHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewUsersFromModels(users))
}

func UpdateUserHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// fixme update не проверяет есть ли запись в бд
	var user apimodels.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateUserHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	passwordHash, err := api_helpers.HashPassword(*user.Password)
	if err != nil {
		slog.Warn(err.Error(), "handler", "UpdateUserHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	repo := repository.NewUserRepository(db)
	updateUser := &dbmodels.User{
		Username:     *user.UserName,
		Role:         string(*user.Role),
		Status:       *user.Status,
		PasswordHash: passwordHash,
		AvatarUrl:    api_helpers.PrepareImage(*user.AvatarUrl),
	}
	vars := mux.Vars(r)
	id, err2 := strconv.Atoi(vars["id"])
	if err2 != nil {
		slog.Warn(err2.Error(), "handler", "UpdateUserHandler")
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": err2.Error()})
		return
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
