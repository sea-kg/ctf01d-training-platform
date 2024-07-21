package migration

import (
	"database/sql"
	"log/slog"
	"runtime"
)

func DatabaseUpdate_update0010_update0010testdata(db *sql.DB, getInfo bool) (string, string, string, error) {
	// WARNING!!!
	// Do not change the update if it has already been installed by other developers or in production.
	// To correct the database, create a new update and register it in the list of updates.

	fromUpdateId, toUpdateId := ParseNameFuncUpdate(runtime.Caller(0))
	description := "Insert test data team_members"
	if getInfo {
		return fromUpdateId, toUpdateId, description, nil
	}
	query := `
		INSERT INTO team_members (user_id, team_id) VALUES
		((SELECT id FROM users WHERE user_name = 'Neo'), (SELECT id FROM teams WHERE name = 'HackersX')),
		((SELECT id FROM users WHERE user_name = 'Morpheus'), (SELECT id FROM teams WHERE name = 'HackersX')),
		((SELECT id FROM users WHERE user_name = 'Trinity'), (SELECT id FROM teams WHERE name = 'HackersX')),
		((SELECT id FROM users WHERE user_name = 'Cipher'), (SELECT id FROM teams WHERE name = 'CodeRed')),
		((SELECT id FROM users WHERE user_name = 'Seraph'), (SELECT id FROM teams WHERE name = 'CodeRed')),
		((SELECT id FROM users WHERE user_name = 'Smith'), (SELECT id FROM teams WHERE name = 'CodeRed')),
		((SELECT id FROM users WHERE user_name = 'Oracle'), (SELECT id FROM teams WHERE name = 'NullByte')),
		((SELECT id FROM users WHERE user_name = 'Sati'), (SELECT id FROM teams WHERE name = 'NullByte')),
		((SELECT id FROM users WHERE user_name = 'Dozer'), (SELECT id FROM teams WHERE name = 'NullByte'))
		;
	`
	_, err := db.Exec(query)
	if err != nil {
		slog.Error("Problem with update, query: " + query + "\n   error:" + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}
	return fromUpdateId, toUpdateId, description, nil
}
