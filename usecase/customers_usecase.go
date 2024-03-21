package usecase

import (
	"enigma_laundry_api/model"
	"enigma_laundry_api/model/dto"
	"enigma_laundry_api/repository"
	"enigma_laundry_api/utils/common"
	"enigma_laundry_api/utils/encryption"
	"errors"
	"fmt"
	"time"
)

type UsersUseCase interface {
	FindById(id string) (model.Users, error)
	LoginCustomer(in dto.LoginRequestDto) (dto.LoginResponseDto, error)
	CreateCustomer(payload dto.UsersRequestDto) (model.Users, error)
}

type usersUseCase struct {
	repo repository.UsersRepository
}

func (cst *usersUseCase) FindById(id string) (model.Users, error) {
	customer, err := cst.repo.Get(id)
	if err != nil {
		return model.Users{}, fmt.Errorf("customer with ID %s not found", id)
	}
	return customer, nil
}

func (cst *usersUseCase) LoginCustomer(in dto.LoginRequestDto) (dto.LoginResponseDto, error) {
	customerData, err := cst.repo.GetByUsername(in.Username)
	if err != nil {
		return dto.LoginResponseDto{}, err
	}
	isValid := encryption.CheckPassword(in.Pass, customerData.Password)
	if !isValid {
		return dto.LoginResponseDto{}, errors.New("1")
	}

	loginExpDuration := time.Duration(10) * time.Minute
	expiresAt := time.Now().Add(loginExpDuration).Unix()
	accesToken, err := common.GenerateTokenJwt(customerData, expiresAt)
	if err != nil {
		return dto.LoginResponseDto{}, err
	}

	return dto.LoginResponseDto{
		AccesToken: accesToken,
		UserId:     customerData.Id,
	}, nil
}

func (cst *usersUseCase) CreateCustomer(payload dto.UsersRequestDto) (model.Users, error) {
	hashPassword, err := encryption.HashPassword(payload.Password)
	if err != nil {
		return model.Users{}, nil
	}
	newCustomers := model.Users{
		Id:           payload.Id,
		CustomerName: payload.CustomerName,
		PhoneNumber:  payload.PhoneNumber,
		Username:     payload.Username,
		Password:     hashPassword,
		Role:         payload.Role,
	}

	customers, err := cst.repo.Create(newCustomers)
	if err != nil {
		return model.Users{}, fmt.Errorf("failed to create customers: %s", err.Error())
	}

	return customers, nil
}

func NewUsersUseCase(repo repository.UsersRepository) UsersUseCase {
	return &usersUseCase{
		repo: repo,
	}
}
