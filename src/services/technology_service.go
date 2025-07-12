package services

import (
	"errors"
	"mime/multipart"
	"strings"

	"icomphub-api/codes"
	"icomphub-api/dtos"
	"icomphub-api/files"
	"icomphub-api/mappers"
	"icomphub-api/models"
	"icomphub-api/repositories"
)

type TechnologyService interface {
	GetAll(req *dtos.TechnologyRequestDTO) ([]dtos.TechnologyDTO, codes.Code, error)
	CountAll(req *dtos.TechnologyRequestDTO) (uint64, codes.Code, error)
	Find(id uint64) (*dtos.TechnologyDTO, codes.Code, error)
	Create(createDTO *dtos.TechnologyCreateRequestDTO) (*dtos.TechnologyDTO, codes.Code, error)
	Delete(id uint64) (codes.Code, error)
	Update(id uint64, updateDTO *dtos.TechnologyUpdateRequestDTO) (*dtos.TechnologyDTO, codes.Code, error)
	UpdateImage(id uint64, image *multipart.FileHeader) (*dtos.TechnologyDTO, codes.Code, error)
	DeleteImage(id uint64) (codes.Code, error)
	GetImageFullPath(id uint64) (string, codes.Code, error)
}

type technologyService struct {
	repository  repositories.TechnologyRepository
	fileService files.FileUploadService
}

func NewTechnologyService(repository repositories.TechnologyRepository, fileService files.FileUploadService) TechnologyService {
	return &technologyService{repository: repository, fileService: fileService}
}

func (service *technologyService) find(id uint64) (*models.Technology, codes.Code, error) {
	technology, err := service.repository.Find(id)
	if err != nil {
		return nil, codes.ErrorFindingTechnology, err
	}

	return technology, codes.FindTechnology, nil
}

func (service *technologyService) validateTechnology(technology *models.Technology, isCreating bool) error {
	if strings.TrimSpace(technology.Slug) == "" {
		return errors.New("invalid technology slug")
	}

	if strings.TrimSpace(technology.Name) == "" {
		return errors.New("invalid technology name")
	}

	if !isCreating {
		if strings.TrimSpace(string(technology.Status)) == "" {
			return errors.New("invalid technology status")
		}
	}

	return nil
}

func (service *technologyService) GetAll(req *dtos.TechnologyRequestDTO) ([]dtos.TechnologyDTO, codes.Code, error) {
	technologies, err := service.repository.GetAll(req)
	if err != nil {
		return nil, codes.ErrorGettingAllTechnologies, err
	}

	return mappers.TechnologiesToDTOs(technologies), codes.GetAllTechnologies, nil
}

func (service *technologyService) CountAll(req *dtos.TechnologyRequestDTO) (uint64, codes.Code, error) {
	count, err := service.repository.CountAll(req)
	if err != nil {
		return 0, codes.ErrorCoutingAllTechnologies, err
	}

	return count, codes.CountAllTechnologies, nil
}

func (service *technologyService) Find(id uint64) (*dtos.TechnologyDTO, codes.Code, error) {
	technology, code, err := service.find(id)

	return mappers.TechnologyToDTO(technology), code, err
}

func (service *technologyService) Create(createDTO *dtos.TechnologyCreateRequestDTO) (*dtos.TechnologyDTO, codes.Code, error) {
	technology := mappers.CreateRequestDTOToTechnology(createDTO)

	err := service.validateTechnology(technology, true)
	if err != nil {
		return nil, codes.InvalidParams, err
	}

	err = service.repository.Create(technology)
	if err != nil {
		return nil, codes.ErrorCreatingTechnology, err
	}

	return mappers.TechnologyToDTO(technology), codes.CreateTechnology, nil
}

func (service *technologyService) Delete(id uint64) (codes.Code, error) {
	technology, code, err := service.find(id)
	if err != nil {
		return code, err
	}

	err = service.repository.Delete(technology)
	if err != nil {
		return codes.ErrorDeletingTechnology, err
	}

	return codes.DeleteTechnology, nil
}

func (service *technologyService) Update(id uint64, updateDTO *dtos.TechnologyUpdateRequestDTO) (*dtos.TechnologyDTO, codes.Code, error) {
	technology, code, err := service.find(id)
	if err != nil {
		return nil, code, err
	}

	if updateDTO.Slug != nil {
		technology.Slug = *updateDTO.Slug
	}

	if updateDTO.Name != nil {
		technology.Name = *updateDTO.Name
	}

	err = service.validateTechnology(technology, true)
	if err != nil {
		return nil, codes.InvalidParams, err
	}

	err = service.repository.Update(technology)
	if err != nil {
		return mappers.TechnologyToDTO(technology), codes.ErrorUpdatingTechnology, err
	}

	return mappers.TechnologyToDTO(technology), codes.UpdateTechnology, nil
}

func (service *technologyService) UpdateImage(id uint64, image *multipart.FileHeader) (*dtos.TechnologyDTO, codes.Code, error) {
	technology, code, err := service.find(id)
	if err != nil {
		return nil, code, err
	}

	if technology.ImageID != nil {
		code, err := service.fileService.DeleteFile(*technology.ImageID, files.UploadConfig{
			TargetFolder: "technologies",
			AllowedTypes: []string{"image/jpeg", "image/png"},
			MaxSizeMB:    5,
		})
		if err != nil {
			return nil, code, err
		}
	}

	filename, code, err := service.fileService.SaveFile(image, files.UploadConfig{
		TargetFolder: "technologies",
		AllowedTypes: []string{"image/jpeg", "image/png"},
		MaxSizeMB:    5,
	})
	if err != nil {
		return nil, code, err
	}

	technology.ImageID = &filename

	err = service.repository.Update(technology)
	if err != nil {
		return nil, codes.ErrorFindingTechnology, err
	}

	return mappers.TechnologyToDTO(technology), codes.CreateTechnology, nil
}

func (service *technologyService) DeleteImage(id uint64) (codes.Code, error) {
	technology, code, err := service.find(id)
	if err != nil {
		return code, err
	}

	if technology.ImageID != nil {
		code, err := service.fileService.DeleteFile(*technology.ImageID, files.UploadConfig{
			TargetFolder: "technologies",
			AllowedTypes: []string{"image/jpeg", "image/png"},
			MaxSizeMB:    5,
		})
		if err != nil {
			return code, err
		}

		technology.ImageID = nil

		err = service.repository.Update(technology)
		if err != nil {
			return codes.ErrorFindingTechnology, err
		}
	}

	return codes.DeleteTechnology, nil
}

func (service *technologyService) GetImageFullPath(id uint64) (string, codes.Code, error) {
	technology, code, err := service.find(id)
	if err != nil {
		return "", code, err
	}

	if technology.ImageID == nil {
		return "", codes.FileNotFound, err
	}

	fullPath, code, err := service.fileService.GetFile(*technology.ImageID, files.UploadConfig{
		TargetFolder: "technologies",
		AllowedTypes: []string{"image/jpeg", "image/png"},
		MaxSizeMB:    5,
	})
	if err != nil {
		return "", code, err
	}

	return fullPath, codes.FileFound, nil
}
