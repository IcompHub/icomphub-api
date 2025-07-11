package services

import (
	"encoding/json"
	"errors"
	"strings"

	"icomphub-api/codes"
	"icomphub-api/dtos"
	"icomphub-api/mappers"
	"icomphub-api/models"
	"icomphub-api/repositories"

	"gorm.io/datatypes"
)

type ProjectService interface {
	GetAll(req *dtos.ProjectRequestDTO) ([]dtos.ProjectDTO, codes.Code, error)
	CountAll(req *dtos.ProjectRequestDTO) (uint64, codes.Code, error)
	Find(id uint64) (*dtos.ProjectDTO, codes.Code, error)
	Create(createDTO *dtos.ProjectCreateRequestDTO) (*dtos.ProjectDTO, codes.Code, error)
	Delete(id uint64) (codes.Code, error)
	Update(id uint64, updateDTO *dtos.ProjectUpdateRequestDTO) (*dtos.ProjectDTO, codes.Code, error)
}

type projectService struct {
	repository repositories.ProjectRepository
}

func NewProjectService(repo repositories.ProjectRepository) ProjectService {
	return &projectService{repository: repo}
}

func (service *projectService) find(id uint64) (*models.Project, codes.Code, error) {
	project, err := service.repository.Find(id)
	if err != nil {
		return nil, codes.ErrorFindingProject, err
	}

	return project, codes.FindProject, nil
}

func (service *projectService) validateProject(project *models.Project) error {
	if strings.TrimSpace(project.Name) == "" {
		return errors.New("invalid project name")
	}

	if project.ClassGroupID == 0 {
		return errors.New("invalid class group ID")
	}

	if project.Data == nil {
		return errors.New("invalid project data")
	}

	return nil
}

func (service *projectService) GetAll(req *dtos.ProjectRequestDTO) ([]dtos.ProjectDTO, codes.Code, error) {
	projects, err := service.repository.GetAll(req)
	if err != nil {
		return nil, codes.ErrorGettingAllProjects, err
	}

	mapper, err := mappers.ProjectsToDTOs(projects)
	if err != nil {
		return nil, codes.ErrorGettingAllProjects, err
	}

	return mapper, codes.GetAllProjects, nil
}

func (service *projectService) CountAll(req *dtos.ProjectRequestDTO) (uint64, codes.Code, error) {
	count, err := service.repository.CountAll(req)
	if err != nil {
		return 0, codes.ErrorCountingAllProjects, err
	}

	return count, codes.CountAllProjects, nil
}

func (service *projectService) Find(id uint64) (*dtos.ProjectDTO, codes.Code, error) {
	project, code, err := service.find(id)
	if err != nil {
		return nil, code, err
	}

	mapper, err := mappers.ProjectToDTO(project)
	if err != nil {
		return nil, codes.ErrorFindingProject, err
	}

	return mapper, code, err
}

func (service *projectService) Create(createDTO *dtos.ProjectCreateRequestDTO) (*dtos.ProjectDTO, codes.Code, error) {
	project, err := mappers.CreateRequestDTOToProject(createDTO)
	if err != nil {
		return nil, codes.ErrorCreatingProject, err
	}

	err = service.validateProject(project)
	if err != nil {
		return nil, codes.InvalidParams, err
	}

	err = service.repository.Create(project)
	if err != nil {
		return nil, codes.ErrorCreatingProject, err
	}

	mapper, err := mappers.ProjectToDTO(project)
	if err != nil {
		return nil, codes.ErrorCreatingProject, err
	}

	return mapper, codes.CreateProject, nil
}

func (service *projectService) Delete(id uint64) (codes.Code, error) {
	project, code, err := service.find(id)
	if err != nil {
		return code, err
	}

	err = service.repository.Delete(project)
	if err != nil {
		return codes.ErrorDeletingProject, err
	}

	return codes.DeleteProject, nil
}

func (service *projectService) Update(id uint64, updateDTO *dtos.ProjectUpdateRequestDTO) (*dtos.ProjectDTO, codes.Code, error) {
	project, code, err := service.find(id)
	if err != nil {
		return nil, code, err
	}

	if updateDTO.Name != nil && strings.TrimSpace(*updateDTO.Name) != "" {
		project.Name = *updateDTO.Name
	}

	// if updateDTO.ClassGroupID != nil && *updateDTO.ClassGroupID != 0 {
	// 	project.ClassGroupID = *updateDTO.ClassGroupID
	// }

	if updateDTO.Data != nil {
		jsonBytes, err := json.Marshal(updateDTO.Data)
		if err != nil {
			panic(err)
		}
		project.Data = datatypes.JSON(jsonBytes)
	}

	err = service.validateProject(project)
	if err != nil {
		return nil, codes.InvalidParams, err
	}

	err = service.repository.Update(project)
	if err != nil {
		return nil, codes.ErrorUpdatingProject, err
	}

	mapper, err := mappers.ProjectToDTO(project)
	if err != nil {
		return nil, codes.ErrorCreatingProject, err
	}

	return mapper, codes.UpdateProject, nil
}
