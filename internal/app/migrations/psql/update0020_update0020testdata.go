package migration

import (
	"database/sql"
	"log/slog"
	"math/rand"
	"runtime"

	"github.com/google/uuid"
)

func DatabaseUpdate_update0020_update0020testdata(db *sql.DB, getInfo bool) (string, string, string, error) {
	// WARNING!!!
	// Do not change the update if it has already been installed by other developers or in production.
	// To correct the database, create a new update and register it in the list of updates.

	fromUpdateId, toUpdateId := ParseNameFuncUpdate(runtime.Caller(0))
	description := "Add test data with new teams, games, and results"
	if getInfo {
		return fromUpdateId, toUpdateId, description, nil
	}

	tx, err := db.Begin()
	if err != nil {
		slog.Error("Failed to begin transaction: " + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}

	// Очистка таблицы team_games
	_, err = tx.Exec(`DELETE FROM team_games`)
	if err != nil {
		tx.Rollback()
		slog.Error("Failed to clear team_games table: " + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}

	query := `
		-- Добавление новых команд
		INSERT INTO teams (name, description) VALUES
		('CyberWarriors', 'Experts in cyber warfare and defense'),
		('ByteBandits', 'Masters of data breaches and exfiltration');

		-- Добавление новых соревнований
		INSERT INTO games (start_time, end_time, description) VALUES
		('2023-10-03 12:00:00', '2023-10-03 15:00:00', 'Break the Encryption'),
		('2023-10-04 12:00:00', '2023-10-04 15:00:00', 'Find the Exploit');
	`

	_, err = tx.Exec(query)
	if err != nil {
		tx.Rollback()
		slog.Error("Problem with query execution, query: " + query + "\n   error:" + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}

	// Получение существующих команд из базы данных
	teamRows, err := tx.Query(`SELECT id, name FROM teams`)
	if err != nil {
		tx.Rollback()
		slog.Error("Failed to fetch existing teams: " + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}
	defer teamRows.Close()

	var allTeams []struct {
		ID   string
		Name string
	}

	for teamRows.Next() {
		var id, name string
		if err := teamRows.Scan(&id, &name); err != nil {
			tx.Rollback()
			slog.Error("Failed to scan team: " + err.Error())
			return fromUpdateId, toUpdateId, description, err
		}
		allTeams = append(allTeams, struct {
			ID   string
			Name string
		}{ID: id, Name: name})
	}

	// Получение существующих игр из базы данных
	gameRows, err := tx.Query(`SELECT id, description FROM games`)
	if err != nil {
		tx.Rollback()
		slog.Error("Failed to fetch existing games: " + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}
	defer gameRows.Close()

	var allGames []struct {
		ID          string
		Description string
	}

	for gameRows.Next() {
		var id, description string
		if err := gameRows.Scan(&id, &description); err != nil {
			tx.Rollback()
			slog.Error("Failed to scan game: " + err.Error())
			return fromUpdateId, toUpdateId, description, err
		}
		allGames = append(allGames, struct {
			ID          string
			Description string
		}{ID: id, Description: description})
	}

	for _, game := range allGames {
		for rank, team := range allTeams {
			score := rand.Float64() * 100
			_, err = tx.Exec(`INSERT INTO results (score, rank, id, team_id, game_id) VALUES ($1, $2, $3, $4, $5)`,
				score, rank+1, uuid.New().String(), team.ID, game.ID)
			if err != nil {
				tx.Rollback()
				slog.Error("Failed to insert result: " + err.Error())
				return fromUpdateId, toUpdateId, description, err
			}

			// Добавление команды в игру
			_, err = tx.Exec(`INSERT INTO team_games (team_id, game_id) VALUES ($1, $2)`,
				team.ID, game.ID)
			if err != nil {
				tx.Rollback()
				slog.Error("Failed to insert team_game: " + err.Error())
				return fromUpdateId, toUpdateId, description, err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		slog.Error("Failed to commit transaction: " + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}

	return fromUpdateId, toUpdateId, description, nil
}
