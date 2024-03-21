package usecase

import (
	"enigma_laundry_api/model"
	"enigma_laundry_api/repository"
	"fmt"
)

type ServicesUseCase interface {
	FindById(id string) (model.Services, error)
	CreateServices(payload model.Services) (model.Services, error)
}

type servicesUseCase struct {
	repo repository.ServicesRepository
}

func (serv *servicesUseCase) FindById(id string) (model.Services, error) {
	services, err := serv.repo.Get(id)
	if err != nil {
		return model.Services{}, fmt.Errorf("service with ID %s not found", id)
	}
	return services, nil
}

func (serv *servicesUseCase) CreateServices(payload model.Services) (model.Services, error) {
	newServices := model.Services{
		Id:          payload.Id,
		ServiceName: payload.ServiceName,
		Unit:        payload.Unit,
		Price:       payload.Price,
	}

	services, err := serv.repo.Create(newServices)
	if err != nil {
		return model.Services{}, fmt.Errorf("failed to create services: %s", err.Error())
	}
	return services, nil
}

func NewServicesUseCase(repo repository.ServicesRepository) ServicesUseCase {
	return &servicesUseCase{
		repo: repo,
	}
}
