package database

import (
	"database/sql"
	"log/slog"
	"runtime"
)

func DatabaseUpdate_update0008_update0008testdata(db *sql.DB, getInfo bool) (string, string, string, error) {

	// WARNING!!!
	// Do not change the update if it has already been installed by other developers or in production.
	// To correct the database, create a new update and register it in the list of updates.

	fromUpdateId, toUpdateId := ParseNameFuncUpdate(runtime.Caller(0))
	description := "Insert test data games"
	if getInfo {
		return fromUpdateId, toUpdateId, description, nil
	}
	query := `
		INSERT INTO games (start_time, end_time, description) VALUES
		('2023-10-01 12:00:00', '2023-10-01 15:00:00', 'Capture the Network Flags'),
		('2023-10-02 12:00:00', '2023-10-02 15:00:00', 'Decrypt the Hidden Messages')
		;
	`
	_, err := db.Exec(query)
	if err != nil {
		slog.Error("Problem with update, query: " + query + "\n   error:" + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}
	return fromUpdateId, toUpdateId, description, nil
}
