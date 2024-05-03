package web

import (
	helpers "ctf01d/internal/app/utils"
	"net/http"
)

func ListGameHandler(w http.ResponseWriter, r *http.Request) {
	helpers.RenderTemplate(w, "games/index.html")
}
func ListUserHandler(w http.ResponseWriter, r *http.Request) {
	helpers.RenderTemplate(w, "users/index.html")
}
func TeamUserHandler(w http.ResponseWriter, r *http.Request) {
	helpers.RenderTemplate(w, "teams/index.html")
}
func ServiceUserHandler(w http.ResponseWriter, r *http.Request) {
	helpers.RenderTemplate(w, "services/index.html")
}
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	helpers.RenderTemplate(w, "index.html")
}
