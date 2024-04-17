// DEPRECATED
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func getConnString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
		os.Getenv("SERVICE2_GO_MYSQL_USER"),
		os.Getenv("SERVICE2_GO_MYSQL_PASSWORD"),
		os.Getenv("SERVICE2_GO_MYSQL_HOST"),
		os.Getenv("SERVICE2_GO_MYSQL_DBNAME"))
}

type FlagDB struct {
	ID         int    `db:"id"`
	ExternalID string `db:"external_id"`
	Flag       string `db:"flag"`
	ServiceID  *int   `db:"service_id"`
	GameID     *int   `db:"game_id"`
}

type FlagResponse struct {
	ExternalID string `json:"flag_id"`
	Flag       string `json:"flag"`
	ServiceID  int    `json:"serviceId,omitempty"`
	GameID     int    `json:"gameId,omitempty"`
}

type ErrorJson struct {
	Error string
}

// --------------------------------------------------------

func errJson(w http.ResponseWriter, msg string) {
	errorJson := ErrorJson{msg}
	js, err := json.Marshal(errorJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)
	w.Write(js)
}

// --------------------------------------------------------

func putFlag(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	externalID := params["flag_id"]
	flagValue := params["flag"]

	db, err := sqlx.Connect("mysql", getConnString())
	if err != nil {
		log.Printf("FAILED connect to database %s\n", err.Error())
		errJson(w, err.Error())
		return
	}

	_, err2 := db.Exec(`INSERT INTO flags (external_id, flag) VALUES(?,?)`, externalID, flagValue)
	if err2 != nil {
		log.Printf("FAILED insert flag=%s with flag_id=%s\n", externalID, flagValue)
		errJson(w, err2.Error())
		db.Close()
		return
	}
	flag2 := FlagResponse{externalID, flagValue, 1, 1}
	js, err := json.Marshal(flag2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		db.Close()
		return
	}

	log.Printf("Inserted flag=%s with flag_id=%s\n", externalID, flagValue)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	db.Close()
}

// --------------------------------------------------------

func getFlag(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	externalID := params["flag_id"]

	db, err := sqlx.Connect("mysql", getConnString())
	if err != nil {
		log.Printf("FAILED connect to database %s\n", err.Error())
		errJson(w, err.Error())
		return
	}

	var flag FlagDB
	err2 := db.Get(&flag, `SELECT * FROM flags WHERE external_id = ? LIMIT 1`, externalID)
	if err2 != nil {
		db.Close()
		errJson(w, err2.Error())
		return
	}

	js, err := json.Marshal(FlagResponse{ExternalID: flag.ExternalID, Flag: flag.Flag})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Get flag=%s with flag_id=%s", flag.ExternalID, flag.Flag)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	db.Close()
}

// --------------------------------------------------------

type FlagIDs struct {
	FlagIDs []string
}

func flagsList(w http.ResponseWriter, r *http.Request) {
	db, err := sqlx.Connect("mysql", getConnString())
	if err != nil {
		log.Printf("FAILED connect to database %s\n", err.Error())
		errJson(w, err.Error())
		return
	}

	var flag_ids []string
	err2 := db.Select(&flag_ids, "SELECT external_id FROM flags")
	if err2 != nil {
		errJson(w, err2.Error())
		db.Close()
		return
	}
	flags2 := FlagIDs{flag_ids}
	js, err := json.Marshal(flags2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		db.Close()
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	db.Close()
}

// --------------------------------------------------------

var templates = template.Must(template.ParseFiles("html/index.html"))

func index(w http.ResponseWriter, r *http.Request) {
	templates.Execute(w, nil)
	// if err != nil {
	// 	return err
	// }
}

// --------------------------------------------------------

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("Request %s processed. Execution time: %s\n", r.RequestURI, time.Since(startTime))
	})
}

// --------------------------------------------------------

func main() {
	host := "localhost"
	port := "4102"

	rtr := mux.NewRouter()
	rtr.Use(loggingMiddleware)
	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/api/flags/{flag_id:[a-zA-Z0-9]+}", getFlag).Methods("GET")
	rtr.HandleFunc("/api/flags/{flag_id:[a-zA-Z0-9]+}/{flag:[a-zA-Z0-9-]+}", putFlag).Methods("POST")
	rtr.HandleFunc("/api/flags/", flagsList)
	http.Handle("/", rtr)

	fs := http.FileServer(http.Dir("html/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Printf("Start server on %s:%s\n", host, port)
	http.ListenAndServe(host+":"+port, nil)
}
