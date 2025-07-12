package mappers

import (
	"icomphub-api/dtos"
	"icomphub-api/models"
)

func TechnologyToDTO(technology *models.Technology) *dtos.TechnologyDTO {
	return &dtos.TechnologyDTO{
		Id:       technology.ID,
		Slug:     technology.Slug,
		Name:     technology.Name,
		HasImage: technology.ImageID != nil,
	}
}

func CreateRequestDTOToTechnology(createDTO *dtos.TechnologyCreateRequestDTO) *models.Technology {
	return &models.Technology{
		Slug: createDTO.Slug,
		Name: createDTO.Name,
	}
}

func TechnologiesToDTOs(technologies []models.Technology) []dtos.TechnologyDTO {
	dtosList := make([]dtos.TechnologyDTO, len(technologies))
	for i, technology := range technologies {
		dtosList[i] = *TechnologyToDTO(&technology)
	}
	return dtosList
}
