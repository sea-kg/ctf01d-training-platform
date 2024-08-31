package migration

import (
	"database/sql"
	"log/slog"
	"runtime"
)

func DatabaseUpdate_update0015_update0016(db *sql.DB, getInfo bool) (string, string, string, error) {
	// WARNING!!!
	// Do not change the update if it has already been installed by other developers or in production.
	// To correct the database, create a new update and register it in the list of updates.

	fromUpdateId, toUpdateId := ParseNameFuncUpdate(runtime.Caller(0))
	description := "Make profile table"
	if getInfo {
		return fromUpdateId, toUpdateId, description, nil
	}
	query := `
		BEGIN;
		ALTER TABLE team_members RENAME TO profiles;
		ALTER TABLE profiles RENAME COLUMN team_id TO current_team_id;
		ALTER TABLE profiles ADD COLUMN created_at TIMESTAMP default now();
		ALTER TABLE profiles ADD COLUMN updated_at TIMESTAMP default now();
		ALTER TABLE profiles ADD COLUMN id UUID DEFAULT gen_random_uuid();

		CREATE TABLE team_history (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			user_id UUID REFERENCES users(id),
			team_id UUID REFERENCES teams(id),
			joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			left_at TIMESTAMP
		);
		COMMIT;
	`
	_, err := db.Query(query)
	if err != nil {
		slog.Error("Problem with select, query: " + query + "\n   error:" + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}
	return fromUpdateId, toUpdateId, description, nil
}
