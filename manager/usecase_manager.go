package manager

import "enigma_laundry_api/usecase"

type UseCaseManager interface {
	UsersUseCase() usecase.UsersUseCase
	ServicesUseCase() usecase.ServicesUseCase
	TransactionUseCase() usecase.TransactionUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) UsersUseCase() usecase.UsersUseCase {
	return usecase.NewUsersUseCase(u.repo.UsersRepo())
}

func (u *useCaseManager) ServicesUseCase() usecase.ServicesUseCase {
	return usecase.NewServicesUseCase(u.repo.ServicesRepo())
}

func (u *useCaseManager) TransactionUseCase() usecase.TransactionUseCase {
	return usecase.NewTransactionUseCase(u.repo.TransactionRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
