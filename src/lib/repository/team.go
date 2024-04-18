package repository

import (
	"context"
	"ctf01d/lib/models"
	"database/sql"
)

type TeamRepository interface {
	Create(ctx context.Context, team *models.Team) error
	GetById(ctx context.Context, id string) (*models.Team, error)
	Update(ctx context.Context, team *models.Team) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]models.Team, error)
}

type teamRepo struct {
	db *sql.DB
}

func NewTeamRepository(db *sql.DB) TeamRepository {
	return &teamRepo{db: db}
}

func (r *teamRepo) Create(ctx context.Context, team *models.Team) error {
	query := `INSERT INTO teams (name, description) VALUES (?, ?)`
	_, err := r.db.ExecContext(ctx, query, team.TeamName, team.Description)
	return err
}

func (r *teamRepo) GetById(ctx context.Context, id string) (*models.Team, error) {
	query := `SELECT id, name, description FROM teams WHERE id = ?`
	team := &models.Team{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&team.Id, &team.TeamName, &team.Description)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (r *teamRepo) Update(ctx context.Context, team *models.Team) error {
	query := `UPDATE teams SET name = ?, description = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, team.TeamName, team.Description, team.Id)
	return err
}

func (r *teamRepo) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM teams WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *teamRepo) List(ctx context.Context) ([]models.Team, error) {
	query := `SELECT id, name, description FROM teams`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []models.Team
	for rows.Next() {
		var team models.Team
		if err := rows.Scan(&team.Id, &team.TeamName, &team.Description); err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}
	return teams, nil
}
