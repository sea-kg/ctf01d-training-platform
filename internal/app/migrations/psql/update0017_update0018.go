package migration

import (
	"database/sql"
	"log/slog"
	"runtime"
)

func DatabaseUpdate_update0017_update0018(db *sql.DB, getInfo bool) (string, string, string, error) {
	// WARNING!!!
	// Do not change the update if it has already been installed by other developers or in production.
	// To correct the database, create a new update and register it in the list of updates.

	fromUpdateId, toUpdateId := ParseNameFuncUpdate(runtime.Caller(0))
	description := "Add role to profile and team history"
	if getInfo {
		return fromUpdateId, toUpdateId, description, nil
	}
	query := `
		BEGIN;
		ALTER TABLE profiles ADD COLUMN role varchar(50) default 'player' NOT NULL;
		ALTER TABLE team_history ADD COLUMN role varchar(50) default 'player' NOT NULL;
		COMMIT;
	`
	_, err := db.Query(query)
	if err != nil {
		slog.Error("Problem with select, query: " + query + "\n   error:" + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}
	return fromUpdateId, toUpdateId, description, nil
}
