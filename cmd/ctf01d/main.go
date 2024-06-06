package main

import (
	config "ctf01d/configs"
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
	db, err := sql.Open(cfg.DB.Driver, cfg.DB.DataSource)
	if err != nil {
		slog.Error("Error connecting to the database: " + err.Error())
	}
	defer db.Close()
	router := routers.NewRouter(db)
	slog.Info("Server started on " + cfg.HTTP.Host+":"+cfg.HTTP.Port)

	err = http.ListenAndServe(cfg.HTTP.Host+":"+cfg.HTTP.Port, router)
	if err != nil {
		slog.Error("Server error: " + err.Error())
	}
}
