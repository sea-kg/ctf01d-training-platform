package handlers

import (
	"ctf01d/internal/app/server"
	"net/http"

	openapi_types "github.com/oapi-codegen/runtime/types"
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

func (siw *ServerInterfaceWrapper) DeleteGame(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	siw.handlers.DeleteGame(w, r, id)
}

func (siw *ServerInterfaceWrapper) GetGameById(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	siw.handlers.GetGameById(w, r, id)
}

func (siw *ServerInterfaceWrapper) UpdateGame(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	siw.handlers.UpdateGame(w, r, id)
}

func (siw *ServerInterfaceWrapper) PostApiV1AuthSignin(w http.ResponseWriter, r *http.Request) {
	siw.handlers.PostApiV1AuthSignin(w, r)
}

func (siw *ServerInterfaceWrapper) PostApiV1AuthSignout(w http.ResponseWriter, r *http.Request) {
	siw.handlers.PostApiV1AuthSignout(w, r)
}

func (siw *ServerInterfaceWrapper) ValidateSession(w http.ResponseWriter, r *http.Request) {
	siw.handlers.ValidateSession(w, r)
}

func (siw *ServerInterfaceWrapper) ListResults(w http.ResponseWriter, r *http.Request) {
	siw.handlers.ListResults(w, r)
}

func (siw *ServerInterfaceWrapper) CreateResult(w http.ResponseWriter, r *http.Request) {
	siw.handlers.CreateResult(w, r)
}

func (siw *ServerInterfaceWrapper) GetResultById(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	siw.handlers.GetResultById(w, r, id)
}

func (siw *ServerInterfaceWrapper) ListServices(w http.ResponseWriter, r *http.Request) {
	siw.handlers.ListServices(w, r)
}

func (siw *ServerInterfaceWrapper) CreateService(w http.ResponseWriter, r *http.Request) {
	siw.handlers.CreateService(w, r)
}

func (siw *ServerInterfaceWrapper) DeleteService(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	siw.handlers.DeleteService(w, r, id)
}

func (siw *ServerInterfaceWrapper) GetServiceById(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	siw.handlers.GetServiceById(w, r, id)
}

func (siw *ServerInterfaceWrapper) UpdateService(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	siw.handlers.UpdateService(w, r, id)
}

func (siw *ServerInterfaceWrapper) ListTeams(w http.ResponseWriter, r *http.Request) {
	siw.handlers.ListTeams(w, r)
}

func (siw *ServerInterfaceWrapper) CreateTeam(w http.ResponseWriter, r *http.Request) {
	siw.handlers.CreateTeam(w, r)
}

func (siw *ServerInterfaceWrapper) DeleteTeam(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	siw.handlers.DeleteTeam(w, r, id)
}

func (siw *ServerInterfaceWrapper) GetTeamById(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	siw.handlers.GetTeamById(w, r, id)
}

func (siw *ServerInterfaceWrapper) UpdateTeam(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	siw.handlers.UpdateTeam(w, r, id)
}

func (siw *ServerInterfaceWrapper) GetApiV1Universities(w http.ResponseWriter, r *http.Request, params server.GetApiV1UniversitiesParams) {
	siw.handlers.GetApiV1Universities(w, r, params)
}

func (siw *ServerInterfaceWrapper) ListUsers(w http.ResponseWriter, r *http.Request) {
	siw.handlers.ListUsers(w, r)
}

func (siw *ServerInterfaceWrapper) CreateUser(w http.ResponseWriter, r *http.Request) {
	siw.handlers.CreateUser(w, r)
}

func (siw *ServerInterfaceWrapper) DeleteUser(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	siw.handlers.DeleteUser(w, r, id)
}

func (siw *ServerInterfaceWrapper) GetUserById(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	siw.handlers.GetUserById(w, r, id)
}

func (siw *ServerInterfaceWrapper) UpdateUser(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	siw.handlers.UpdateUser(w, r, id)
}
