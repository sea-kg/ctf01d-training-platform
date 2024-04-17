package routers

import (
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
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"CreateGame",
		strings.ToUpper("Post"),
		"/games",
		api.CreateGame,
	},

	Route{
		"DeleteGame",
		strings.ToUpper("Delete"),
		"/games/{id}",
		api.DeleteGame,
	},

	Route{
		"GetGameById",
		strings.ToUpper("Get"),
		"/games/{id}",
		api.GetGameById,
	},

	Route{
		"ListGames",
		strings.ToUpper("Get"),
		"/games",
		api.ListGames,
	},

	Route{
		"UpdateGame",
		strings.ToUpper("Put"),
		"/games/{id}",
		api.UpdateGame,
	},

	Route{
		"CreateResult",
		strings.ToUpper("Post"),
		"/results",
		api.CreateResult,
	},

	Route{
		"GetResultById",
		strings.ToUpper("Get"),
		"/results/{id}",
		api.GetResultById,
	},

	Route{
		"ListResults",
		strings.ToUpper("Get"),
		"/results",
		api.ListResults,
	},

	Route{
		"CreateService",
		strings.ToUpper("Post"),
		"/services",
		api.CreateService,
	},

	Route{
		"DeleteService",
		strings.ToUpper("Delete"),
		"/services/{id}",
		api.DeleteService,
	},

	Route{
		"GetServiceById",
		strings.ToUpper("Get"),
		"/services/{id}",
		api.GetServiceById,
	},

	Route{
		"ListServices",
		strings.ToUpper("Get"),
		"/services",
		api.ListServices,
	},

	Route{
		"UpdateService",
		strings.ToUpper("Put"),
		"/services/{id}",
		api.UpdateService,
	},

	Route{
		"CreateTeam",
		strings.ToUpper("Post"),
		"/teams",
		api.CreateTeam,
	},

	Route{
		"DeleteTeam",
		strings.ToUpper("Delete"),
		"/teams/{id}",
		api.DeleteTeam,
	},

	Route{
		"GetTeamById",
		strings.ToUpper("Get"),
		"/teams/{id}",
		api.GetTeamById,
	},

	Route{
		"ListTeams",
		strings.ToUpper("Get"),
		"/teams",
		api.ListTeams,
	},

	Route{
		"UpdateTeam",
		strings.ToUpper("Put"),
		"/teams/{id}",
		api.UpdateTeam,
	},

	Route{
		"CreateUser",
		strings.ToUpper("Post"),
		"/users",
		api.CreateUser,
	},

	Route{
		"DeleteUser",
		strings.ToUpper("Delete"),
		"/users/{id}",
		api.DeleteUser,
	},

	Route{
		"GetUserById",
		strings.ToUpper("Get"),
		"/users/{id}",
		api.GetUserById,
	},

	Route{
		"ListUsers",
		strings.ToUpper("Get"),
		"/users",
		api.ListUsers,
	},

	Route{
		"UpdateUser",
		strings.ToUpper("Put"),
		"/users/{id}",
		api.UpdateUser,
	},
}
