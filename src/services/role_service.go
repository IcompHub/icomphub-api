package services

import (
	"icomphub-api/codes"
	"icomphub-api/dtos"
	"icomphub-api/mappers"
	"icomphub-api/models"
	"icomphub-api/repositories"
)

type RoleService interface {
	GetAll() ([]dtos.RoleDTO, codes.Code, error)
	GetByID(id uint64) (*dtos.RoleDTO, codes.Code, error)
	Create(dto *dtos.RoleCreateRequestDTO) (*dtos.RoleDTO, codes.Code, error)
	Update(id uint64, dto *dtos.RoleUpdateRequestDTO) (*dtos.RoleDTO, codes.Code, error)
	Delete(id uint64) (codes.Code, error)
}

type roleService struct {
	repository repositories.RoleRepository
}

func NewRoleService(repo repositories.RoleRepository) RoleService {
	return &roleService{repository: repo}
}

func (s *roleService) GetAll() ([]dtos.RoleDTO, codes.Code, error) {
	roles, err := s.repository.FindAll()
	if err != nil {
		return nil, codes.ErrorGettingAllRoles, err
	}
	return mappers.RolesToDTOs(roles), codes.GetAllRoles, nil
}

func (s *roleService) GetByID(id uint64) (*dtos.RoleDTO, codes.Code, error) {
	role, err := s.repository.FindByID(id)
	if err != nil {
		return nil, codes.ErrorFindingRole, err
	}
	return mappers.RoleToDTO(role), codes.FindRole, nil
}

func (s *roleService) Create(dto *dtos.RoleCreateRequestDTO) (*dtos.RoleDTO, codes.Code, error) {
	role := &models.Role{
		Slug: dto.Slug,
		Name: dto.Name,
	}
	err := s.repository.Create(role)
	if err != nil {
		return nil, codes.ErrorCreatingRole, err
	}
	return mappers.RoleToDTO(role), codes.CreateRole, nil
}

func (s *roleService) Update(id uint64, dto *dtos.RoleUpdateRequestDTO) (*dtos.RoleDTO, codes.Code, error) {
	role, err := s.repository.FindByID(id)
	if err != nil {
		return nil, codes.ErrorFindingRole, err
	}

	if dto.Name != nil {
		role.Name = *dto.Name
	}
	if dto.Slug != nil {
		role.Slug = *dto.Slug
	}

	err = s.repository.Update(role)
	if err != nil {
		return nil, codes.ErrorUpdatingRole, err
	}

	return mappers.RoleToDTO(role), codes.UpdateRole, nil
}

func (s *roleService) Delete(id uint64) (codes.Code, error) {
	role, err := s.repository.FindByID(id)
	if err != nil {
		return codes.ErrorFindingRole, err
	}

	err = s.repository.Delete(role)
	if err != nil {
		return codes.ErrorDeletingRole, err
	}

	return codes.DeleteRole, nil
}
