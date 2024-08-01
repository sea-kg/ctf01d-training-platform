package migration

import (
	"database/sql"
	"log/slog"
	"runtime"
)

func DatabaseUpdate_update0019_update0019testdata(db *sql.DB, getInfo bool) (string, string, string, error) {
	// WARNING!!!
	// Do not change the update if it has already been installed by other developers or in production.
	// To correct the database, create a new update and register it in the list of updates.

	fromUpdateId, toUpdateId := ParseNameFuncUpdate(runtime.Caller(0))
	description := "clear robo-hash url"
	if getInfo {
		return fromUpdateId, toUpdateId, description, nil
	}
	query := `
		BEGIN;
		update users set avatar_url = null;
		update teams set avatar_url = null;
		update services set logo_url = null;
		COMMIT;
	`
	_, err := db.Query(query)
	if err != nil {
		slog.Error("Problem with select, query: " + query + "\\n   error:" + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}
	return fromUpdateId, toUpdateId, description, nil
}
