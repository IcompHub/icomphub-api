package services

import (
	"errors"
	"mime/multipart"
	"strings"

	"icomphub-api/auth"
	"icomphub-api/codes"
	"icomphub-api/dtos"
	"icomphub-api/files"
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
	UpdateProfilePicture(id uint64, image *multipart.FileHeader) (*dtos.UserDTO, codes.Code, error)
	DeleteProfilePicture(id uint64) (*dtos.UserDTO, codes.Code, error)
	GetProfilePictureFullPath(id uint64) (string, codes.Code, error)
}

type userService struct {
	repository  repositories.UserRepository
	fileService files.FileUploadService
}

func NewUserService(repo repositories.UserRepository, fileService files.FileUploadService) UserService {
	return &userService{repository: repo, fileService: fileService}
}

func (service *userService) find(id uint64) (*models.User, codes.Code, error) {
	user, err := service.repository.Find(id)
	if err != nil {
		return nil, codes.ErrorFindingUser, err
	}

	return user, codes.FindUser, nil
}

func (service *userService) validateUser(user *models.User, isCreating bool) error {
	if strings.TrimSpace(user.Slug) == "" {
		return errors.New("invalid user slug")
	}

	if strings.TrimSpace(user.Nickname) == "" {
		return errors.New("invalid user nickname")
	}

	if strings.TrimSpace(user.FullName) == "" {
		return errors.New("invalid user full name")
	}

	if strings.TrimSpace(user.PersonalEmail) == "" {
		return errors.New("invalid user personal email")
	}

	if strings.TrimSpace(user.Password) == "" {
		return errors.New("invalid user password")
	}

	if !isCreating {
		if strings.TrimSpace(string(user.Status)) == "" {
			return errors.New("invalid user status")
		}

		if strings.TrimSpace(string(user.SystemRole)) == "" {
			return errors.New("invalid user systemRole")
		}
	}

	if user.InstitutionalEmail != nil && strings.TrimSpace(*user.InstitutionalEmail) == "" {
		return errors.New("invalid user institutional email")
	}

	if user.Registration != nil && strings.TrimSpace(*user.Registration) == "" {
		return errors.New("invalid user registration")
	}

	return nil
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

	hashedPassword, err := auth.HashPassword(createDTO.Password)
	if err != nil {
		return nil, codes.ErrorCreatingUser, err
	}
	user.Password = hashedPassword

	err = service.validateUser(user, true)
	if err != nil {
		return nil, codes.InvalidParams, err
	}

	err = service.repository.Create(user)
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

	if updateDTO.Slug != nil {
		user.Slug = *updateDTO.Slug
	}

	if updateDTO.FullName != nil {
		user.FullName = *updateDTO.FullName
	}

	if updateDTO.Nickname != nil {
		user.Nickname = *updateDTO.Nickname
	}

	if updateDTO.PersonalEmail != nil {
		user.PersonalEmail = *updateDTO.PersonalEmail
	}

	user.Registration = updateDTO.Registration

	user.InstitutionalEmail = updateDTO.InstitutionalEmail

	err = service.validateUser(user, false)
	if err != nil {
		return nil, codes.InvalidParams, err
	}

	err = service.repository.Update(user)
	if err != nil {
		return mappers.UserToDTO(user), codes.ErrorUpdatingUser, err
	}

	return mappers.UserToDTO(user), codes.UpdateUser, nil
}

func (service *userService) UpdateProfilePicture(id uint64, image *multipart.FileHeader) (*dtos.UserDTO, codes.Code, error) {
	user, code, err := service.find(id)
	if err != nil {
		return nil, code, err
	}

	path, code, err := service.fileService.SaveFile(image, files.UploadConfig{
		TargetFolder: "users",
		AllowedTypes: []string{"image/jpeg", "image/png"},
		MaxSizeMB:    5,
	})
	if err != nil {
		return nil, code, err
	}

	if user.ProfilePictureId != nil {
		_, code, err = service.DeleteProfilePicture(id)
		if err != nil {
			return nil, code, err
		}

		user.ProfilePictureId = nil
	}

	user.ProfilePictureId = &path

	err = service.validateUser(user, false)
	if err != nil {
		return nil, codes.InvalidParams, err
	}

	err = service.repository.Update(user)
	if err != nil {
		return mappers.UserToDTO(user), codes.ErrorUpdatingUser, err
	}

	return mappers.UserToDTO(user), codes.UpdateUser, nil
}

func (service *userService) DeleteProfilePicture(id uint64) (*dtos.UserDTO, codes.Code, error) {
	user, code, err := service.find(id)
	if err != nil {
		return nil, code, err
	}

	if user.ProfilePictureId == nil {
		return mappers.UserToDTO(user), codes.UpdateUser, nil
	}

	code, err = service.fileService.DeleteFile(*user.ProfilePictureId, files.UploadConfig{
		TargetFolder: "users",
		AllowedTypes: []string{"image/jpeg", "image/png"},
		MaxSizeMB:    5,
	})
	if err != nil {
		return nil, code, err
	}

	user.ProfilePictureId = nil

	err = service.validateUser(user, false)
	if err != nil {
		return nil, codes.InvalidParams, err
	}

	err = service.repository.Update(user)
	if err != nil {
		return mappers.UserToDTO(user), codes.ErrorUpdatingUser, err
	}

	return mappers.UserToDTO(user), codes.UpdateUser, nil
}

func (service *userService) GetProfilePictureFullPath(id uint64) (string, codes.Code, error) {
	user, code, err := service.find(id)
	if err != nil {
		return "", code, err
	}

	if user.ProfilePictureId == nil {
		return "", codes.FileNotFound, errors.New("no profile picture registered")
	}

	fullPath, code, err := service.fileService.GetFile(*user.ProfilePictureId, files.UploadConfig{
		TargetFolder: "users",
		AllowedTypes: []string{"image/jpeg", "image/png"},
		MaxSizeMB:    5,
	})
	if err != nil {
		return "", code, err
	}

	return fullPath, codes.FileFound, nil
}
