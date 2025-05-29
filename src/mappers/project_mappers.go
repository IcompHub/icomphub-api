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

	return &dtos.ProjectDTO{
		Id:           project.ID,
		Slug:         project.Slug,
		Name:         project.Name,
		Status:       string(project.Status),
		Data:         data,
		ClassGroupId: project.ClassGroupID,
	}, nil
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
		Status:       enums.StatusWaitingApproval, // default status
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
