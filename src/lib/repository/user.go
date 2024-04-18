package repository

import (
	"context"
	"ctf01d/lib/models"
	"database/sql"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetById(ctx context.Context, id string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]models.User, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (username, avatar_url, role) VALUES (?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, user.Username, user.AvatarUrl, user.Role)
	return err
}

func (r *userRepo) GetById(ctx context.Context, id string) (*models.User, error) {
	query := `SELECT id, username, avatar_url, role FROM users WHERE id = ?`
	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.Id, &user.Username, &user.AvatarUrl, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepo) Update(ctx context.Context, user *models.User) error {
	query := `UPDATE users SET username = ?, avatar_url = ?, role = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, user.Username, user.AvatarUrl, user.Role, user.Id)
	return err
}

func (r *userRepo) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *userRepo) List(ctx context.Context) ([]models.User, error) {
	query := `SELECT id, username, avatar_url, role FROM users`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.AvatarUrl, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
