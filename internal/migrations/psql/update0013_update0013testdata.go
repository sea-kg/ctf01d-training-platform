package migration

import (
	"database/sql"
	"log"
	"log/slog"
	"runtime"

	"github.com/jaswdr/faker/v2"
)

func DatabaseUpdate_update0013_update0013testdata(db *sql.DB, getInfo bool) (string, string, string, error) {
	// WARNING!!!
	// Do not change the update if it has already been installed by other developers or in production.
	// To correct the database, create a new update and register it in the list of updates.

	fromUpdateId, toUpdateId := ParseNameFuncUpdate(runtime.Caller(0))
	description := "Insert test data to game_services"
	if getInfo {
		return fromUpdateId, toUpdateId, description, nil
	}
	query := `
		SELECT id from users
	`
	rows, err := db.Query(query)
	if err != nil {
		slog.Error("Problem with select, query: " + query + "\n   error:" + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}
	fake := faker.New()
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}
		name := fake.Person().Name()
		query := "UPDATE users SET display_name = $1 WHERE id = $2;"
		_, err := db.Exec(query, name, id)
		if err != nil {
			slog.Error("Problem with update, query: " + query + "\n   error:" + err.Error())
			return fromUpdateId, toUpdateId, description, err
		}
	}
	return fromUpdateId, toUpdateId, description, nil
}
