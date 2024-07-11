package database

import (
	"database/sql"
	"log/slog"
	"runtime"
)

func DatabaseUpdate_update0011_update0011testdata(db *sql.DB, getInfo bool) (string, string, string, error) {
	// WARNING!!!
	// Do not change the update if it has already been installed by other developers or in production.
	// To correct the database, create a new update and register it in the list of updates.

	fromUpdateId, toUpdateId := ParseNameFuncUpdate(runtime.Caller(0))
	description := "Insert testdata team_games"
	if getInfo {
		return fromUpdateId, toUpdateId, description, nil
	}
	query := `
		INSERT INTO team_games (team_id, game_id) VALUES
		((SELECT id FROM teams WHERE name = 'HackersX'), 1),
		((SELECT id FROM teams WHERE name = 'CodeRed'), 1),
		((SELECT id FROM teams WHERE name = 'NullByte'), 2),
		((SELECT id FROM teams WHERE name = 'HackersX'), 2)
		;
	`
	_, err := db.Exec(query)
	if err != nil {
		slog.Error("Problem with update, query: " + query + "\n   error:" + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}
	return fromUpdateId, toUpdateId, description, nil
}
