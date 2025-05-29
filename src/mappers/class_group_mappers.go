package mappers

import (
	"icomphub-api/dtos"
	"icomphub-api/enums"
	"icomphub-api/models"
)

func ClassGroupToDTO(classGroup *models.ClassGroup) *dtos.ClassGroupDTO {
	return &dtos.ClassGroupDTO{
		Id:     classGroup.ID,
		Slug:   classGroup.Slug,
		Name:   classGroup.Name,
		Status: string(classGroup.Status),
	}
}

func CreateRequestDTOToClassGroup(createDTO *dtos.ClassGroupCreateRequestDTO) *models.ClassGroup {
	return &models.ClassGroup{
		Slug:   createDTO.Slug,
		Name:   createDTO.Name,
		Status: enums.StatusWaitingApproval, // default status
	}
}

func ClassGroupsToDTOs(groups []models.ClassGroup) []dtos.ClassGroupDTO {
	result := make([]dtos.ClassGroupDTO, len(groups))
	for i, g := range groups {
		result[i] = *ClassGroupToDTO(&g)
	}
	return result
}
