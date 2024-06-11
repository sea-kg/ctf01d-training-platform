package database

import (
	"database/sql"
	"log/slog"
	"runtime"
)

func DatabaseUpdate_update0007_update0007testdata(db *sql.DB, getInfo bool) (string, string, string, error) {

	// WARNING!!!
	// Do not change the update if it has already been installed by other developers or in production.
	// To correct the database, create a new update and register it in the list of updates.

	fromUpdateId, toUpdateId := ParseNameFuncUpdate(runtime.Caller(0))
	description := "Added table services"
	if getInfo {
		return fromUpdateId, toUpdateId, description, nil
	}
	query := `
		INSERT INTO services (name, author, logo_url, description, is_public) VALUES
		('NetAttack', 'Phantom', '', 'Simulated network attack platform', TRUE),
		('CryptoBox', 'Enigma', '', 'Cryptography challenge service', TRUE)
		;
	`
	_, err := db.Exec(query)
	if err != nil {
		slog.Error("Problem with update, query: " + query + "\n   error:" + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}
	return fromUpdateId, toUpdateId, description, nil
}
