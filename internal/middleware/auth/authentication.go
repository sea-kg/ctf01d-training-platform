package auth

import (
	"database/sql"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"ctf01d/internal/httpserver"
)

func AuthenticationMiddleware(db *sql.DB) httpserver.MiddlewareFunc {
	return func(c *gin.Context) {
		// Проверка, нужен ли вообще роуту токен для авторизации
		sessionScopes, sessionExists := c.Keys["sessionAuth.Scopes"]

		// Если оба отсутствуют, пропускаем запрос дальше
		if !sessionExists {
			c.Next()
			return
		}

		// Session Authentication
		if sessionExists {
			session := sessions.Default(c)
			if userID := session.Get("user_id"); userID != nil {
				c.Set("user_id", userID)
				c.Set("auth_type", "session")
				c.Set("scopes", sessionScopes)
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized", "message": "no valid authentication method found"})
		c.Abort()
	}
}
