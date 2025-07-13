package mappers

import (
	"icomphub-api/dtos"
	"icomphub-api/models"
)

func RoleToDTO(Role *models.Role) *dtos.RoleDTO {
	return &dtos.RoleDTO{
		Id:     Role.ID,
		Slug:   Role.Slug,
		Name:   Role.Name,
		Status: string(Role.Status),
	}
}

func CreateRequestDTOToRole(createDTO *dtos.RoleCreateRequestDTO) *models.Role {
	return &models.Role{
		Slug: createDTO.Slug,
		Name: createDTO.Name,
	}
}

func RoleToDTOs(role []models.Role) []dtos.RoleDTO {
	dtosList := make([]dtos.RoleDTO, len(role))
	for i, Role := range role {
		dtosList[i] = *RoleToDTO(&Role)
	}
	return dtosList
}
