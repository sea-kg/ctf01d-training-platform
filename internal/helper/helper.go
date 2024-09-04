package helper

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"

	"ctf01d/internal/server"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPasswordHash(s, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(s))
	return err == nil
}

func ToNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{
			String: "",
			Valid:  false,
		}
	}
	return sql.NullString{
		String: *s,
		Valid:  true,
	}
}

func WithDefault(img string) string {
	return "api/v1/avatar/" + url.QueryEscape(img)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(`{"error": "Error marshaling the response object"}`)); err != nil {
			slog.Error("Error writing error response: " + err.Error())
		}
		return
	}
	w.WriteHeader(code)
	if _, err := w.Write(response); err != nil {
		slog.Error("Error writing response: " + err.Error())
	}
}

func ConvertUserRequestRoleToUserResponseRole(role server.UserRequestRole) server.UserResponseRole {
	switch role {
	case server.UserRequestRoleAdmin:
		return server.UserResponseRoleAdmin
	case server.UserRequestRoleGuest:
		return server.UserResponseRoleGuest
	case server.UserRequestRolePlayer:
		return server.UserResponseRolePlayer
	default:
		return ""
	}
}

func ConvertUserRequestRoleToString(role server.UserRequestRole) string {
	switch role {
	case server.UserRequestRoleAdmin:
		return "admin"
	case server.UserRequestRoleGuest:
		return "guest"
	case server.UserRequestRolePlayer:
		return "player"
	default:
		return ""
	}
}

func ConvertUserResponseRoleToUserRequestRole(role server.UserResponseRole) server.UserRequestRole {
	switch role {
	case server.UserResponseRoleAdmin:
		return server.UserRequestRoleAdmin
	case server.UserResponseRoleGuest:
		return server.UserRequestRoleGuest
	case server.UserResponseRolePlayer:
		return server.UserRequestRolePlayer
	default:
		return ""
	}
}
