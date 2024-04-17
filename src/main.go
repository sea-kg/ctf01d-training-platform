package main

import (
	"log"
	"net/http"

	sw "ctf01d/routers"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":4102", router))
}
