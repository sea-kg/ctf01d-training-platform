package migration

import (
	"database/sql"
	"log/slog"
	"runtime"
)

func DatabaseUpdate_update0019_update0020(db *sql.DB, getInfo bool) (string, string, string, error) {
	// WARNING!!!
	// Do not change the update if it has already been installed by other developers or in production.
	// To correct the database, create a new update and register it in the list of updates.

	fromUpdateId, toUpdateId := ParseNameFuncUpdate(runtime.Caller(0))
	description := "add table team_member_requests"
	if getInfo {
		return fromUpdateId, toUpdateId, description, nil
	}
	query := `
		CREATE TABLE team_member_requests (
			id SERIAL PRIMARY KEY,
			team_id UUID NOT NULL,
			user_id UUID NOT NULL,
			role VARCHAR(50) NOT NULL, -- роль в команде (player, captain)
			status VARCHAR(50) NOT NULL DEFAULT 'pending',  -- статус запроса (pending, approved, rejected)
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err := db.Query(query)
	if err != nil {
		slog.Error("Problem with select, query: " + query + "\\n   error:" + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}
	return fromUpdateId, toUpdateId, description, nil
}
