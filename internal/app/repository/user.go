package repository

import (
	"context"
	"ctf01d/internal/app/models"
	"database/sql"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	AddUserToTeams(ctx context.Context, userId int, teamIds []string) error
	GetById(ctx context.Context, id int) (*models.User, error)
	GetByUserName(ctx context.Context, id string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]*models.User, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (user_name, avatar_url, role, status, password_hash) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, user.Username, user.AvatarUrl, user.Role, user.Status, user.PasswordHash).Scan(&user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) AddUserToTeams(ctx context.Context, userId int, teamIds []string) error {
	for _, teamId := range teamIds {
		_, err := r.db.ExecContext(ctx, "INSERT INTO team_members (user_id, team_id) VALUES ($1, $2)", userId, teamId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *userRepo) GetById(ctx context.Context, id int) (*models.User, error) {
	query := `SELECT id, user_name, avatar_url, role, status FROM users WHERE id = $1`
	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.Id, &user.Username, &user.AvatarUrl, &user.Role, &user.Status)
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
	query := `UPDATE users SET user_name = $1, avatar_url = $2, role = $3, status = $4, password_hash = $5 WHERE id = $6`
	_, err := r.db.ExecContext(ctx, query, user.Username, user.AvatarUrl, user.Role, user.Status, user.PasswordHash, user.Id)
	return err
}

func (r *userRepo) Delete(ctx context.Context, id int) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, "DELETE FROM team_members WHERE user_id = $1", id); err != nil {
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
	query := `SELECT id, user_name, avatar_url, role, status FROM users`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.AvatarUrl, &user.Role, &user.Status); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
