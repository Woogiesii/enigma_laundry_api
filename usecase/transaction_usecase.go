package usecase

import (
	"enigma_laundry_api/model"
	"enigma_laundry_api/repository"
	"fmt"
)

type TransactionUseCase interface {
	RegisterTransaction(payload model.Transaction) (model.Transaction, error)
}

type transactionUseCase struct {
	repo repository.TransactionRepository
}

func (t *transactionUseCase) RegisterTransaction(payload model.Transaction) (model.Transaction, error) {
	newTransaction := model.Transaction{
		Id:             payload.Id,
		Users:          payload.Users,
		Services:       payload.Services,
		TransactionIn:  payload.TransactionIn,
		TransactionOut: payload.TransactionOut,
		Amount:         payload.Amount,
		CreatedAt:      payload.CreatedAt,
		UpdatedAt:      payload.UpdatedAt,
	}
	transaction, err := t.repo.Create(newTransaction)
	if err != nil {
		return model.Transaction{}, fmt.Errorf("failed to create services: %s", err.Error())
	}
	return transaction, nil
}

func NewTransactionUseCase(repo repository.TransactionRepository) TransactionUseCase {
	return &transactionUseCase{
		repo: repo,
	}
}
