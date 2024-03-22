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
	Update(payload model.Users) (model.Users, error)
	Delete(id string) (model.Users, error)
}

type usersRepository struct {
	db *sql.DB
}

func (cst *usersRepository) Get(id string) (model.Users, error) {
	var customer model.Users
	err := cst.db.QueryRow(`SELECT id, full_name, phone_number, username, password, role, date_created FROM mst_users WHERE id = $1`, id).Scan(
		&customer.Id,
		&customer.FullName,
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
	err := cst.db.QueryRow(`SELECT id, full_name, phone_number, username, password, role, date_created FROM mst_users WHERE username = $1`, username).Scan(
		&customer.Id,
		&customer.FullName,
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
	err := cst.db.QueryRow(`INSERT INTO mst_users (full_name, phone_number, username, password, role, date_created) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, full_name, phone_number, username, password, role, date_created`,
		payload.FullName,
		payload.PhoneNumber,
		payload.Username,
		payload.Password,
		payload.Role,
		time.Now(),
	).Scan(
		&customer.Id,
		&customer.FullName,
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

func (cst *usersRepository) Update(payload model.Users) (model.Users, error) {
	var customer model.Users
	err := cst.db.QueryRow(`UPDATE mst_users SET full_name = $1, phone_number = $2, username = $3, password = $4, role = $5 WHERE id = $6 RETURNING id, full_name, phone_number, username, password, role, date_created`,
		payload.FullName,
		payload.PhoneNumber,
		payload.Username,
		payload.Password,
		payload.Role,
		payload.Id,
	).Scan(
		&customer.Id,
		&customer.FullName,
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

func (cst *usersRepository) Delete(id string) (model.Users, error) {
	var customer model.Users
	err := cst.db.QueryRow(`DELETE FROM mst_users WHERE id = $1 RETURNING id, full_name, phone_number, username, password, role, date_created`, id).Scan(
		&customer.Id,
		&customer.FullName,
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
