package repository

import (
	"context"
	"database/sql"

	models "ctf01d/internal/app/db"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type ResultRepository interface {
	Create(ctx context.Context, result *models.Result) error
	GetById(ctx context.Context, id openapi_types.UUID) (*models.Result, error)
	Update(ctx context.Context, result *models.Result) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, gameId openapi_types.UUID) ([]*models.Result, error)
}

type resultRepo struct {
	db *sql.DB
}

func NewResultRepository(db *sql.DB) ResultRepository {
	return &resultRepo{db: db}
}

func (r *resultRepo) Create(ctx context.Context, result *models.Result) error {
	query := `INSERT INTO results (team_id, game_id, score)
	          VALUES ($1, $2, $3)
	          RETURNING id, team_id, game_id, score`
	row := r.db.QueryRowContext(ctx, query, result.TeamId, result.GameId, result.Score)
	err := row.Scan(&result.Id, &result.TeamId, &result.GameId, &result.Score)
	if err != nil {
		return err
	}
	return nil
}

func (r *resultRepo) GetById(ctx context.Context, gameId openapi_types.UUID) (*models.Result, error) {
	query := `SELECT id, team_id, game_id, score FROM results WHERE game_id = $1 order by score desc`
	result := &models.Result{}
	err := r.db.QueryRowContext(ctx, query, gameId).Scan(&result.Id, &result.TeamId, &result.GameId, &result.Score)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *resultRepo) Update(ctx context.Context, result *models.Result) error {
	query := `UPDATE results SET team_id = $1, game_id = $2, score = $3 WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query, result.TeamId, result.GameId, result.Score, result.Id)
	return err
}

func (r *resultRepo) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM results WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *resultRepo) List(ctx context.Context, gameId openapi_types.UUID) ([]*models.Result, error) {
	query := `SELECT id, team_id, game_id, score FROM results WHERE game_id = $1 ORDER BY score DESC`
	rows, err := r.db.QueryContext(ctx, query, gameId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*models.Result
	for rows.Next() {
		var result models.Result
		if err := rows.Scan(&result.Id, &result.TeamId, &result.GameId, &result.Score); err != nil {
			return nil, err
		}
		results = append(results, &result)
	}
	return results, nil
}
