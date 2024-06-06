package database

import (
	"database/sql"
	"log/slog"
	"runtime"
)

func DatabaseUpdate_update0000_update0001(db *sql.DB, getInfo bool) (string, string, string, error) {

	// WARNING!!!
	// Do not change the update if it has already been installed by other developers or in production.
	// To correct the database, create a new update and register it in the list of updates.

	fromUpdateId, toUpdateId := ParseNameFuncUpdate(runtime.Caller(0))
	description := "Added table users"
	if getInfo {
		return fromUpdateId, toUpdateId, description, nil
	}
	query := `
		CREATE TABLE users (
		    id SERIAL PRIMARY KEY,
		    user_name VARCHAR(255) UNIQUE NOT NULL,
		    password_hash VARCHAR(255) NOT NULL,
		    avatar_url VARCHAR(255),
		    status VARCHAR(255),
		    role VARCHAR(255) NOT NULL CHECK (role IN ('admin', 'player', 'guest'))
		);
	`
	_, err := db.Exec(query)
	if err != nil {
		slog.Error("Problem with update, query: " + query + "\n   error:" + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}
	return fromUpdateId, toUpdateId, description, nil
}
