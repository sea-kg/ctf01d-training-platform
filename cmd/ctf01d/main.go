package main

import (
	"log/slog"
	"net/http"
	"os"

	"ctf01d/config"
	"ctf01d/internal/app/database"
	"ctf01d/internal/app/handlers"
	"ctf01d/internal/app/server"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("Config error: " + err.Error())
	}
	db, err := database.InitDatabase()
	if err != nil {
		slog.Error("Error opening database connection: " + err.Error())
		return
	}
	slog.Info("Database connection established successfully")
	defer db.Close()
	r := chi.NewRouter()
	h := &handlers.Handlers{
		DB: db,
	}
	s := handlers.NewServerInterfaceWrapper(h)

	r.Mount("/", server.HandlerFromMux(s, r))
	err = http.ListenAndServe(cfg.HTTP.Host+":"+cfg.HTTP.Port, r)
	if err != nil {
		slog.Error("Server error: " + err.Error())
	}
}
