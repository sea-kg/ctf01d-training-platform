package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"ctf01d/internal/config"
	"ctf01d/internal/handler"
	"ctf01d/internal/httpserver"
	migration "ctf01d/internal/migrations/psql"

	"ctf01d/internal/middleware/auth"

	"ctf01d/pkg/ginmiddleware"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

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
	router := gin.New()

	// Добавление CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	hndlr := &handler.Handler{
		DB: db,
	}
	svr := handler.NewServerInterfaceWrapper(hndlr)

	// router.Mount("/api/", httpserver.HandlerFromMux(svr, router))
	// router.Mount("/", http.HandlerFunc(httpserver.NewHtmlRouter))
	router.Any("/api/", gin.WrapH(httpserver.Handler(svr, http.DefaultServeMux)))
	router.GET("/", gin.WrapF(httpserver.NewHtmlRouter))

	slog.Info("Server run on", slog.String("host", cfg.HTTP.Host), slog.String("port", cfg.HTTP.Port))
	err = http.ListenAndServe(cfg.HTTP.Host+":"+cfg.HTTP.Port, router)
	if err != nil {
		slog.Error("Server error: " + err.Error())
	}
	swagger, err := openapi3.NewLoader().LoadFromFile("api/openapi.yaml")
	if err != nil {
		log.Fatalf("Failed to load OpenAPI spec: %v", err)
	}
	middlewareOptions := &ginmiddleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: func(_ context.Context, _ *openapi3filter.AuthenticationInput) error {
				return nil
			},
		},
	}
	requestValidator := ginmiddleware.OapiRequestValidatorWithOptions(swagger, middlewareOptions)
	router.Use(requestValidator)
	responseValidator := ginmiddleware.OapiResponseValidatorWithOptions(swagger, middlewareOptions)
	router.Use(responseValidator)
	options := httpserver.GinServerOptions{
		Middlewares: []httpserver.MiddlewareFunc{
			auth.AuthenticationMiddleware(db),
		},
	}
	httpserver.RegisterHandlersWithOptions(r, si, options)
	log.Fatal(router.Run(cfg.HTTP.Host + ":" + cfg.HTTP.Port))
}
