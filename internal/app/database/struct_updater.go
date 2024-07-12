package database

import (
	"database/sql"
	"log/slog"
	"reflect"
	"runtime"
	"strings"
	"time"

	"ctf01d/config"
)

type DatabaseUpdateFunc func(db *sql.DB, getInfo bool) (string, string, string, error)

func RegisterAllUpdates() map[string][]DatabaseUpdateFunc {
	allUpdates := make(map[string][]DatabaseUpdateFunc)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0000_update0001)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0001_update0001admin)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0001_update0001testdata)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0001_update0002)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0002_update0003)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0003_update0004)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0004_update0005)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0005_update0006)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0006_update0006testdata)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0006_update0007)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0007_update0007testdata)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0007_update0008)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0008_update0008testdata)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0008_update0009)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0009_update0010)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0010_update0010testdata)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0010_update0011)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0011_update0011testdata)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0011_update0012)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0012_update0012testdata)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0012_update0013)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0013_update0013testdata)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0013_update0014)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0014_update0014fix1)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0014_update0015)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0015_update0015testdata)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0015_update0016)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0016_update0016testdata)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0016_update0017_no_university)
	allUpdates = RegisterDatabaseUpdate(allUpdates, DatabaseUpdate_update0017_update0018)
	return allUpdates
}

func RegisterDatabaseUpdate(allUpdates map[string][]DatabaseUpdateFunc, f DatabaseUpdateFunc) map[string][]DatabaseUpdateFunc {
	fromUpdateId, _, _, _ := f(nil, true)
	_, foundFrom := allUpdates[fromUpdateId]
	if !foundFrom {
		allUpdates[fromUpdateId] = []DatabaseUpdateFunc{}
	}
	allUpdates[fromUpdateId] = append(allUpdates[fromUpdateId], f)
	return allUpdates
}

func getTimestampNow() time.Time {
	return time.Now()
}

func insertUpdateInfo(db *sql.DB, fromUpdateId string, toUpdateId string, description string) error {
	query := "" +
		"INSERT INTO database_updates (" +
		"    updated_at, from_update_id, to_update_id, description" +
		") VALUES(" +
		"    $1, $2, $3, $4" +
		")"
	_, err := db.Exec(query, getTimestampNow(), fromUpdateId, toUpdateId, description)
	if err != nil {
		slog.Error("insertUpdateInfo: " + err.Error())
		return err
	}
	return nil
}

