package repository

import (
	"context"
	"ctf01d/internal/app/models"
	"database/sql"
)

type TeamRepository interface {
	Create(ctx context.Context, team *models.Team) error
	GetById(ctx context.Context, id string) (*models.Team, error)
	Update(ctx context.Context, team *models.Team) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*models.Team, error)
}

type teamRepo struct {
	db *sql.DB
}

func NewTeamRepository(db *sql.DB) TeamRepository {
	return &teamRepo{db: db}
}

func (r *teamRepo) Create(ctx context.Context, team *models.Team) error {
	query := `INSERT INTO teams (name, description, university_id, social_links, avatar_url) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query, team.Name, team.Description, team.UniversityId, team.SocialLinks, team.AvatarUrl)
	return err
}

func (r *teamRepo) GetById(ctx context.Context, id string) (*models.Team, error) {
	query := "SELECT t.*, u.name as university_name FROM teams t JOIN universities u ON t.university_id = u.id"
	team := &models.Team{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&team.Id, &team.Name, &team.Description, &team.UniversityId, &team.University, &team.SocialLinks, &team.AvatarUrl)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (r *teamRepo) Update(ctx context.Context, team *models.Team) error {
	query := `UPDATE teams SET name = $1, description = $2 university = $3 social_links = $4 avatar_url = $5 WHERE id = $6`
	_, err := r.db.ExecContext(ctx, query, team.Name, team.Description, team.UniversityId, team.SocialLinks, team.AvatarUrl, team.Id)
	return err
}

func (r *teamRepo) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM teams WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *teamRepo) List(ctx context.Context) ([]*models.Team, error) {
	query := "SELECT t.*, u.name as university_name FROM teams t JOIN universities u ON t.university_id = u.id"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []*models.Team
	for rows.Next() {
		var team models.Team
		if err := rows.Scan(&team.Id, &team.Name, &team.Description, &team.UniversityId, &team.SocialLinks, &team.AvatarUrl, &team.University); err != nil {
			return nil, err
		}
		teams = append(teams, &team)
	}
	return teams, nil
}