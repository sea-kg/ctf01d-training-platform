package main

import (
	"log/slog"
	"net/http"
	"os"

	"ctf01d/config"
	"ctf01d/internal/app/handlers"
	migration "ctf01d/internal/app/migrations/psql"
	"ctf01d/internal/app/server"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func main() {
	path := "./config/config.development.yml"
	if envPath, exists := os.LookupEnv("CONFIG_PATH"); exists {
		path = envPath
	}

	cfg, err := config.NewConfig(path)
	if err != nil {
		slog.Error("Config error: " + err.Error())
		os.Exit(1)
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.Level(
			cfg.ParseLogLevel(cfg.Log.Level),
		),
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
	router := chi.NewRouter()
	hndlrs := &handlers.Handlers{
		DB: db,
	}
	svr := handlers.NewServerInterfaceWrapper(hndlrs)

	router.Mount("/api/", server.HandlerFromMux(svr, router))
	router.Mount("/", http.HandlerFunc(server.NewHtmlRouter))

	slog.Info("Server run on", slog.String("host", cfg.HTTP.Host), slog.String("port", cfg.HTTP.Port))
	err = http.ListenAndServe(cfg.HTTP.Host+":"+cfg.HTTP.Port, router)
	if err != nil {
		slog.Error("Server error: " + err.Error())
	}
}
