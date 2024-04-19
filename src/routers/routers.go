package routers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"ctf01d/lib/api"
	"ctf01d/lib/logger"
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

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{"CreateGame", strings.ToUpper("Post"), "/games", api.CreateGameHandler},
	Route{"DeleteGame", strings.ToUpper("Delete"), "/games/{id}", api.DeleteGameHandler},
	Route{"GetGameById", strings.ToUpper("Get"), "/games/{id}", api.GetGameByIdHandler},
	Route{"ListGames", strings.ToUpper("Get"), "/games", api.ListGamesHandler},
	Route{"UpdateGame", strings.ToUpper("Put"), "/games/{id}", api.UpdateGameHandler},

	Route{"CreateResult", strings.ToUpper("Post"), "/results", api.CreateResultHandler},
	Route{"GetResultById", strings.ToUpper("Get"), "/results/{id}", api.GetResultByIdHandler},
	Route{"ListResults", strings.ToUpper("Get"), "/results", api.ListResultsHandler},

	Route{"CreateService", strings.ToUpper("Post"), "/services", api.CreateServiceHandler},
	Route{"DeleteService", strings.ToUpper("Delete"), "/services/{id}", api.DeleteServiceHandler},
	Route{"GetServiceById", strings.ToUpper("Get"), "/services/{id}", api.GetServiceByIdHandler},
	Route{"ListServices", strings.ToUpper("Get"), "/services", api.ListServicesHandler},
	Route{"UpdateService", strings.ToUpper("Put"), "/services/{id}", api.UpdateServiceHandler},

	Route{"CreateTeam", strings.ToUpper("Post"), "/teams", api.CreateTeamHandler},
	Route{"DeleteTeam", strings.ToUpper("Delete"), "/teams/{id}", api.DeleteTeamHandler},
	Route{"GetTeamById", strings.ToUpper("Get"), "/teams/{id}", api.GetTeamByIdHandler},
	Route{"ListTeams", strings.ToUpper("Get"), "/teams", api.ListTeamsHandler},
	Route{"UpdateTeam", strings.ToUpper("Put"), "/teams/{id}", api.UpdateTeamHandler},

	Route{"CreateUser", strings.ToUpper("Post"), "/users", api.CreateUserHandler},
	Route{"DeleteUser", strings.ToUpper("Delete"), "/users/{id}", api.DeleteUserHandler},
	Route{"GetUserById", strings.ToUpper("Get"), "/users/{id}", api.GetUserByIdHandler},
	Route{"ListUsers", strings.ToUpper("Get"), "/users", api.ListUsersHandler},
	Route{"UpdateUser", strings.ToUpper("Put"), "/users/{id}", api.UpdateUserHandler},
}
