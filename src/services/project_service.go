package services

import (
	"encoding/json"
	"errors"
	"mime/multipart"
	"strings"

	"icomphub-api/codes"
	"icomphub-api/dtos"
	"icomphub-api/files"
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
	UpdateThumbnail(id uint64, image *multipart.FileHeader) (*dtos.ProjectDTO, codes.Code, error)
	DeleteThumbnail(id uint64) (*dtos.ProjectDTO, codes.Code, error)
	GetThumbnailFullPath(id uint64) (string, codes.Code, error)
	CreateImage(id uint64, image *multipart.FileHeader) (*dtos.ProjectImageDTO, codes.Code, error)
	UpdateImage(imageID uint64, image *multipart.FileHeader) (*dtos.ProjectImageDTO, codes.Code, error)
	DeleteImage(imageID uint64) (codes.Code, error)
	GetImageFullPath(imageID uint64) (string, codes.Code, error)
}

type projectService struct {
	repository        repositories.ProjectRepository
	fileService       files.FileUploadService
	imageRepository   repositories.ProjectImageRepository
	technologyService TechnologyService
	memberService     MemberService
}

func NewProjectService(repo repositories.ProjectRepository, fileService files.FileUploadService, imageRepository repositories.ProjectImageRepository, technologyService TechnologyService, memberService MemberService) ProjectService {
	return &projectService{repository: repo, fileService: fileService, imageRepository: imageRepository, technologyService: technologyService, memberService: memberService}
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

	var validTechIds []uint64

	for _, techId := range createDTO.TechnologyIDs {
		tech, code, err := service.technologyService.Find(uint64(techId))
		if err != nil {
			return nil, code, err
		}

		validTechIds = append(validTechIds, tech.Id)
	}

	err = service.repository.ReplaceTechnologies(project.ID, validTechIds)
	if err != nil {
		return nil, codes.ErrorCreatingProject, err
	}

	for _, memberCreateDTO := range createDTO.Members {
		memberCreateDTO.ProjectID = project.ID
		_, code, err := service.memberService.Create(&memberCreateDTO)
		if err != nil {
			return nil, code, err
		}
	}

	project, code, err := service.find(project.ID)
	if err != nil {
		return nil, code, err
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

	if updateDTO.ClassGroupId != nil && *updateDTO.ClassGroupId != 0 {
		project.ClassGroupID = *updateDTO.ClassGroupId
	}

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

	var validTechIds []uint64

	for _, techId := range updateDTO.TechnologyIDs {
		tech, code, err := service.technologyService.Find(uint64(techId))
		if err != nil {
			return nil, code, err
		}

		validTechIds = append(validTechIds, tech.Id)
	}

	err = service.repository.ReplaceTechnologies(project.ID, validTechIds)
	if err != nil {
		return nil, codes.ErrorCreatingProject, err
	}

	updatedMemberIDs := make(map[uint64]bool)

	for _, memberUpdateDTO := range updateDTO.Members {
		if memberUpdateDTO.ID != nil {
			updatedMemberIDs[*memberUpdateDTO.ID] = true

			_, code, err := service.memberService.Update(*memberUpdateDTO.ID, &dtos.MemberUpdateRequestDTO{
				Nickname:  memberUpdateDTO.Nickname,
				ProjectID: &project.ID,
				UserId:    memberUpdateDTO.UserId,
				RoleIDs:   memberUpdateDTO.RoleIDs,
			})
			if err != nil {
				return nil, code, err
			}
		} else {
			_, code, err := service.memberService.Create(&dtos.MemberCreateRequestDTO{
				Nickname:  *memberUpdateDTO.Nickname,
				ProjectID: project.ID,
				UserId:    memberUpdateDTO.UserId,
				RoleIDs:   memberUpdateDTO.RoleIDs,
			})
			if err != nil {
				return nil, code, err
			}
		}
	}

	for _, member := range project.Members {
		if _, found := updatedMemberIDs[member.ID]; !found {
			code, err := service.memberService.Delete(member.ID)
			if err != nil {
				return nil, code, err
			}
		}
	}

	project, code, err = service.find(project.ID)
	if err != nil {
		return nil, code, err
	}

	mapper, err := mappers.ProjectToDTO(project)
	if err != nil {
		return nil, codes.ErrorCreatingProject, err
	}

	return mapper, codes.UpdateProject, nil
}

func (service *projectService) UpdateThumbnail(id uint64, image *multipart.FileHeader) (*dtos.ProjectDTO, codes.Code, error) {
	project, code, err := service.find(id)
	if err != nil {
		return nil, code, err
	}

	path, code, err := service.fileService.SaveFile(image, files.UploadConfig{
		TargetFolder: "projects",
		AllowedTypes: []string{"image/jpeg", "image/png"},
		MaxSizeMB:    5,
	})
	if err != nil {
		return nil, code, err
	}

	if project.ThumbnailID != nil {
		_, code, err = service.DeleteThumbnail(id)
		if err != nil {
			return nil, code, err
		}

		project.ThumbnailID = nil
	}

	project.ThumbnailID = &path

	err = service.validateProject(project)
	if err != nil {
		return nil, codes.InvalidParams, err
	}

	err = service.repository.Update(project)
	if err != nil {
		return nil, codes.ErrorUpdatingProject, err
	}

	projectDTO, err := mappers.ProjectToDTO(project)
	if err != nil {
		return nil, codes.ErrorUpdatingProject, err
	}

	return projectDTO, codes.UpdateProject, nil
}

func (service *projectService) DeleteThumbnail(id uint64) (*dtos.ProjectDTO, codes.Code, error) {
	project, code, err := service.find(id)
	if err != nil {
		return nil, code, err
	}

	if project.ThumbnailID == nil {
		projectDTO, err := mappers.ProjectToDTO(project)
		if err != nil {
			return nil, codes.ErrorUpdatingProject, err
		}

		return projectDTO, codes.UpdateProject, nil
	}

	code, err = service.fileService.DeleteFile(*project.ThumbnailID, files.UploadConfig{
		TargetFolder: "projects",
		AllowedTypes: []string{"image/jpeg", "image/png"},
		MaxSizeMB:    5,
	})
	if err != nil {
		return nil, code, err
	}

	project.ThumbnailID = nil

	err = service.validateProject(project)
	if err != nil {
		return nil, codes.InvalidParams, err
	}

	err = service.repository.Update(project)
	if err != nil {
		return nil, codes.ErrorUpdatingProject, err
	}

	projectDTO, err := mappers.ProjectToDTO(project)
	if err != nil {
		return nil, codes.ErrorUpdatingProject, err
	}

	return projectDTO, codes.UpdateProject, nil
}

func (service *projectService) GetThumbnailFullPath(id uint64) (string, codes.Code, error) {
	project, code, err := service.find(id)
	if err != nil {
		return "", code, err
	}

	if project.ThumbnailID == nil {
		return "", codes.FileNotFound, errors.New("no profile picture registered")
	}

	fullPath, code, err := service.fileService.GetFile(*project.ThumbnailID, files.UploadConfig{
		TargetFolder: "projects",
		AllowedTypes: []string{"image/jpeg", "image/png"},
		MaxSizeMB:    5,
	})
	if err != nil {
		return "", code, err
	}

	return fullPath, codes.FileFound, nil
}

func (service *projectService) CreateImage(id uint64, image *multipart.FileHeader) (*dtos.ProjectImageDTO, codes.Code, error) {
	project, code, err := service.find(id)
	if err != nil {
		return nil, code, err
	}

	filename, code, err := service.fileService.SaveFile(image, files.UploadConfig{
		TargetFolder: "projects",
		AllowedTypes: []string{"image/jpeg", "image/png"},
		MaxSizeMB:    5,
	})
	if err != nil {
		return nil, code, err
	}

	projectImage := &models.ProjectImage{
		ProjectID: project.ID,
		ImageID:   filename,
	}

	err = service.imageRepository.Create(projectImage)
	if err != nil {
		return nil, codes.ErrorFindingProject, err
	}

	return mappers.ProjectImageToDTO(projectImage), codes.CreateProject, nil
}

func (service *projectService) UpdateImage(imageID uint64, image *multipart.FileHeader) (*dtos.ProjectImageDTO, codes.Code, error) {
	projectImage, err := service.imageRepository.Find(imageID)
	if err != nil {
		return nil, codes.FileNotFound, err
	}

	code, err := service.fileService.DeleteFile(projectImage.ImageID, files.UploadConfig{
		TargetFolder: "projects",
		AllowedTypes: []string{"image/jpeg", "image/png"},
		MaxSizeMB:    5,
	})
	if err != nil {
		return nil, code, err
	}

	projectImage.ImageID, code, err = service.fileService.SaveFile(image, files.UploadConfig{
		TargetFolder: "projects",
		AllowedTypes: []string{"image/jpeg", "image/png"},
		MaxSizeMB:    5,
	})
	if err != nil {
		return nil, code, err
	}

	err = service.imageRepository.Update(projectImage)
	if err != nil {
		return nil, codes.ErrorFindingProject, err
	}

	return mappers.ProjectImageToDTO(projectImage), codes.CreateProject, nil
}

func (service *projectService) DeleteImage(imageID uint64) (codes.Code, error) {
	projectImage, err := service.imageRepository.Find(imageID)
	if err != nil {
		return codes.FileNotFound, err
	}

	err = service.imageRepository.Delete(projectImage)
	if err != nil {
		return codes.UnknowError, err
	}

	code, err := service.fileService.DeleteFile(projectImage.ImageID, files.UploadConfig{
		TargetFolder: "projects",
		AllowedTypes: []string{"image/jpeg", "image/png"},
		MaxSizeMB:    5,
	})
	if err != nil {
		return code, err
	}

	return codes.DeleteProject, nil
}

func (service *projectService) GetImageFullPath(imageID uint64) (string, codes.Code, error) {
	projectImage, err := service.imageRepository.Find(imageID)
	if err != nil {
		return "", codes.FileNotFound, err
	}

	fullPath, code, err := service.fileService.GetFile(projectImage.ImageID, files.UploadConfig{
		TargetFolder: "projects",
		AllowedTypes: []string{"image/jpeg", "image/png"},
		MaxSizeMB:    5,
	})
	if err != nil {
		return "", code, err
	}

	return fullPath, codes.FileFound, nil
}
