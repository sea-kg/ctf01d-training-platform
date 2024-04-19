package routers

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"ctf01d/internal/app/api"
	"ctf01d/internal/app/logger"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc HandlerFunc
}

type HandlerFunc func(db *sql.DB, w http.ResponseWriter, r *http.Request)

type Routes []Route

func NewRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		handler := func(w http.ResponseWriter, r *http.Request) {
			route.HandlerFunc(db, w, r)
		}
		loggedHandler := logger.Logger(http.HandlerFunc(handler), route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(loggedHandler)
	}

	return router
}

var routes = Routes{
	Route{"CreateGame", strings.ToUpper("Post"), "/api/games", api.CreateGameHandler},
	Route{"DeleteGame", strings.ToUpper("Delete"), "/api/games/{id}", api.DeleteGameHandler},
	Route{"GetGameById", strings.ToUpper("Get"), "/api/games/{id}", api.GetGameByIdHandler},
	Route{"ListGames", strings.ToUpper("Get"), "/api/games", api.ListGamesHandler},
	Route{"UpdateGame", strings.ToUpper("Put"), "/api/games/{id}", api.UpdateGameHandler},

	Route{"CreateResult", strings.ToUpper("Post"), "/api/results", api.CreateResultHandler},
	Route{"GetResultById", strings.ToUpper("Get"), "/api/results/{id}", api.GetResultByIdHandler},
	Route{"ListResults", strings.ToUpper("Get"), "/api/results", api.ListResultsHandler},

	Route{"CreateService", strings.ToUpper("Post"), "/api/services", api.CreateServiceHandler},
	Route{"DeleteService", strings.ToUpper("Delete"), "/api/services/{id}", api.DeleteServiceHandler},
	Route{"GetServiceById", strings.ToUpper("Get"), "/api/services/{id}", api.GetServiceByIdHandler},
	Route{"ListServices", strings.ToUpper("Get"), "/api/services", api.ListServicesHandler},
	Route{"UpdateService", strings.ToUpper("Put"), "/api/services/{id}", api.UpdateServiceHandler},

	Route{"CreateTeam", strings.ToUpper("Post"), "/api/teams", api.CreateTeamHandler},
	Route{"DeleteTeam", strings.ToUpper("Delete"), "/api/teams/{id}", api.DeleteTeamHandler},
	Route{"GetTeamById", strings.ToUpper("Get"), "/api/teams/{id}", api.GetTeamByIdHandler},
	Route{"ListTeams", strings.ToUpper("Get"), "/api/teams", api.ListTeamsHandler},
	Route{"UpdateTeam", strings.ToUpper("Put"), "/api/teams/{id}", api.UpdateTeamHandler},

	Route{"CreateUser", strings.ToUpper("Post"), "/api/users", api.CreateUserHandler},
	Route{"DeleteUser", strings.ToUpper("Delete"), "/api/users/{id}", api.DeleteUserHandler},
	Route{"GetUserById", strings.ToUpper("Get"), "/api/users/{id}", api.GetUserByIdHandler},
	Route{"ListUsers", strings.ToUpper("Get"), "/api/users", api.ListUsersHandler},
	Route{"UpdateUser", strings.ToUpper("Put"), "/api/users/{id}", api.UpdateUserHandler},
}
