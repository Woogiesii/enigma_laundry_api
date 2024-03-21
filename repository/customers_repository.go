package repository

import (
	"database/sql"
	"enigma_laundry_api/model"
	"time"
)

type UsersRepository interface {
	Get(id string) (model.Users, error)
	GetByUsername(username string) (model.Users, error)
	Create(payload model.Users) (model.Users, error)
}

type usersRepository struct {
	db *sql.DB
}

func (cst *usersRepository) Get(id string) (model.Users, error) {
	var customer model.Users
	err := cst.db.QueryRow(`SELECT id, customer_name, phone_number, username, password, role, date_created FROM mst_customers WHERE id = $1`, id).Scan(
		&customer.Id,
		&customer.CustomerName,
		&customer.PhoneNumber,
		&customer.Username,
		&customer.Password,
		&customer.Role,
		&customer.DateCreated,
	)

	if err != nil {
		return model.Users{}, err
	}

	return customer, nil
}

func (cst *usersRepository) GetByUsername(username string) (model.Users, error) {
	var customer model.Users
	err := cst.db.QueryRow(`SELECT id, customer_name, phone_number, username, password, role, date_created FROM mst_customers WHERE username = $1`, username).Scan(
		&customer.Id,
		&customer.CustomerName,
		&customer.PhoneNumber,
		&customer.Username,
		&customer.Password,
		&customer.Role,
		&customer.DateCreated,
	)
	if err != nil {
		return model.Users{}, err
	}
	return customer, nil
}

func (cst *usersRepository) Create(payload model.Users) (model.Users, error) {
	var customer model.Users
	err := cst.db.QueryRow(`INSERT INTO mst_customers (customer_name, phone_number, username, password, role, date_created) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, customer_name, phone_number, username, password, role, date_created`,
		payload.CustomerName,
		payload.PhoneNumber,
		payload.Username,
		payload.Password,
		payload.Role,
		time.Now(),
	).Scan(
		&customer.Id,
		&customer.CustomerName,
		&customer.PhoneNumber,
		&customer.Username,
		&customer.Password,
		&customer.Role,
		&customer.DateCreated,
	)

	if err != nil {
		return model.Users{}, err
	}
	return customer, nil
}

func NewUsersRepository(db *sql.DB) UsersRepository {
	return &usersRepository{db: db}
}
