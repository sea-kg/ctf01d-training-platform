package repository

import (
	"context"
	"database/sql"
	"time"

	models "ctf01d/internal/app/db"

	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type GameRepository interface {
	Create(ctx context.Context, game *models.Game) error
	GetById(ctx context.Context, id openapi_types.UUID) (*models.Game, error)
	GetGameDetails(ctx context.Context, id openapi_types.UUID) (*models.GameDetails, error)
	Update(ctx context.Context, game *models.Game) error
	Delete(ctx context.Context, id openapi_types.UUID) error
	List(ctx context.Context) ([]*models.Game, error)
}

type gameRepo struct {
	db *sql.DB
}

func NewGameRepository(db *sql.DB) GameRepository {
	return &gameRepo{db: db}
}

func (r *gameRepo) Create(ctx context.Context, game *models.Game) error {
	query := `INSERT INTO games (start_time, end_time, description)
	          VALUES ($1, $2, $3)
	          RETURNING id, start_time, end_time, description`
	row := r.db.QueryRowContext(ctx, query, game.StartTime, game.EndTime, game.Description)
	err := row.Scan(&game.Id, &game.StartTime, &game.EndTime, &game.Description)
	if err != nil {
		return err
	}
	return nil
}

func (r *gameRepo) GetGameDetails(ctx context.Context, id openapi_types.UUID) (*models.GameDetails, error) {
	query := `
        SELECT g.id, g.start_time, g.end_time, g.description, t.id, t.name, t.description
        FROM games g
        LEFT JOIN team_games tg ON g.id = tg.game_id
        LEFT JOIN teams t ON tg.team_id = t.id
        WHERE g.id = $1;
    `
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	gameDetails := &models.GameDetails{}
	teams := make(map[uuid.UUID]models.Team)

	for rows.Next() {
		var gameId openapi_types.UUID
		var startTime, endTime time.Time
		var description string
		var teamId sql.NullString
		var teamName sql.NullString
		var teamDescription sql.NullString

		err := rows.Scan(&gameId, &startTime, &endTime, &description, &teamId, &teamName, &teamDescription)
		if err != nil {
			return nil, err
		}

		gameDetails.Id = id
		gameDetails.StartTime = startTime
		gameDetails.EndTime = endTime
		gameDetails.Description = description
		if teamId.Valid {
			teamUUID, err := uuid.Parse(teamId.String)
			if err != nil {
				return nil, err
			}
			team := models.Team{
				Id:          teamUUID,
				Name:        teamName.String,
				Description: teamDescription.String,
			}
			teams[teamUUID] = team
		}
	}

	for _, team := range teams {
		gameDetails.Teams = append(gameDetails.Teams, &team)
	}

	return gameDetails, nil
}

func (r *gameRepo) GetById(ctx context.Context, id openapi_types.UUID) (*models.Game, error) {
	query := `SELECT id, start_time, end_time, description FROM games WHERE id = $1`
	game := &models.Game{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&game.Id, &game.StartTime, &game.EndTime, &game.Description)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (r *gameRepo) Update(ctx context.Context, game *models.Game) error {
	query := `UPDATE games SET start_time = $1, end_time = $2, description = $3 WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query, game.StartTime, game.EndTime, game.Description, game.Id)
	return err
}

func (r *gameRepo) Delete(ctx context.Context, id openapi_types.UUID) error {
	query := `DELETE FROM games WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *gameRepo) List(ctx context.Context) ([]*models.Game, error) {
	query := `SELECT id, start_time, end_time, description FROM games`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []*models.Game
	for rows.Next() {
		var game models.Game
		if err := rows.Scan(&game.Id, &game.StartTime, &game.EndTime, &game.Description); err != nil {
			return nil, err
		}
		games = append(games, &game)
	}
	return games, nil
}
