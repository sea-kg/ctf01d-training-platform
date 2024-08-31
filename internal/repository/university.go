package repository

import (
	"context"
	"database/sql"

	models "ctf01d/internal/model"
)

type UniversityRepository interface {
	List(ctx context.Context) ([]*models.University, error)
	Search(ctx context.Context, query string) ([]*models.University, error)
}

type universityRepo struct {
	db *sql.DB
}

func NewUniversityRepository(db *sql.DB) UniversityRepository {
	return &universityRepo{db: db}
}

func (repo *universityRepo) Search(ctx context.Context, query string) ([]*models.University, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name FROM universities WHERE name ILIKE '%' || $1 || '%' LIMIT 10", query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var universities []*models.University
	for rows.Next() {
		var u models.University
		if err := rows.Scan(&u.Id, &u.Name); err != nil {
			return nil, err
		}
		universities = append(universities, &u)
	}
	return universities, nil
}

func (repo *universityRepo) List(ctx context.Context) ([]*models.University, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name FROM universities LIMIT 10")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var universities []*models.University
	for rows.Next() {
		var u models.University
		if err := rows.Scan(&u.Id, &u.Name); err != nil {
			return nil, err
		}
		universities = append(universities, &u)
	}
	return universities, nil
}
