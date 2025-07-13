package mappers

import (
	"encoding/json"

	"icomphub-api/dtos"
	"icomphub-api/enums"
	"icomphub-api/models"

	"gorm.io/datatypes"
)

func ProjectToDTO(project *models.Project) (*dtos.ProjectDTO, error) {
	var data map[string]any

	err := json.Unmarshal(project.Data, &data)
	if err != nil {
		return nil, err
	}

	var technologiesDTO []dtos.TechnologyDTO
	if project.Technologies != nil {
		for _, techModel := range project.Technologies {
			technologiesDTO = append(technologiesDTO, *TechnologyToDTO(&techModel))
		}
	}

	var imagesDTO []dtos.ProjectImageDTO
	if project.ProjectImages != nil {
		for _, imageModel := range project.ProjectImages {
			imagesDTO = append(imagesDTO, *ProjectImageToDTO(&imageModel))
		}
	}

	return &dtos.ProjectDTO{
		Id:           project.ID,
		Slug:         project.Slug,
		Name:         project.Name,
		Status:       string(project.Status),
		ThumbnailID:  project.ThumbnailID,
		Data:         data,
		ClassGroupId: project.ClassGroupID,
		Technologies: technologiesDTO,
		ImageIDs:     imagesDTO,
		ClassGroup:   *ClassGroupToDTO(&project.ClassGroup),
		Members:      MembersToShortDTOs(project.Members),
	}, nil
}

func ProjectToShortDTO(project *models.Project) *dtos.ProjectShortDTO {
	return &dtos.ProjectShortDTO{
		Id:           project.ID,
		Slug:         project.Slug,
		Name:         project.Name,
		Status:       string(project.Status),
		ThumbnailID:  project.ThumbnailID,
		ClassGroupId: project.ClassGroupID,
	}
}

func CreateRequestDTOToProject(createDTO *dtos.ProjectCreateRequestDTO) (*models.Project, error) {
	jsonBytes, err := json.Marshal(createDTO.Data)
	if err != nil {
		return nil, err
	}

	return &models.Project{
		Slug:         createDTO.Slug,
		Name:         createDTO.Name,
		Data:         datatypes.JSON(jsonBytes),
		ClassGroupID: createDTO.ClassGroupId,
		Status:       enums.StatusWaitingApproval,
	}, nil
}

func ProjectsToDTOs(projects []models.Project) ([]dtos.ProjectDTO, error) {
	result := make([]dtos.ProjectDTO, len(projects))
	for i, p := range projects {
		dto, err := ProjectToDTO(&p)
		if err != nil {
			return nil, err
		}
		result[i] = *dto
	}
	return result, nil
}
