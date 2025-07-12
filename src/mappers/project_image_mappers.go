package mappers

import (
	"icomphub-api/dtos"
	"icomphub-api/models"
)

func ProjectImageToDTO(projectImage *models.ProjectImage) *dtos.ProjectImageDTO {
	return &dtos.ProjectImageDTO{
		Id:      projectImage.ID,
		ImageID: projectImage.ImageID,
	}
}

func ProjectImagesToDTOs(projectImages []models.ProjectImage) []dtos.ProjectImageDTO {
	result := make([]dtos.ProjectImageDTO, len(projectImages))
	for i, p := range projectImages {
		dto := ProjectImageToDTO(&p)
		result[i] = *dto
	}
	return result
}
