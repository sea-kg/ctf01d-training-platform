package repository

import (
	"context"
	"ctf01d/lib/models"
	"database/sql"
)

type ServiceRepository interface {
	Create(ctx context.Context, service *models.Service) error
	GetById(ctx context.Context, id string) (*models.Service, error)
	Update(ctx context.Context, service *models.Service) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]models.Service, error)
}

type serviceRepo struct {
	db *sql.DB
}

func NewServiceRepository(db *sql.DB) ServiceRepository {
	return &serviceRepo{db: db}
}

func (r *serviceRepo) Create(ctx context.Context, service *models.Service) error {
	// fixme check params
	query := `INSERT INTO services (name author, logo_url, description, is_public) VALUES (?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, service.Name, service.Author, service.LogoUrl, service.Description, service.IsPublic)
	return err
}

func (r *serviceRepo) GetById(ctx context.Context, id string) (*models.Service, error) {
	// fixme - check params
	query := `SELECT id name, author, logo_url, description, is_public FROM services WHERE id = ?`
	service := &models.Service{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&service.Id, &service.Name, &service.Author, &service.LogoUrl, &service.Description, &service.IsPublic)
	if err != nil {
		return nil, err
	}
	return service, nil
}

func (r *serviceRepo) Update(ctx context.Context, service *models.Service) error {
	// fixme - check params
	query := `UPDATE services SET start_time = ?, end_time = ?, description = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, service.Name, service.Author, service.LogoUrl, service.Description, service.IsPublic)
	return err
}

func (r *serviceRepo) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM services WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *serviceRepo) List(ctx context.Context) ([]models.Service, error) {
	// fixme - check params
	query := `SELECT id, start_time, end_time, description FROM services`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []models.Service
	for rows.Next() {
		var service models.Service
		if err := rows.Scan(&service.Id, &service.Id, &service.Name, &service.Author, &service.LogoUrl, &service.Description, &service.IsPublic); err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	return services, nil
}
