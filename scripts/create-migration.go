package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func main() {
	dir := "./internal/migrations/psql/"
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var maxUpdate int
	re := regexp.MustCompile(`update(\d{4})_update(\d{4}).go`)

	for _, file := range files {
		matches := re.FindStringSubmatch(file.Name())
		if len(matches) == 3 {
			toUpdate, err := strconv.Atoi(matches[2])
			if err != nil {
				log.Fatal(err)
			}
			if toUpdate > maxUpdate {
				maxUpdate = toUpdate
			}
		}
	}

	nextUpdate := maxUpdate + 1
	newFileName := fmt.Sprintf("update%04d_update%04d.go", maxUpdate, nextUpdate)
	newFilePath := filepath.Join(dir, newFileName)

	template := fmt.Sprintf(`package migration

import (
	"database/sql"
	"log/slog"
	"runtime"
)

func DatabaseUpdate_update%04d_update%04d(db *sql.DB, getInfo bool) (string, string, string, error) {
	// WARNING!!!
	// Do not change the update if it has already been installed by other developers or in production.
	// To correct the database, create a new update and register it in the list of updates.

	fromUpdateId, toUpdateId := ParseNameFuncUpdate(runtime.Caller(0))
	description := ""
	if getInfo {
		return fromUpdateId, toUpdateId, description, nil
	}
	query := ""
	_, err := db.Query(query)
	if err != nil {
		slog.Error("Problem with select, query: " + query + "\\n   error:" + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}
	return fromUpdateId, toUpdateId, description, nil
}
`, maxUpdate, nextUpdate)

	err = os.WriteFile(newFilePath, []byte(template), 0o644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created new migration file: %s\n", newFilePath)
}
