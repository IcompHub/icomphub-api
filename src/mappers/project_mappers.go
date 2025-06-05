package mappers

import (
	"encoding/json"

	"icomphub-api/dtos"
	"icomphub-api/enums"
	"icomphub-api/models"

	"gorm.io/datatypes"
)

func TechnologyToDTO2(tech *models.Technology) dtos.TechnologyDTO {
	return dtos.TechnologyDTO{
		Id:   tech.ID,
		Slug: tech.Slug,
		Name: tech.Name,
	}
}

func ProjectToDTO(project *models.Project) (*dtos.ProjectDTO, error) {
	var data map[string]any

	err := json.Unmarshal(project.Data, &data)
	if err != nil {
		return nil, err
	}

	var technologiesDTO []dtos.TechnologyDTO
	if project.Technologies != nil {
		for _, techModel := range project.Technologies {
			technologiesDTO = append(technologiesDTO, TechnologyToDTO2(&techModel))
		}
	}

	return &dtos.ProjectDTO{
		Id:           project.ID,
		Slug:         project.Slug,
		Name:         project.Name,
		Status:       string(project.Status),
		Data:         data,
		ClassGroupId: project.ClassGroupID,
		Technologies: technologiesDTO,
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
