package repository

import (
	"context"
	"database/sql"

	models "ctf01d/internal/app/db"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	AddUserToTeams(ctx context.Context, userId openapi_types.UUID, teamIds *[]openapi_types.UUID) error
	GetById(ctx context.Context, id openapi_types.UUID) (*models.User, error)
	GetProfileWithHistory(ctx context.Context, id openapi_types.UUID) (*models.ProfileWithHistory, error)
	GetByUserName(ctx context.Context, id string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id openapi_types.UUID) error
	List(ctx context.Context) ([]*models.User, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (display_name, user_name, avatar_url, role, status, password_hash) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, user.DisplayName, user.Username, user.AvatarUrl, user.Role, user.Status, user.PasswordHash).Scan(&user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) AddUserToTeams(ctx context.Context, userId openapi_types.UUID, teamIds *[]openapi_types.UUID) error {
	for _, teamId := range *teamIds {
		_, err := r.db.ExecContext(ctx, "INSERT INTO profiles (user_id, current_team_id) VALUES ($1, $2)", userId, teamId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *userRepo) GetProfileWithHistory(ctx context.Context, id openapi_types.UUID) (*models.ProfileWithHistory, error) {
	query := `
		SELECT teams.name, created_at, updated_at
		FROM profiles JOIN teams on profiles.current_team_id=teams.id
		WHERE profiles.user_id = $1
	`
	profile := models.Profile{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&profile.CurrentTeam, &profile.CreatedAt, &profile.UpdatedAt)
	if err != nil {
		return nil, err
	}
	query = `
	SELECT joined_at, left_at, name
		FROM team_history
		JOIN teams ON teams.id = team_history.team_id
		WHERE user_id = $1
	`
	rows, err := r.db.QueryContext(ctx, query, id)
	var history []models.ProfileTeams
	for rows.Next() {
		var team models.ProfileTeams
		err := rows.Scan(&team.JoinedAt, &team.LeftAt, &team.Name)
		if err != nil {
			return nil, err
		}
		history = append(history, team)
	}
	if err != nil {
		return nil, err
	}
	return &models.ProfileWithHistory{
		Profile: profile,
		History: history,
	}, nil
}

func (r *userRepo) GetById(ctx context.Context, id openapi_types.UUID) (*models.User, error) {
	query := `SELECT id, display_name, user_name, avatar_url, role, status FROM users WHERE id = $1`
	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.Id, &user.DisplayName, &user.Username, &user.AvatarUrl, &user.Role, &user.Status)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepo) GetByUserName(ctx context.Context, name string) (*models.User, error) {
	query := `SELECT id, password_hash FROM users WHERE user_name = $1`
	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, name).Scan(&user.Id, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepo) Update(ctx context.Context, user *models.User) error {
	query := `UPDATE users SET user_name = $1, avatar_url = $2, role = $3, status = $4, password_hash = $5, display_name = $6 WHERE id = $7`
	_, err := r.db.ExecContext(ctx, query, user.Username, user.AvatarUrl, user.Role, user.Status, user.PasswordHash, user.DisplayName, user.Id)
	return err
}

func (r *userRepo) Delete(ctx context.Context, id openapi_types.UUID) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, "DELETE FROM profiles WHERE user_id = $1", id); err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}
	if _, err := tx.ExecContext(ctx, "DELETE FROM users WHERE id = $1", id); err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}
	return tx.Commit()
}

func (r *userRepo) List(ctx context.Context) ([]*models.User, error) {
	query := `SELECT id, display_name, user_name, avatar_url, role, status FROM users`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.DisplayName, &user.Username, &user.AvatarUrl, &user.Role, &user.Status); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
