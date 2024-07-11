package helpers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"ctf01d/internal/app/server"

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

func PrepareImage(avatarUrl string) string {
	// fixme подумать за генерацию аватарок, пока для mvp - сойдет robohash.org
	if strings.Contains(avatarUrl, "robohash.org") {
		return avatarUrl
	}
	// fixme подумать что делать с http контентом
	re := regexp.MustCompile(`(?i)^https?://.*\.(jpg|jpeg|png|gif)$`)
	if re.MatchString(avatarUrl) {
		return avatarUrl
	}
	return "https://robohash.org/" + url.QueryEscape(avatarUrl)
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
