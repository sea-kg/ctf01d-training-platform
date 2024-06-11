package database

import (
	"database/sql"
	"log/slog"
	"runtime"
)

func DatabaseUpdate_update0001_update0001testdata(db *sql.DB, getInfo bool) (string, string, string, error) {

	// WARNING!!!
	// Do not change the update if it has already been installed by other developers or in production.
	// To correct the database, create a new update and register it in the list of updates.

	fromUpdateId, toUpdateId := ParseNameFuncUpdate(runtime.Caller(0))
	description := "Insert test data users"
	if getInfo {
		return fromUpdateId, toUpdateId, description, nil
	}
	query := `
		INSERT INTO users (user_name, password_hash, role, avatar_url, status) VALUES
		('Neo', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/neo', 'active'),
		('Morpheus', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/morpheus', 'active'),
		('Trinity', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/trinity', 'active'),
		('Cipher', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/cipher', 'active'),
		('Seraph', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/seraph', 'active'),
		('Smith', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/smith', 'inactive'),
		('Oracle', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/oracle', 'active'),
		('Sati', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/sati', 'active'),
		('Apoc', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/apoc', 'active'),
		('Dozer', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'admin', 'https://robohash.org/dozer', '')
		;
	`
	_, err := db.Exec(query)
	if err != nil {
		slog.Error("Problem with update, query: " + query + "\n   error:" + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}
	return fromUpdateId, toUpdateId, description, nil
}
