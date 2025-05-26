package services

import (
	"icomphub-api/codes"
	"icomphub-api/dtos"
	"icomphub-api/mappers"
	"icomphub-api/models"
	"icomphub-api/repositories"
)

type UserService interface {
	GetAll(req *dtos.UserRequestDTO) ([]dtos.UserDTO, codes.Code, error)
	CountAll(req *dtos.UserRequestDTO) (uint64, codes.Code, error)
	Find(id uint64) (*dtos.UserDTO, codes.Code, error)
	Create(createDTO *dtos.UserCreateRequestDTO) (*dtos.UserDTO, codes.Code, error)
	Delete(id uint64) (codes.Code, error)
	Update(id uint64, updateDTO *dtos.UserUpdateRequestDTO) (*dtos.UserDTO, codes.Code, error)
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repository: repo}
}

func (service *userService) find(id uint64) (*models.User, codes.Code, error) {
	user, err := service.repository.Find(id)
	if err != nil {
		return nil, codes.ErrorFindingUser, err
	}

	return user, codes.FindUser, nil
}

func (service *userService) GetAll(req *dtos.UserRequestDTO) ([]dtos.UserDTO, codes.Code, error) {
	users, err := service.repository.GetAll(req)
	if err != nil {
		return nil, codes.ErrorGettingAllUsers, err
	}

	return mappers.UsersToDTOs(users), codes.GetAllUsers, nil
}

func (service *userService) CountAll(req *dtos.UserRequestDTO) (uint64, codes.Code, error) {
	count, err := service.repository.CountAll(req)
	if err != nil {
		return 0, codes.ErrorCoutingAllUsers, err
	}

	return count, codes.CountAllUsers, nil
}

func (service *userService) Find(id uint64) (*dtos.UserDTO, codes.Code, error) {
	user, code, err := service.find(id)

	return mappers.UserToDTO(user), code, err
}

func (service *userService) Create(createDTO *dtos.UserCreateRequestDTO) (*dtos.UserDTO, codes.Code, error) {
	user := mappers.CreateRequestDTOToUser(createDTO)

	err := service.repository.Create(user)
	if err != nil {
		return nil, codes.ErrorCreatingUser, err
	}

	return mappers.UserToDTO(user), codes.CreateUser, nil
}

func (service *userService) Delete(id uint64) (codes.Code, error) {
	user, code, err := service.find(id)
	if err != nil {
		return code, err
	}

	err = service.repository.Delete(user)
	if err != nil {
		return codes.ErrorDeletingUser, err
	}

	return codes.DeleteUser, nil
}

func (service *userService) Update(id uint64, updateDTO *dtos.UserUpdateRequestDTO) (*dtos.UserDTO, codes.Code, error) {
	user, code, err := service.find(id)
	if err != nil {
		return nil, code, err
	}

	if updateDTO.Slug != "" {
		user.Slug = updateDTO.Slug
	}

	if updateDTO.FullName != "" {
		user.FullName = updateDTO.FullName
	}

	err = service.repository.Update(user)
	if err != nil {
		return mappers.UserToDTO(user), codes.ErrorUpdatingUser, err
	}

	return mappers.UserToDTO(user), codes.UpdateUser, nil
}
