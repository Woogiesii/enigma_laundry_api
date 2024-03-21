package manager

import "enigma_laundry_api/repository"

type RepoManager interface {
	CustomersRepo() repository.UsersRepository
	ServicesRepo() repository.ServicesRepository
	TransactionRepo() repository.TransactionRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) CustomersRepo() repository.UsersRepository {
	return repository.NewUsersRepository(r.infra.Conn())
}

func (r *repoManager) ServicesRepo() repository.ServicesRepository {
	return repository.NewServicesRepository(r.infra.Conn())
}

func (r *repoManager) TransactionRepo() repository.TransactionRepository {
	return repository.NewTransactionRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
