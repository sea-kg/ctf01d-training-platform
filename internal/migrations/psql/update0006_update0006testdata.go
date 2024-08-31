package migration

import (
	"database/sql"
	"log/slog"
	"runtime"
)

func DatabaseUpdate_update0006_update0006testdata(db *sql.DB, getInfo bool) (string, string, string, error) {
	// WARNING!!!
	// Do not change the update if it has already been installed by other developers or in production.
	// To correct the database, create a new update and register it in the list of updates.

	fromUpdateId, toUpdateId := ParseNameFuncUpdate(runtime.Caller(0))
	description := "Insert test data teams"
	if getInfo {
		return fromUpdateId, toUpdateId, description, nil
	}
	query := `
		INSERT INTO teams (name, description, university_id, social_links, avatar_url) VALUES
		('HackersX', 'Specialized in network attacks and defense', 401, '', 'https://robohash.org/hackersx'),
		('CodeRed', 'Expert in cryptography and steganography', 402, '', 'https://robohash.org/codered'),
		('NullByte', 'Skilled in web security and binary exploitation', 403, '', 'https://robohash.org/nullbyte')
		;
	`
	_, err := db.Exec(query)
	if err != nil {
		slog.Error("Problem with update, query: " + query + "\n   error:" + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}
	return fromUpdateId, toUpdateId, description, nil
}
