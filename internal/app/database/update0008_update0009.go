package database

import (
	"database/sql"
	"log/slog"
	"runtime"
)

func DatabaseUpdate_update0008_update0009(db *sql.DB, getInfo bool) (string, string, string, error) {

	// WARNING!!!
	// Do not change the update if it has already been installed by other developers or in production.
	// To correct the database, create a new update and register it in the list of updates.

	fromUpdateId, toUpdateId := ParseNameFuncUpdate(runtime.Caller(0))
	description := "Added table results"
	if getInfo {
		return fromUpdateId, toUpdateId, description, nil
	}
	query := `
		CREATE TABLE results (
		    id SERIAL PRIMARY KEY,
		    team_id INTEGER REFERENCES teams(id),
		    game_id INTEGER REFERENCES games(id),
		    score INTEGER NOT NULL,
		    rank INTEGER NOT NULL
		);
	`
	_, err := db.Exec(query)
	if err != nil {
		slog.Error("Problem with update, query: " + query + "\n   error:" + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}
	return fromUpdateId, toUpdateId, description, nil
}
