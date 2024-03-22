package repository

import (
	"database/sql"
	"enigma_laundry_api/model"
	"errors"
	"fmt"
	"time"
)

type TransactionRepository interface {
	Create(payload model.Transaction) (model.Transaction, error)
	Delete(id string) (model.Transaction, error)
}

type transactionRepository struct {
	db *sql.DB
}

func (t *transactionRepository) Create(payload model.Transaction) (model.Transaction, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return model.Transaction{}, err
	}
	var transaction model.Transaction
	now := time.Now()
	err = tx.QueryRow(`INSERT INTO tx_enigma_laundry
	(id_users, id_services, transaction_in, transaction_out, amount, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id, id_users, id_services, transaction_in, transaction_out, amount, created_at, updated_at`,
		payload.Users,
		payload.Services,
		payload.TransactionIn,
		payload.TransactionOut,
		payload.Amount,
		now.Unix(),
		now.Unix(),
	).Scan(
		&transaction.Id,
		&transaction.Users,
		&transaction.Services,
		&transaction.TransactionIn,
		&transaction.TransactionOut,
		&transaction.Amount,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)

	if err != nil {
		tx.Rollback()
		errMsg := fmt.Sprintf("Error ON DB : %v", err)
		return model.Transaction{}, errors.New(errMsg)
	}

	err = tx.Commit()
	if err != nil {
		return model.Transaction{}, err
	}

	return transaction, nil
}

func (t *transactionRepository) Delete(id string) (model.Transaction, error) {
	var transaction model.Transaction
	err := t.db.QueryRow(`DELETE FROM tx_enigma_laundry WHERE id = $1 RETURNING id, id_users, id_services, transaction_in, transaction_out, amount, created_at, updated_at`, id).Scan(
		&transaction.Id,
		&transaction.Users,
		&transaction.Services,
		&transaction.TransactionIn,
		&transaction.TransactionOut,
		&transaction.Amount,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)

	if err != nil {
		return model.Transaction{}, err
	}
	return transaction, nil
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{db: db}
}
