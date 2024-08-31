package migration

import (
	"database/sql"
	"log/slog"
	"runtime"
)

func DatabaseUpdate_update0002_update0003(db *sql.DB, getInfo bool) (string, string, string, error) {
	// WARNING!!!
	// Do not change the update if it has already been installed by other developers or in production.
	// To correct the database, create a new update and register it in the list of updates.

	fromUpdateId, toUpdateId := ParseNameFuncUpdate(runtime.Caller(0))
	description := "Added table sessions"
	if getInfo {
		return fromUpdateId, toUpdateId, description, nil
	}
	query := `
		CREATE TABLE sessions (
		    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
		    user_id INTEGER UNIQUE REFERENCES users(id),
		    expires_at TIMESTAMP NOT NULL
		);
	`
	_, err := db.Exec(query)
	if err != nil {
		slog.Error("Problem with update, query: " + query + "\n   error:" + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}
	return fromUpdateId, toUpdateId, description, nil
}
