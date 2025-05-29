package services

import (
	"errors"
	"strings"

	"icomphub-api/codes"
	"icomphub-api/dtos"
	"icomphub-api/mappers"
	"icomphub-api/models"
	"icomphub-api/repositories"
)

type ClassGroupService interface {
	GetAll(req *dtos.ClassGroupRequestDTO) ([]dtos.ClassGroupDTO, codes.Code, error)
	CountAll(req *dtos.ClassGroupRequestDTO) (uint64, codes.Code, error)
	Find(id uint64) (*dtos.ClassGroupDTO, codes.Code, error)
	Create(createDTO *dtos.ClassGroupCreateRequestDTO) (*dtos.ClassGroupDTO, codes.Code, error)
	Delete(id uint64) (codes.Code, error)
	Update(id uint64, updateDTO *dtos.ClassGroupUpdateRequestDTO) (*dtos.ClassGroupDTO, codes.Code, error)
}

type classGroupService struct {
	repository repositories.ClassGroupRepository
}

func NewClassGroupService(repo repositories.ClassGroupRepository) ClassGroupService {
	return &classGroupService{repository: repo}
}

func (service *classGroupService) find(id uint64) (*models.ClassGroup, codes.Code, error) {
	classGroup, err := service.repository.Find(id)
	if err != nil {
		return nil, codes.ErrorFindingClassGroup, err
	}

	return classGroup, codes.FindClassGroup, nil
}

func (service *classGroupService) validateClassGroup(classGroup *models.ClassGroup, isCreating bool) error {
	if strings.TrimSpace(classGroup.Name) == "" {
		return errors.New("invalid class group name")
	}

	return nil
}

func (service *classGroupService) GetAll(req *dtos.ClassGroupRequestDTO) ([]dtos.ClassGroupDTO, codes.Code, error) {
	classGroups, err := service.repository.GetAll(req)
	if err != nil {
		return nil, codes.ErrorGettingAllClassGroups, err
	}

	return mappers.ClassGroupsToDTOs(classGroups), codes.GetAllClassGroups, nil
}

func (service *classGroupService) CountAll(req *dtos.ClassGroupRequestDTO) (uint64, codes.Code, error) {
	count, err := service.repository.CountAll(req)
	if err != nil {
		return 0, codes.ErrorCountingAllClassGroups, err
	}

	return count, codes.CountAllClassGroups, nil
}

func (service *classGroupService) Find(id uint64) (*dtos.ClassGroupDTO, codes.Code, error) {
	classGroup, code, err := service.find(id)
	return mappers.ClassGroupToDTO(classGroup), code, err
}

func (service *classGroupService) Create(createDTO *dtos.ClassGroupCreateRequestDTO) (*dtos.ClassGroupDTO, codes.Code, error) {
	classGroup := mappers.CreateRequestDTOToClassGroup(createDTO)

	err := service.validateClassGroup(classGroup, true)
	if err != nil {
		return nil, codes.InvalidParams, err
	}

	err = service.repository.Create(classGroup)
	if err != nil {
		return nil, codes.ErrorCreatingClassGroup, err
	}

	return mappers.ClassGroupToDTO(classGroup), codes.CreateClassGroup, nil
}

func (service *classGroupService) Delete(id uint64) (codes.Code, error) {
	classGroup, code, err := service.find(id)
	if err != nil {
		return code, err
	}

	err = service.repository.Delete(classGroup)
	if err != nil {
		return codes.ErrorDeletingClassGroup, err
	}

	return codes.DeleteClassGroup, nil
}

func (service *classGroupService) Update(id uint64, updateDTO *dtos.ClassGroupUpdateRequestDTO) (*dtos.ClassGroupDTO, codes.Code, error) {
	classGroup, code, err := service.find(id)
	if err != nil {
		return nil, code, err
	}

	if updateDTO.Name != nil && strings.TrimSpace(*updateDTO.Name) != "" {
		classGroup.Name = *updateDTO.Name
	}

	err = service.validateClassGroup(classGroup, false)
	if err != nil {
		return nil, codes.InvalidParams, err
	}

	err = service.repository.Update(classGroup)
	if err != nil {
		return mappers.ClassGroupToDTO(classGroup), codes.ErrorUpdatingClassGroup, err
	}

	return mappers.ClassGroupToDTO(classGroup), codes.UpdateClassGroup, nil
}
