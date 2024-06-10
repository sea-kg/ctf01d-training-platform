package handlers

import (
	"ctf01d/internal/app/server"
	"net/http"
)

// ServerInterfaceWrapper wraps Handlers to conform to the generated interface
type ServerInterfaceWrapper struct {
	handlers *Handlers
}

func NewServerInterfaceWrapper(handlers *Handlers) *ServerInterfaceWrapper {
	return &ServerInterfaceWrapper{handlers: handlers}
}

func (siw *ServerInterfaceWrapper) ListGames(w http.ResponseWriter, r *http.Request) {
	siw.handlers.ListGames(w, r)
}

func (siw *ServerInterfaceWrapper) CreateGame(w http.ResponseWriter, r *http.Request) {
	siw.handlers.CreateGame(w, r)
}

func (siw *ServerInterfaceWrapper) DeleteGame(w http.ResponseWriter, r *http.Request, id int) {
	siw.handlers.DeleteGame(w, r, id)
}

func (siw *ServerInterfaceWrapper) GetGameById(w http.ResponseWriter, r *http.Request, id int) {
	siw.handlers.GetGameById(w, r, id)
}

func (siw *ServerInterfaceWrapper) UpdateGame(w http.ResponseWriter, r *http.Request, id int) {
	siw.handlers.UpdateGame(w, r, id)
}

func (siw *ServerInterfaceWrapper) PostApiLogin(w http.ResponseWriter, r *http.Request) {
	siw.handlers.PostApiLogin(w, r)
}

func (siw *ServerInterfaceWrapper) PostApiLogout(w http.ResponseWriter, r *http.Request) {
	siw.handlers.PostApiLogout(w, r)
}

func (siw *ServerInterfaceWrapper) ListResults(w http.ResponseWriter, r *http.Request) {
	siw.handlers.ListResults(w, r)
}

func (siw *ServerInterfaceWrapper) CreateResult(w http.ResponseWriter, r *http.Request) {
	siw.handlers.CreateResult(w, r)
}

func (siw *ServerInterfaceWrapper) GetResultById(w http.ResponseWriter, r *http.Request, id int) {
	siw.handlers.GetResultById(w, r, id)
}

func (siw *ServerInterfaceWrapper) ListServices(w http.ResponseWriter, r *http.Request) {
	siw.handlers.ListServices(w, r)
}

func (siw *ServerInterfaceWrapper) CreateService(w http.ResponseWriter, r *http.Request) {
	siw.handlers.CreateService(w, r)
}

func (siw *ServerInterfaceWrapper) DeleteService(w http.ResponseWriter, r *http.Request, id int) {
	siw.handlers.DeleteService(w, r, id)
}

func (siw *ServerInterfaceWrapper) GetServiceById(w http.ResponseWriter, r *http.Request, id int) {
	siw.handlers.GetServiceById(w, r, id)
}

func (siw *ServerInterfaceWrapper) UpdateService(w http.ResponseWriter, r *http.Request, id int) {
	siw.handlers.UpdateService(w, r, id)
}

func (siw *ServerInterfaceWrapper) ListTeams(w http.ResponseWriter, r *http.Request) {
	siw.handlers.ListTeams(w, r)
}

func (siw *ServerInterfaceWrapper) CreateTeam(w http.ResponseWriter, r *http.Request) {
	siw.handlers.CreateTeam(w, r)
}

func (siw *ServerInterfaceWrapper) DeleteTeam(w http.ResponseWriter, r *http.Request, id int) {
	siw.handlers.DeleteTeam(w, r, id)
}

func (siw *ServerInterfaceWrapper) GetTeamById(w http.ResponseWriter, r *http.Request, id int) {
	siw.handlers.GetTeamById(w, r, id)
}

func (siw *ServerInterfaceWrapper) UpdateTeam(w http.ResponseWriter, r *http.Request, id int) {
	siw.handlers.UpdateTeam(w, r, id)
}

func (siw *ServerInterfaceWrapper) GetApiUniversities(w http.ResponseWriter, r *http.Request, params server.GetApiUniversitiesParams) {
	siw.handlers.GetApiUniversities(w, r, params)
}

func (siw *ServerInterfaceWrapper) ListUsers(w http.ResponseWriter, r *http.Request) {
	siw.handlers.ListUsers(w, r)
}

func (siw *ServerInterfaceWrapper) CreateUser(w http.ResponseWriter, r *http.Request) {
	siw.handlers.CreateUser(w, r)
}

func (siw *ServerInterfaceWrapper) DeleteUser(w http.ResponseWriter, r *http.Request, id int) {
	siw.handlers.DeleteUser(w, r, id)
}

func (siw *ServerInterfaceWrapper) GetUserById(w http.ResponseWriter, r *http.Request, id int) {
	siw.handlers.GetUserById(w, r, id)
}

func (siw *ServerInterfaceWrapper) UpdateUser(w http.ResponseWriter, r *http.Request, id int) {
	siw.handlers.UpdateUser(w, r, id)
}