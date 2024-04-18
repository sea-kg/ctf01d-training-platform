package repository

import (
	"context"
	"ctf01d/lib/models"
	"database/sql"
)

type GameRepository interface {
	Create(ctx context.Context, game *models.Game) error
	GetById(ctx context.Context, id string) (*models.Game, error)
	Update(ctx context.Context, game *models.Game) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]models.Game, error)
}

type gameRepo struct {
	db *sql.DB
}

func NewGameRepository(db *sql.DB) GameRepository {
	return &gameRepo{db: db}
}

func (r *gameRepo) Create(ctx context.Context, game *models.Game) error {
	query := `INSERT INTO games (start_time, end_time, description) VALUES (?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, game.StartTime, game.EndTime, game.Description)
	return err
}

func (r *gameRepo) GetById(ctx context.Context, id string) (*models.Game, error) {
	query := `SELECT id, start_time, end_time, description FROM games WHERE id = ?`
	game := &models.Game{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&game.Id, &game.StartTime, &game.EndTime, &game.Description)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (r *gameRepo) Update(ctx context.Context, game *models.Game) error {
	query := `UPDATE games SET start_time = ?, end_time = ?, description = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, game.StartTime, game.EndTime, game.Description, game.Id)
	return err
}

func (r *gameRepo) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM games WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *gameRepo) List(ctx context.Context) ([]models.Game, error) {
	query := `SELECT id, start_time, end_time, description FROM games`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []models.Game
	for rows.Next() {
		var game models.Game
		if err := rows.Scan(&game.Id, &game.StartTime, &game.EndTime, &game.Description); err != nil {
			return nil, err
		}
		games = append(games, game)
	}
	return games, nil
}
