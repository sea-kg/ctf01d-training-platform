package main

import (
	"log"
	"net/http"

	"ctf01d/config"
	"ctf01d/routers"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	log.Printf("Server started")
	router := routers.NewRouter()

	log.Fatal(
		http.ListenAndServe(cfg.HTTP.Host+":"+cfg.HTTP.Port, router),
	)
}
