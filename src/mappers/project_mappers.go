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

func ProjectToDetailDTO(project *models.Project) (*dtos.ProjectDetailDTO, error) {
	var data map[string]any
	if err := json.Unmarshal(project.Data, &data); err != nil {
		return nil, err
	}

	var technologiesDTO []dtos.TechnologyDTO
	if project.Technologies != nil {
		for _, techModel := range project.Technologies {
			technologiesDTO = append(technologiesDTO, *TechnologyToDTO(&techModel))
		}
	}

	var membersDTO []dtos.ProjectMemberDTO
	if project.Members != nil {
		for _, memberModel := range project.Members {
			memberDTO := dtos.ProjectMemberDTO{
				Id:       memberModel.ID,
				Nickname: memberModel.Nickname,
				Status:   string(memberModel.Status),
				User: dtos.MemberUserDTO{
					Id:       memberModel.User.ID,
					Slug:     memberModel.User.Slug,
					Nickname: memberModel.User.Nickname,
					FullName: memberModel.User.FullName,
				},
				Role: dtos.MemberRoleDTO{
					Id:   memberModel.Role.ID,
					Slug: memberModel.Role.Slug,
					Name: memberModel.Role.Name,
				},
			}
			membersDTO = append(membersDTO, memberDTO)
		}
	}

	return &dtos.ProjectDetailDTO{
		Id:           project.ID,
		Slug:         project.Slug,
		Name:         project.Name,
		Status:       string(project.Status),
		Data:         data,
		ClassGroupId: project.ClassGroupID,
		Technologies: technologiesDTO,
		Members:      membersDTO,
	}, nil
}
