package database

import (
	"database/sql"
	"log/slog"
	"runtime"
)

func DatabaseUpdate_update0016_update0017_no_university(db *sql.DB, getInfo bool) (string, string, string, error) {

	// WARNING!!!
	// Do not change the update if it has already been installed by other developers or in production.
	// To correct the database, create a new update and register it in the list of updates.

	fromUpdateId, toUpdateId := ParseNameFuncUpdate(runtime.Caller(0))
	description := "Added row to universities (no university)"
	if getInfo {
		return fromUpdateId, toUpdateId, description, nil
	}
	query := `
		INSERT INTO universities (id, name) VALUES
		('00000000-0000-0000-0000-000000000000', 'Без университета')
		;
	`
	_, err := db.Query(query)
	if err != nil {
		slog.Error("Problem with select, query: " + query + "\n   error:" + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}
	return fromUpdateId, toUpdateId, description, nil
}
