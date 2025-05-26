package services

import (
	"icomphub-api/codes"
	"icomphub-api/dtos"
	"icomphub-api/mappers"
	"icomphub-api/repositories"
)

type UserService interface {
	GetAll(req *dtos.UserRequestDTO) ([]dtos.UserDTO, codes.Code, error)
	CountAll(req *dtos.UserRequestDTO) (uint64, codes.Code, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{userRepo: repo}
}

func (service *userService) GetAll(req *dtos.UserRequestDTO) ([]dtos.UserDTO, codes.Code, error) {
	users, err := service.userRepo.GetAll(req)
	if err != nil {
		return nil, codes.ErrorGettingAllUsers, err
	}

	return mappers.UsersToDTOs(users), codes.GetAllUsers, nil
}

func (service *userService) CountAll(req *dtos.UserRequestDTO) (uint64, codes.Code, error) {
	count, err := service.userRepo.CountAll(req)
	if err != nil {
		return 0, codes.ErrorCoutingAllUsers, err
	}

	return count, codes.CountAllUsers, nil
}
