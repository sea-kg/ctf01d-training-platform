package api_helpers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"regexp"

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
	// fixme подумать что делать с http контентом
	re := regexp.MustCompile(`(?i)^https?://.*\.(jpg|jpeg|png|gif)$`)
	if re.MatchString(avatarUrl) {
		return avatarUrl
	}
	// fixme подумать за генерацию аватарок, пока для mvp - сойдет
	return "https://robohash.org/" + url.QueryEscape(avatarUrl)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(`{"error": "Error marshaling the response object"}`)); err != nil {
			log.Printf("Error writing error response: %v", err)
		}
		return
	}
	w.WriteHeader(code)
	if _, err := w.Write(response); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
