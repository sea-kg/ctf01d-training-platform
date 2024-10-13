package helper

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"ctf01d/internal/httpserver"

	openapi_types "github.com/oapi-codegen/runtime/types"
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

func ConvertUserRequestRoleToUserResponseRole(role httpserver.UserRequestRole) httpserver.UserResponseRole {
	switch role {
	case httpserver.UserRequestRoleAdmin:
		return httpserver.UserResponseRoleAdmin
	case httpserver.UserRequestRoleGuest:
		return httpserver.UserResponseRoleGuest
	case httpserver.UserRequestRolePlayer:
		return httpserver.UserResponseRolePlayer
	default:
		return ""
	}
}

func ConvertUserRequestRoleToString(role httpserver.UserRequestRole) string {
	switch role {
	case httpserver.UserRequestRoleAdmin:
		return "admin"
	case httpserver.UserRequestRoleGuest:
		return "guest"
	case httpserver.UserRequestRolePlayer:
		return "player"
	default:
		return ""
	}
}

func ConvertUserResponseRoleToUserRequestRole(role httpserver.UserResponseRole) httpserver.UserRequestRole {
	switch role {
	case httpserver.UserResponseRoleAdmin:
		return httpserver.UserRequestRoleAdmin
	case httpserver.UserResponseRoleGuest:
		return httpserver.UserRequestRoleGuest
	case httpserver.UserResponseRolePlayer:
		return httpserver.UserRequestRolePlayer
	default:
		return ""
	}
}

var zipSignature = []byte{0x50, 0x4b, 0x03, 0x04}

func IsZip(data []byte) bool {
	return len(data) >= 4 && bytes.Equal(zipSignature, data[:4])
}

func Unhyphensify(id openapi_types.UUID) string {
	var sb strings.Builder
	str := id.String()
	sb.WriteString(str[0:8])
	sb.WriteString(str[8:12])
	sb.WriteString(str[12:16])
	sb.WriteString(str[16:20])
	sb.WriteString(str[20:32])

	return sb.String()
}

func IdToPath(id openapi_types.UUID) string {
	raw := Unhyphensify(id)
	const rank = 8
	size := len(raw) / rank
	var path string
	for i := 0; i < rank; i++ {
		base := i * size
		path = filepath.Join(path, raw[base:base+size])
	}

	return path
}
