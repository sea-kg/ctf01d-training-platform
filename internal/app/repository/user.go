package repository

import (
	"context"
	"ctf01d/internal/app/models"
	"database/sql"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetById(ctx context.Context, id string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*models.User, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (user_name, avatar_url, role, status, password_hash) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query, user.Username, user.AvatarUrl, user.Role, user.Status, user.PasswordHash)
	return err
}

func (r *userRepo) GetById(ctx context.Context, id string) (*models.User, error) {
	query := `SELECT id, user_name, avatar_url, role, status FROM users WHERE id = $1`
	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.Id, &user.Username, &user.AvatarUrl, &user.Role, &user.Status)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepo) Update(ctx context.Context, user *models.User) error {
	query := `UPDATE users SET user_name = $1, avatar_url = $2, role = $3 WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query, user.Username, user.AvatarUrl, user.Role, user.Id)
	return err
}

func (r *userRepo) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
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
