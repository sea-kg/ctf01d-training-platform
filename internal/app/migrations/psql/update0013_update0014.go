package migration

import (
	"database/sql"
	"log/slog"
	"runtime"
)

func DatabaseUpdate_update0013_update0014(db *sql.DB, getInfo bool) (string, string, string, error) {
	// WARNING!!!
	// Do not change the update if it has already been installed by other developers or in production.
	// To correct the database, create a new update and register it in the list of updates.

	fromUpdateId, toUpdateId := ParseNameFuncUpdate(runtime.Caller(0))
	description := "Change int id to string uuid in all table"
	if getInfo {
		return fromUpdateId, toUpdateId, description, nil
	}
	query := `
		BEGIN;

		ALTER TABLE results ADD COLUMN new_id UUID DEFAULT gen_random_uuid();
		UPDATE results SET new_id = gen_random_uuid();
		ALTER TABLE results DROP COLUMN id;
		ALTER TABLE results ADD CONSTRAINT new_id_pkey PRIMARY KEY (new_id);
		ALTER TABLE results RENAME COLUMN new_id TO id;


		ALTER TABLE games ADD COLUMN new_id UUID DEFAULT gen_random_uuid();
		ALTER TABLE services ADD COLUMN new_id UUID DEFAULT gen_random_uuid();
		ALTER TABLE sessions ADD COLUMN new_id UUID DEFAULT gen_random_uuid();
		ALTER TABLE teams ADD COLUMN new_id UUID DEFAULT gen_random_uuid();
		ALTER TABLE universities ADD COLUMN new_id UUID DEFAULT gen_random_uuid();
		ALTER TABLE users ADD COLUMN new_id UUID DEFAULT gen_random_uuid();


		ALTER TABLE team_games ADD COLUMN new_team_id UUID;
		ALTER TABLE team_games ADD COLUMN new_game_id UUID;
		---
		UPDATE team_games tg
		SET new_team_id = t.new_id
		FROM teams t
		WHERE tg.team_id = t.id;
		---
		UPDATE team_games tg
		SET new_game_id = g.new_id
		FROM games g
		WHERE tg.game_id = g.id;
		---
		ALTER TABLE team_games DROP COLUMN team_id;
		ALTER TABLE team_games DROP COLUMN game_id;
		ALTER TABLE team_games RENAME COLUMN new_team_id TO team_id;
		ALTER TABLE team_games RENAME COLUMN new_game_id TO game_id;


		ALTER TABLE team_members ADD COLUMN new_team_id UUID;
		ALTER TABLE team_members ADD COLUMN new_user_id UUID;
		---
		UPDATE team_members tm
		SET new_team_id = t.new_id
		FROM teams t
		WHERE tm.team_id = t.id;
		---
		UPDATE team_members tm
		SET new_user_id = u.new_id
		FROM users u
		WHERE tm.user_id = u.id;
		---
		ALTER TABLE team_members DROP COLUMN team_id;
		ALTER TABLE team_members DROP COLUMN user_id;
		ALTER TABLE team_members RENAME COLUMN new_team_id TO team_id;
		ALTER TABLE team_members RENAME COLUMN new_user_id TO user_id;


		ALTER TABLE game_services ADD COLUMN new_game_id UUID;
		ALTER TABLE game_services ADD COLUMN new_service_id UUID;
		---
		UPDATE game_services gm
		SET new_game_id = g.new_id
		FROM games g
		WHERE gm.game_id = g.id;
		---
		UPDATE game_services gs
		SET new_service_id = s.new_id
		FROM services s
		WHERE gs.service_id = s.id;
		---
		ALTER TABLE game_services DROP COLUMN game_id;
		ALTER TABLE game_services DROP COLUMN service_id;
		ALTER TABLE game_services RENAME COLUMN new_game_id TO game_id;
		ALTER TABLE game_services RENAME COLUMN new_service_id TO service_id;


		ALTER TABLE results drop constraint results_game_id_fkey;
		ALTER TABLE results drop constraint results_team_id_fkey;
		ALTER TABLE results DROP COLUMN team_id;
		ALTER TABLE results DROP COLUMN game_id;
		ALTER TABLE results ADD COLUMN team_id UUID not null;
		ALTER TABLE results ADD COLUMN game_id UUID not null;

		ALTER TABLE sessions drop constraint sessions_user_id_fkey;
		ALTER TABLE sessions DROP COLUMN user_id;
		ALTER TABLE sessions ADD COLUMN user_id UUID not null;

		ALTER TABLE teams DROP constraint teams_university_id_fkey;
		ALTER TABLE teams DROP COLUMN university_id;
		ALTER TABLE teams ADD COLUMN university_id UUID;

		ALTER TABLE games DROP COLUMN id;
		ALTER TABLE games RENAME COLUMN new_id TO id;
		ALTER TABLE teams DROP COLUMN id;
		ALTER TABLE teams RENAME COLUMN new_id TO id;
		ALTER TABLE users DROP COLUMN id;
		ALTER TABLE users RENAME COLUMN new_id TO id;
		ALTER TABLE universities DROP COLUMN id;
		ALTER TABLE universities RENAME COLUMN new_id TO id;
		ALTER TABLE services DROP COLUMN id;
		ALTER TABLE services RENAME COLUMN new_id TO id;


		ALTER TABLE games ADD CONSTRAINT games_id_unique UNIQUE (id);
		ALTER TABLE teams ADD CONSTRAINT teams_id_unique UNIQUE (id);
		ALTER TABLE users ADD CONSTRAINT users_id_unique UNIQUE (id);
		ALTER TABLE universities ADD CONSTRAINT universities_id_unique UNIQUE (id);
		ALTER TABLE services ADD CONSTRAINT services_id_unique UNIQUE (id);


		ALTER TABLE results ADD CONSTRAINT results_game_id_fkey FOREIGN KEY (game_id) REFERENCES games(id);
		ALTER TABLE results ADD CONSTRAINT results_team_id_fkey FOREIGN KEY (team_id) REFERENCES teams(id);
		ALTER TABLE sessions ADD CONSTRAINT sessions_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id);
		ALTER TABLE teams ADD CONSTRAINT teams_university_id_fkey FOREIGN KEY (university_id) REFERENCES universities(id);

		COMMIT;
	`
	_, err := db.Exec(query)
	if err != nil {
		slog.Error("Problem with update, query: " + query + "\n   error:" + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}

	return fromUpdateId, toUpdateId, description, nil
}
