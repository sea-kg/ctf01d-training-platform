package middleware

import (
	"net/http"

	"ctf01d/internal/httpserver"
	"ctf01d/internal/repository"
)

func Auth(repo repository.SessionRepository) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Проверяем наличие признака в контексте
			if scopes, ok := r.Context().Value(httpserver.SessionCookieScopes).([]string); ok && len(scopes) > 0 {
				// Достаём сессию
				cookie, err := r.Cookie("session_id")
				if err != nil {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}

				// Проверяем сессию в базе данных
				_, err = repo.GetSessionFromDB(r.Context(), cookie.Value)
				if err != nil {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}
