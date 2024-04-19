package main

import (
	config "ctf01d/configs"
	"ctf01d/internal/app/routers"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	db, err := sql.Open(cfg.DB.Driver, cfg.DB.DataSource)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()
	log.Printf("Server started")
	router := routers.NewRouter(db)

	log.Fatal(
		http.ListenAndServe(cfg.HTTP.Host+":"+cfg.HTTP.Port, router),
	)
}
