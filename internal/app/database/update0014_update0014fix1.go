package database

import (
	"database/sql"
	"log/slog"
	"runtime"
)

func DatabaseUpdate_update0014_update0014fix1(db *sql.DB, getInfo bool) (string, string, string, error) {

	// WARNING!!!
	// Do not change the update if it has already been installed by other developers or in production.
	// To correct the database, create a new update and register it in the list of updates.

	fromUpdateId, toUpdateId := ParseNameFuncUpdate(runtime.Caller(0))
	description := "Rename new_id to id for table sessions"
	if getInfo {
		return fromUpdateId, toUpdateId, description, nil
	}
	query := `
		BEGIN;

		ALTER TABLE sessions DROP COLUMN id;
		ALTER TABLE sessions RENAME COLUMN new_id TO id;

		ALTER TABLE sessions ADD CONSTRAINT sessions_id_unique UNIQUE (id);

		COMMIT;
	`
	_, err := db.Exec(query)
	if err != nil {
		slog.Error("Problem with update, query: " + query + "\n   error:" + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}

	return fromUpdateId, toUpdateId, description, nil
}
