package main

import (
	"ctf01d/config"
	"ctf01d/internal/app/routers"
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("Config error: " + err.Error())
	}

	slog.Info("Connecting to the database ...")
	db, err := sql.Open(cfg.DB.Driver, cfg.DB.DataSource)
	if err != nil {
		slog.Error("Error opening database connection: " + err.Error())
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		slog.Error("Error pinging database: " + err.Error())
		return
	}
	slog.Info("Database connection established successfully")

	router := routers.NewRouter(db)
	slog.Info("Server started on http://" + cfg.HTTP.Host + ":" + cfg.HTTP.Port)

	err = http.ListenAndServe(cfg.HTTP.Host+":"+cfg.HTTP.Port, router)
	if err != nil {
		slog.Error("Server error: " + err.Error())
	}
}