func getInstalledDatabaseVersions(db *sql.DB) ([]string, error) {
	query := `SELECT to_update_id FROM database_updates;`
	var installedUpdates []string
	rows, err := db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Warn("database_updates - No rows")
			err = insertUpdateInfo(db, "", "update0000", "Added table database_updates")
			if err != nil {
				slog.Error("Problem with insert into database_updates (1) " + err.Error())
				return installedUpdates, err
			}
			installedUpdates = append(installedUpdates, "update0000")
			return installedUpdates, nil
		} else if strings.Contains(err.Error(), "relation \"database_updates\" does not exist") {
			query = "" +
				"CREATE TABLE database_updates (" +
				"    updated_at TIMESTAMP NOT NULL," +
				"    from_update_id VARCHAR(32) NOT NULL," +
				"    to_update_id VARCHAR(32) UNIQUE NOT NULL," +
				"    description TEXT, " +
				"    PRIMARY KEY (from_update_id, to_update_id)" +
				");"
			_, err := db.Exec(query)
			if err != nil {
				slog.Error("Create table database_updates: " + err.Error())
				return installedUpdates, err
			}
			err = insertUpdateInfo(db, "", "update0000", "Added table database_updates")
			if err != nil {
				slog.Error("Problem with insert into database_updates (2) " + err.Error())
				return installedUpdates, err
			}
			installedUpdates = append(installedUpdates, "update0000")
			return installedUpdates, nil
		} else {
			return installedUpdates, err
		}
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var update string
		if err := rows.Scan(&update); err != nil {
			return installedUpdates, err
		}
		installedUpdates = append(installedUpdates, update)
	}
	if err = rows.Err(); err != nil {
		return installedUpdates, err
	}
	return installedUpdates, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func InitDatabase(cfg *config.Config) (*sql.DB, error) {
	slog.Info("Initing database...")
	db, err := sql.Open(cfg.DB.Driver, cfg.DB.DataSource)
	if err != nil {
		slog.Error("Error connecting to the database: " + err.Error())
		defer db.Close()
		return nil, err
	}

	if err := db.Ping(); err != nil {
		slog.Error("Error pinging database: " + err.Error())
		return nil, err
	}
	slog.Info("Database connection established successfully")

	// last_update := &models.DatabaseUpdate{}
	installedUpdates, err := getInstalledDatabaseVersions(db)
	if err != nil {
		slog.Error("Problem with database (InitDatabase): " + err.Error())
		defer db.Close()
		return nil, err
	}
	slog.Info("Found last update: " + installedUpdates[len(installedUpdates)-1])

	allUpdates := RegisterAllUpdates()

	var alreadyCheckedUpdates []string
	installedSomeUpdate := true

	for installedSomeUpdate { // while
		installedSomeUpdate = false
		installedUpdates, _ := getInstalledDatabaseVersions(db)
		for _, installed_update_id := range installedUpdates {
			slog.Debug("installed_update: " + installed_update_id)
			updates, found_update := allUpdates[installed_update_id]
			if found_update {
				for _, update_func := range updates {
					fromUpdateId, toUpdateId, description, _ := update_func(db, true)
					if contains(alreadyCheckedUpdates, toUpdateId) {
						continue
					} else if contains(installedUpdates, toUpdateId) && !contains(alreadyCheckedUpdates, toUpdateId) {
						slog.Info("Update already installed " + fromUpdateId + " -> " + toUpdateId)
						alreadyCheckedUpdates = append(alreadyCheckedUpdates, toUpdateId)
					} else {
						_, _, _, err := update_func(db, false)
						if err != nil {
							slog.Error("Problem with update (1) " + fromUpdateId + " -> " + toUpdateId + " (" + description + "): " + err.Error())
							defer db.Close()
							return db, err
						}
						err = insertUpdateInfo(db, fromUpdateId, toUpdateId, description)
						if err != nil {
							slog.Error("Problem with update (2) " + fromUpdateId + " -> " + toUpdateId + " (" + description + "): " + err.Error())
							defer db.Close()
							return db, err
						}
						slog.Info("Successfully applied update " + fromUpdateId + " -> " + toUpdateId + " (" + description + ")")
						alreadyCheckedUpdates = append(alreadyCheckedUpdates, toUpdateId)
						installedSomeUpdate = true // try find next updates
					}
				}
			}
		}
	}
	// successfully
	slog.Info("Database inited.")
	return db, nil
}

func GetFunctionName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

func ParseNameFuncUpdate(pc uintptr, file string, line int, ok bool) (string, string) {
	functionName := runtime.FuncForPC(pc).Name()

	// remove prefix ctf01d/internal/app/database.DatabaseUpdate_update0000_update0001, leave only the name of func: DatabaseUpdate_update0000_update0001
	functionName = strings.Split(functionName, "database.")[1]

	res1 := strings.Split(functionName, "_")
	if res1[0] != "DatabaseUpdate" {
		slog.Error("Wrong func name is not a database update func: " + functionName + ", expected like DatabaseUpdate_xxxx0000_xxxxx0001")
		return "", ""
	}
	if res1[1] == res1[2] {
		slog.Error("Could not be same update from and to: " + res1[2] + ", expected like DatabaseUpdate_xxxx0000_xxxxx0001")
		return "", ""
	}
	return res1[1], res1[2] // from / to
}
