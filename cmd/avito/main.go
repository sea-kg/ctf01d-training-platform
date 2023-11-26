package main

import (
	"avitoStart/internal/app"
	"errors"
	"log"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Println(errors.New("Error New in main"))
	}
	err = a.Run()
	if err != nil {
		log.Println(errors.New("Error Start in main"))
	}

}
