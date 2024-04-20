package api

import (
	"ctf01d/internal/app/models"
	"ctf01d/internal/app/repository"
	api_helpers "ctf01d/internal/app/utils"
	"ctf01d/internal/app/view"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type RequestUser struct {
	Username  string   `json:"user_name"`
	Role      string   `json:"role"`
	AvatarUrl string   `json:"avatar_url"`
	Status    string   `json:"status"`
	Password  string   `json:"password"`
	TeamsId   []string `json:"team_ids"`
}

func CreateUserHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// fixme обернуть в транзакцию, т.к. две вставки подряд
	var user RequestUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload: " + err.Error()})
		return
	}
	userRepo := repository.NewUserRepository(db)
	newUser := &models.User{
		Username:     user.Username,
		Role:         user.Role,
		Status:       user.Status,
		PasswordHash: api_helpers.HashPassword(user.Password),
		AvatarUrl:    api_helpers.PrepareImage(user.AvatarUrl),
	}
	if err := userRepo.Create(r.Context(), newUser); err != nil {
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create user: " + err.Error()})
		return
	}
	if len(user.TeamsId) > 0 {
		if err := userRepo.AddUserToTeams(r.Context(), newUser.Id, user.TeamsId); err != nil {
			api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to add user to teams: " + err.Error()})
			return
		}
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "User created successfully"})
}

func DeleteUserHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	userRepo := repository.NewUserRepository(db)
	if err := userRepo.Delete(r.Context(), id); err != nil {
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "User deleted successfully"})
}

func GetUserByIdHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	userRepo := repository.NewUserRepository(db)
	user, err := userRepo.GetById(r.Context(), id)
	if err != nil {
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch user"})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewUserFromModel(user))
}

func ListUsersHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	userRepo := repository.NewUserRepository(db)
	users, err := userRepo.List(r.Context())
	if err != nil {
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, view.NewUsersFromModels(users))
}

func UpdateUserHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// fixme update не проверяет есть ли запись в бд
	var user RequestUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		api_helpers.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	userRepo := repository.NewUserRepository(db)
	updateUser := &models.User{
		Username: user.Username,
		Role:     user.Role,
		Status:   user.Status,
	}
	vars := mux.Vars(r)
	id, err2 := strconv.Atoi(vars["id"])
	if err2 != nil {
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": err2.Error()})
		return
	}
	updateUser.Id = id
	err := userRepo.Update(r.Context(), updateUser)
	if err != nil {
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "User updated successfully"})
}
