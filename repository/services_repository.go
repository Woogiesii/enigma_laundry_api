package repository

import (
	"database/sql"
	"enigma_laundry_api/model"
)

type ServicesRepository interface {
	Get(id string) (model.Services, error)
	Create(payload model.Services) (model.Services, error)
}

type servicesRepository struct {
	db *sql.DB
}

func (serv *servicesRepository) Get(id string) (model.Services, error) {
	var services model.Services
	err := serv.db.QueryRow(`SELECT id, service_name, unit, price FROM mst_services WHERE id = $1`, id).Scan(
		&services.Id,
		&services.ServiceName,
		&services.Unit,
		&services.Price,
	)

	if err != nil {
		return model.Services{}, err
	}

	return services, nil
}

func (serv *servicesRepository) Create(payload model.Services) (model.Services, error) {
	var services model.Services
	err := serv.db.QueryRow(`INSERT INTO mst_services (service_name, unit, price) VALUES ($1, $2, $3) RETURNING id, service_name, unit, price`,
		payload.ServiceName,
		payload.Unit,
		payload.Price,
	).Scan(
		&services.Id,
		&services.ServiceName,
		&services.Unit,
		&services.Price,
	)

	if err != nil {
		return model.Services{}, err
	}

	return services, nil
}

func NewServicesRepository(db *sql.DB) ServicesRepository {
	return &servicesRepository{db: db}
}
