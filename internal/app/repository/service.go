package repository

import (
	"context"
	"database/sql"

	models "ctf01d/internal/app/db"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type ServiceRepository interface {
	Create(ctx context.Context, service *models.Service) error
	GetById(ctx context.Context, id openapi_types.UUID) (*models.Service, error)
	Update(ctx context.Context, service *models.Service) error
	Delete(ctx context.Context, id openapi_types.UUID) error
	List(ctx context.Context) ([]*models.Service, error)
}

type serviceRepo struct {
	db *sql.DB
}

func NewServiceRepository(db *sql.DB) ServiceRepository {
	return &serviceRepo{db: db}
}

func (r *serviceRepo) Create(ctx context.Context, service *models.Service) error {
	query := `INSERT INTO services (name, author, logo_url, description, is_public)
	          VALUES ($1, $2, $3, $4, $5)
	          RETURNING id, name, author, logo_url, description, is_public`
	row := r.db.QueryRowContext(ctx, query, service.Name, service.Author, service.LogoUrl, service.Description, service.IsPublic)
	err := row.Scan(&service.Id, &service.Name, &service.Author, &service.LogoUrl, &service.Description, &service.IsPublic)
	if err != nil {
		return err
	}
	return nil
}

func (r *serviceRepo) GetById(ctx context.Context, id openapi_types.UUID) (*models.Service, error) {
	query := `SELECT id, name, author, logo_url, description, is_public FROM services WHERE id = $1`
	service := &models.Service{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&service.Id, &service.Name, &service.Author, &service.LogoUrl, &service.Description, &service.IsPublic)
	if err != nil {
		return nil, err
	}
	return service, nil
}

func (r *serviceRepo) Update(ctx context.Context, service *models.Service) error {
	query := `UPDATE services SET name = $1, author = $2, logo_url = $3, description = $4, is_public = $5 WHERE id = $6`
	_, err := r.db.ExecContext(ctx, query, service.Name, service.Author, service.LogoUrl, service.Description, service.IsPublic, service.Id)
	return err
}

func (r *serviceRepo) Delete(ctx context.Context, id openapi_types.UUID) error {
	query := `DELETE FROM services WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *serviceRepo) List(ctx context.Context) ([]*models.Service, error) {
	query := `SELECT id, name, author, logo_url, description, is_public FROM services`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []*models.Service
	for rows.Next() {
		var service models.Service
		if err := rows.Scan(&service.Id, &service.Name, &service.Author, &service.LogoUrl, &service.Description, &service.IsPublic); err != nil {
			return nil, err
		}
		services = append(services, &service)
	}
	return services, nil
}
