package main

import (
	config "ctf01d/configs"
	"ctf01d/internal/app/routers"
	models "ctf01d/internal/app/db"
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Error("Starting server...")

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("Config error: " + err.Error())
	}
	db, err := sql.Open(cfg.DB.Driver, cfg.DB.DataSource)
	if err != nil {
		slog.Error("Error connecting to the database: " + err.Error())
	} else {
		query := `SELECT update_id FROM database_updates WHERE id=(SELECT max(id) FROM database_updates)`
		last_update := &models.DatabaseUpdate{}
		err := db.QueryRow(query).Scan(&last_update.Id, &last_update.StartTime, &last_update.UpdateId, &last_update.Description)
		if err != nil {
			slog.Error("Problem with database: " + err.Error())
			return
		}
	}
	defer db.Close()
	router := routers.NewRouter(db)
	slog.Info("Server started on http://" + cfg.HTTP.Host+":"+cfg.HTTP.Port)

	err = http.ListenAndServe(cfg.HTTP.Host+":"+cfg.HTTP.Port, router)
	if err != nil {
		slog.Error("Server error: " + err.Error())
	}
}
