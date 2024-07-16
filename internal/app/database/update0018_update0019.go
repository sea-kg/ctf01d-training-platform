package database

import (
	"database/sql"
	"log/slog"
	"runtime"
)

func DatabaseUpdate_update0018_update0019(db *sql.DB, getInfo bool) (string, string, string, error) {

	// WARNING!!!
	// Do not change the update if it has already been installed by other developers or in production.
	// To correct the database, create a new update and register it in the list of updates.

	fromUpdateId, toUpdateId := ParseNameFuncUpdate(runtime.Caller(0))
	description := "Change type of score"
	if getInfo {
		return fromUpdateId, toUpdateId, description, nil
	}
	query := `alter table results alter column score type float;`
	_, err := db.Query(query)
	if err != nil {
		slog.Error("Problem with select, query: " + query + "\\n   error:" + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}
	return fromUpdateId, toUpdateId, description, nil
}
