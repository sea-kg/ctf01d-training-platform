package api_helpers

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

// fixme тут ли ей место? вынести в - tool
func HashPassword(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	// fixme добавить соль в конфиг и солить пароли, переделать на bcrypt
	return hex.EncodeToString(h.Sum(nil))
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
