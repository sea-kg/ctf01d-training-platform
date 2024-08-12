package migration

import (
	"database/sql"
	"log/slog"
	"runtime"
)

func DatabaseUpdate_update0021_update0022(db *sql.DB, getInfo bool) (string, string, string, error) {
	// WARNING!!!
	// Do not change the update if it has already been installed by other developers or in production.
	// To correct the database, create a new update and register it in the list of updates.

	fromUpdateId, toUpdateId := ParseNameFuncUpdate(runtime.Caller(0))
	description := "Add cascade delete for results when deleting games"
	if getInfo {
		return fromUpdateId, toUpdateId, description, nil
	}

	query := `
		BEGIN;
		ALTER TABLE results DROP CONSTRAINT IF EXISTS results_game_id_fkey;
		ALTER TABLE results ADD CONSTRAINT results_game_id_fkey
		    FOREIGN KEY (game_id) REFERENCES games(id) ON DELETE CASCADE;
		COMMIT;
	`
	_, err := db.Exec(query)
	if err != nil {
		slog.Error("Problem with select, query: " + query + "\\n   error:" + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}
	return fromUpdateId, toUpdateId, description, nil
}
