package repository

import (
	"context"
	"database/sql"

	"ctf01d/internal/model"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type TeamRepository interface {
	Create(ctx context.Context, team *model.Team) error
	GetById(ctx context.Context, id openapi_types.UUID) (*model.Team, error)
	Update(ctx context.Context, team *model.Team) error
	Delete(ctx context.Context, id openapi_types.UUID) error
	List(ctx context.Context) ([]*model.Team, error)
}

type teamRepo struct {
	db *sql.DB
}

func NewTeamRepository(db *sql.DB) TeamRepository {
	return &teamRepo{db: db}
}

func (r *teamRepo) Create(ctx context.Context, team *model.Team) error {
	query := `INSERT INTO teams (name, description, university_id, social_links, avatar_url)
	          VALUES ($1, $2, $3, $4, $5)
	          RETURNING id, name, description, university_id, social_links, avatar_url`
	row := r.db.QueryRowContext(ctx, query, team.Name, team.Description, team.UniversityId, team.SocialLinks, team.AvatarUrl)
	err := row.Scan(&team.Id, &team.Name, &team.Description, &team.UniversityId, &team.SocialLinks, &team.AvatarUrl)
	if err != nil {
		return err
	}
	return nil
}

func (r *teamRepo) GetById(ctx context.Context, id openapi_types.UUID) (*model.Team, error) {
	query := `SELECT t.id, t.name, t.description, t.social_links, t.avatar_url, u.name as university_name
			FROM teams t
			LEFT JOIN universities u ON t.university_id = u.id
			WHERE t.id = $1`
	team := &model.Team{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&team.Id, &team.Name, &team.Description, &team.SocialLinks, &team.AvatarUrl, &team.University)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (r *teamRepo) Update(ctx context.Context, team *model.Team) error {
	query := `UPDATE teams SET name = $1, description = $2, university_id = $3, social_links = $4, avatar_url = $5 WHERE id = $6`
	_, err := r.db.ExecContext(ctx, query, team.Name, team.Description, team.UniversityId, team.SocialLinks, team.AvatarUrl, team.Id)
	return err
}

func (r *teamRepo) Delete(ctx context.Context, id openapi_types.UUID) error {
	query := `DELETE FROM teams WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *teamRepo) List(ctx context.Context) ([]*model.Team, error) {
	query := "SELECT t.*, u.name as university_name FROM teams t LEFT JOIN universities u ON t.university_id = u.id"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []*model.Team
	for rows.Next() {
		var team model.Team
		if err := rows.Scan(&team.Name, &team.Description, &team.SocialLinks, &team.AvatarUrl, &team.Id, &team.UniversityId, &team.University); err != nil {
			return nil, err
		}
		teams = append(teams, &team)
	}
	return teams, nil
}
