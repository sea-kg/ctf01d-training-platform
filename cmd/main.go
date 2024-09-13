package main

import (
	"log/slog"
	"net/http"
	"os"

	"ctf01d/internal/config"
	"ctf01d/internal/handler"
	"ctf01d/internal/httpserver"
	"ctf01d/internal/middleware"
	migration "ctf01d/internal/migrations/psql"
	"ctf01d/internal/repository"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	_ "go.uber.org/automaxprocs"
)

func main() {
	path := "./configs/config.development.yml"
	if envPath, exists := os.LookupEnv("CONFIG_PATH"); exists {
		path = envPath
	}

	cfg, err := config.New(path)
	if err != nil {
		slog.Error("Config error: " + err.Error())
		os.Exit(1)
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.Level(cfg.ParseLogLevel(cfg.Log.Level)),
	}))
	slog.SetDefault(logger)
	slog.Info("Config path is - " + path)
	db, err := migration.InitDatabase(cfg)
	if err != nil {
		slog.Error("Error opening database connection: " + err.Error())
		return
	}
	slog.Info("Database connection established successfully")
	defer db.Close()

	if err := db.Ping(); err != nil {
		slog.Error("Database is not reachable: %v", err)
	}

	swagger, err := openapi3.NewLoader().LoadFromFile("api/openapi.yaml")
	if err != nil {
		slog.Error("Failed to load OpenAPI spec: %v", err)
	}

	validatorMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			route, pathParams, err := swagger.Paths.Find(r.Method, r.URL.Path)
			if err != nil || route == nil {
				http.Error(w, "Route not found", http.StatusNotFound)
				return
			}

			if security := route.GetOperation().Security; security != nil && len(*security) > 0 {
				err := middleware.Auth(db)(next).ServeHTTP(w, r)
				if err != nil {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}

	router := chi.NewRouter()

	// Добавление CORS middleware
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	router.Use(corsMiddleware.Handler)

	// Auth Middleware
	sessionRepo := repository.NewSessionRepository(db)
	router.Use(middleware.Auth(sessionRepo))

	hndlr := &handler.Handler{
		DB: db,
	}
	svr := handler.NewServerInterfaceWrapper(hndlr)

	router.Mount("/api/", httpserver.HandlerFromMux(svr, router))
	router.Mount("/", http.HandlerFunc(httpserver.NewHtmlRouter))

	slog.Info("Server run on", slog.String("host", cfg.HTTP.Host), slog.String("port", cfg.HTTP.Port))
	err = http.ListenAndServe(cfg.HTTP.Host+":"+cfg.HTTP.Port, router)
	if err != nil {
		slog.Error("Server error: " + err.Error())
	}
}
