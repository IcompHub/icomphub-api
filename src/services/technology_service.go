package services

import (
	"icomphub-api/codes"
	"icomphub-api/dtos"
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
}

type technologyService struct {
	repository repositories.TechnologyRepository
}

func NewTechnologyService(repository repositories.TechnologyRepository) TechnologyService {
	return &technologyService{repository: repository}
}

func (service *technologyService) find(id uint64) (*models.Technology, codes.Code, error) {
	technology, err := service.repository.Find(id)
	if err != nil {
		return nil, codes.ErrorFindingTechnology, err
	}

	return technology, codes.FindTechnology, nil
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

	err := service.repository.Create(technology)
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

	if updateDTO.Slug != "" {
		technology.Slug = updateDTO.Slug
	}

	if updateDTO.Name != "" {
		technology.Name = updateDTO.Name
	}

	err = service.repository.Update(technology)
	if err != nil {
		return mappers.TechnologyToDTO(technology), codes.ErrorUpdatingTechnology, err
	}

	return mappers.TechnologyToDTO(technology), codes.UpdateTechnology, nil
}
