package routers

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"ctf01d/internal/app/api"
	"ctf01d/internal/app/logger"
	"ctf01d/internal/app/web"
)

type ApiRoutes []ApiRoute
type ApiRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc HandlerFunc
}

type FrontRoutes []FrontRoute
type FrontRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type HandlerFunc func(db *sql.DB, w http.ResponseWriter, r *http.Request)

func NewRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// api backend
	for _, route := range apiRoutes {
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
	// frontend kek
	for _, route := range frontRoutes {
		handler := func(w http.ResponseWriter, r *http.Request) {
			route.HandlerFunc(w, r)
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

var apiRoutes = ApiRoutes{
	ApiRoute{"CreateGame", strings.ToUpper("Post"), "/api/games", api.CreateGameHandler},
	ApiRoute{"DeleteGame", strings.ToUpper("Delete"), "/api/games/{id}", api.DeleteGameHandler},
	ApiRoute{"GetGameById", strings.ToUpper("Get"), "/api/games/{id}", api.GetGameByIdHandler},
	ApiRoute{"ListGames", strings.ToUpper("Get"), "/api/games", api.ListGamesHandler},
	ApiRoute{"UpdateGame", strings.ToUpper("Put"), "/api/games/{id}", api.UpdateGameHandler},

	ApiRoute{"CreateResult", strings.ToUpper("Post"), "/api/results", api.CreateResultHandler},
	ApiRoute{"GetResultById", strings.ToUpper("Get"), "/api/results/{id}", api.GetResultByIdHandler},
	ApiRoute{"ListResults", strings.ToUpper("Get"), "/api/results", api.ListResultsHandler},

	ApiRoute{"CreateService", strings.ToUpper("Post"), "/api/services", api.CreateServiceHandler},
	ApiRoute{"DeleteService", strings.ToUpper("Delete"), "/api/services/{id}", api.DeleteServiceHandler},
	ApiRoute{"GetServiceById", strings.ToUpper("Get"), "/api/services/{id}", api.GetServiceByIdHandler},
	ApiRoute{"ListServices", strings.ToUpper("Get"), "/api/services", api.ListServicesHandler},
	ApiRoute{"UpdateService", strings.ToUpper("Put"), "/api/services/{id}", api.UpdateServiceHandler},

	ApiRoute{"CreateTeam", strings.ToUpper("Post"), "/api/teams", api.CreateTeamHandler},
	ApiRoute{"DeleteTeam", strings.ToUpper("Delete"), "/api/teams/{id}", api.DeleteTeamHandler},
	ApiRoute{"GetTeamById", strings.ToUpper("Get"), "/api/teams/{id}", api.GetTeamByIdHandler},
	ApiRoute{"ListTeams", strings.ToUpper("Get"), "/api/teams", api.ListTeamsHandler},
	ApiRoute{"UpdateTeam", strings.ToUpper("Put"), "/api/teams/{id}", api.UpdateTeamHandler},

	ApiRoute{"CreateUser", strings.ToUpper("Post"), "/api/users", api.CreateUserHandler},
	ApiRoute{"DeleteUser", strings.ToUpper("Delete"), "/api/users/{id}", api.DeleteUserHandler},
	ApiRoute{"GetUserById", strings.ToUpper("Get"), "/api/users/{id}", api.GetUserByIdHandler},
	ApiRoute{"ListUsers", strings.ToUpper("Get"), "/api/users", api.ListUsersHandler},
	ApiRoute{"UpdateUser", strings.ToUpper("Put"), "/api/users/{id}", api.UpdateUserHandler},

	ApiRoute{"LoginUser", strings.ToUpper("Post"), "/api/login", api.LoginSessionHandler},
	ApiRoute{"LogoutUser", strings.ToUpper("Post"), "/api/logout", api.LogoutSessionHandler},

	ApiRoute{"ListUniversities", strings.ToUpper("Get"), "/api/universities", api.ListUniversitiesHandler},
}

var frontRoutes = FrontRoutes{
	FrontRoute{"ListGame", strings.ToUpper("Get"), "/games/index.html", web.ListGameHandler},
	FrontRoute{"ListUser", strings.ToUpper("Get"), "/users/index.html", web.ListUserHandler},
	FrontRoute{"TeamUser", strings.ToUpper("Get"), "/teams/index.html", web.TeamUserHandler},
	FrontRoute{"ServiceUser", strings.ToUpper("Get"), "/services/index.html", web.ServiceUserHandler},
	FrontRoute{"Index", strings.ToUpper("Get"), "/", web.IndexHandler},
}
