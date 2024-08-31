package migration

import (
	"database/sql"
	"log/slog"
	"runtime"
)

func DatabaseUpdate_update0001_update0001admin(db *sql.DB, getInfo bool) (string, string, string, error) {
	// WARNING!!!
	// Do not change the update if it has already been installed by other developers or in production.
	// To correct the database, create a new update and register it in the list of updates.

	fromUpdateId, toUpdateId := ParseNameFuncUpdate(runtime.Caller(0))
	description := "Insert admin/admin/admin user"
	if getInfo {
		return fromUpdateId, toUpdateId, description, nil
	}
	// admin / admin
	query := `
		INSERT INTO users (user_name, password_hash, role, avatar_url, status) VALUES
		('admin', '$2a$10$zHYCbFK9BkGRL6oxy91GiOgrqTs/0A.K4FUpA/d9h..aJiVyePxqi', 'admin', 'https://robohash.org/admin', 'active')
		;
	`
	_, err := db.Exec(query)
	if err != nil {
		slog.Error("Problem with update, query: " + query + "\n   error:" + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}
	return fromUpdateId, toUpdateId, description, nil
}
