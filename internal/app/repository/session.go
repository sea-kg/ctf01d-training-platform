package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type SessionRepository interface {
	GetSessionFromDB(ctx context.Context, sessionId string) (int, error)
	StoreSessionInDB(ctx context.Context, userId int) (string, error)
	DeleteSessionInDB(ctx context.Context, cookie string) error
}

type sessionRepo struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) SessionRepository {
	return &sessionRepo{db: db}
}

func (r *sessionRepo) GetSessionFromDB(ctx context.Context, sessionId string) (int, error) {
	var userId int
	err := r.db.QueryRowContext(ctx, "SELECT user_id FROM sessions WHERE id = $1 AND expires_at > NOW()", sessionId).Scan(&userId)
	return userId, err
}

func (r *sessionRepo) StoreSessionInDB(ctx context.Context, userId int) (string, error) {
	var session string
	query := `
		INSERT INTO sessions (user_id, expires_at)
		VALUES ($1, $2)
		ON CONFLICT (user_id) DO
		UPDATE SET expires_at = EXCLUDED.expires_at
		RETURNING id
		`
	err := r.db.QueryRowContext(ctx, query, userId, time.Now().Add(96*time.Hour)).Scan(&session)
	fmt.Println(session)
	if err != nil {
		return "", err
	}
	return session, nil
}

// fixme - implement me
func (r *sessionRepo) DeleteSessionInDB(ctx context.Context, cookie string) error {
	return nil
}
