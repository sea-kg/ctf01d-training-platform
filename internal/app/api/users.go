package api

import (
	"crypto/sha1"
	"ctf01d/internal/app/models"
	"ctf01d/internal/app/repository"
	"ctf01d/internal/app/view"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type RequestUser struct {
	Username  string `json:"user_name"`
	Role      string `json:"role"`
	AvatarUrl string `json:"avatar_url"`
	Status    string `json:"status"`
	Password  string `json:"password"`
}

func CreateUserHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var user RequestUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	userRepo := repository.NewUserRepository(db)
	newUser := &models.User{
		Username:     user.Username,
		Role:         user.Role,
		Status:       user.Status,
		PasswordHash: HashPassword(user.Password),
	}
	if err := userRepo.Create(r.Context(), newUser); err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create user: " + err.Error()})
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"data": "User created successfully"})
}

// ей тут не место, вынести в - tool, добавить соль в конфиг и солить пароли
func HashPassword(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func DeleteUserHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	userRepo := repository.NewUserRepository(db)
	if err := userRepo.Delete(r.Context(), id); err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"data": "User deleted successfully"})
}

func GetUserByIdHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	userRepo := repository.NewUserRepository(db)
	user, err := userRepo.GetById(r.Context(), id)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch user"})
		return
	}
	respondWithJSON(w, http.StatusOK, view.NewUserFromModel(user))
}

func ListUsersHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	userRepo := repository.NewUserRepository(db)
	users, err := userRepo.List(r.Context())
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	respondWithJSON(w, http.StatusOK, view.NewUsersFromModels(users))
}

func UpdateUserHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// fixme update не проверяет есть ли запись в бд
	var user RequestUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	userRepo := repository.NewUserRepository(db)
	updateUser := &models.User{
		Username: user.Username,
		Role:     user.Role,
		Status:   user.Status,
	}
	vars := mux.Vars(r)
	id := vars["id"]
	updateUser.Id = id
	err := userRepo.Update(r.Context(), updateUser)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"data": "User updated successfully"})
}
