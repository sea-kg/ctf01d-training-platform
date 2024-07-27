package handlers

import (
	"net/http"

	"ctf01d/internal/app/server"

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

func (siw *ServerInterfaceWrapper) SignInUser(w http.ResponseWriter, r *http.Request) {
	siw.handlers.SignInUser(w, r)
}

func (siw *ServerInterfaceWrapper) SignOutUser(w http.ResponseWriter, r *http.Request) {
	siw.handlers.SignOutUser(w, r)
}

func (siw *ServerInterfaceWrapper) ValidateSession(w http.ResponseWriter, r *http.Request) {
	siw.handlers.ValidateSession(w, r)
}

func (siw *ServerInterfaceWrapper) GetResult(w http.ResponseWriter, r *http.Request, gameId openapi_types.UUID) {
	siw.handlers.GetResult(w, r, gameId)
}

func (siw *ServerInterfaceWrapper) CreateResult(w http.ResponseWriter, r *http.Request, gameId openapi_types.UUID) {
	siw.handlers.CreateResult(w, r, gameId)
}

func (siw *ServerInterfaceWrapper) UpdateResult(w http.ResponseWriter, r *http.Request, gameId openapi_types.UUID) {
	siw.handlers.CreateResult(w, r, gameId)
}

func (siw *ServerInterfaceWrapper) GetScoreboard(w http.ResponseWriter, r *http.Request, gameId openapi_types.UUID) {
	siw.handlers.GetScoreboard(w, r, gameId)
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

func (siw *ServerInterfaceWrapper) ConnectUserTeam(w http.ResponseWriter, r *http.Request, teamId openapi_types.UUID, userId openapi_types.UUID) {
	siw.handlers.ConnectUserTeam(w, r, teamId, userId)
}

func (siw *ServerInterfaceWrapper) LeaveUserFromTeam(w http.ResponseWriter, r *http.Request, teamId openapi_types.UUID, userId openapi_types.UUID) {
	siw.handlers.LeaveUserFromTeam(w, r, teamId, userId)
}

func (siw *ServerInterfaceWrapper) ApproveUserTeam(w http.ResponseWriter, r *http.Request, teamId openapi_types.UUID, userId openapi_types.UUID) {
	siw.handlers.ApproveUserTeam(w, r, teamId, userId)
}

func (siw *ServerInterfaceWrapper) TeamMembers(w http.ResponseWriter, r *http.Request, teamId openapi_types.UUID) {
	siw.handlers.TeamMembers(w, r, teamId)
}

func (siw *ServerInterfaceWrapper) ListUniversities(w http.ResponseWriter, r *http.Request, params server.ListUniversitiesParams) {
	siw.handlers.ListUniversities(w, r, params)
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

func (siw *ServerInterfaceWrapper) GetProfileById(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	siw.handlers.GetProfileById(w, r, id)
}

func (siw *ServerInterfaceWrapper) UpdateUser(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	siw.handlers.UpdateUser(w, r, id)
}

func (siw *ServerInterfaceWrapper) UploadChecker(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	siw.handlers.UploadChecker(w, r, id)
}

func (siw *ServerInterfaceWrapper) UploadService(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	siw.handlers.UploadService(w, r, id)
}

func (siw *ServerInterfaceWrapper) UniqueAvatar(w http.ResponseWriter, r *http.Request, username string) {
	siw.handlers.UniqueAvatar(w, r, username)
}
