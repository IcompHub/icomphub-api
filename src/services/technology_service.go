package services

import (
	"icomphub-api/dto"
	"icomphub-api/models"
	"icomphub-api/repositories"
)

type TechnologyService interface {
	CreateTechnology(techDTO *dto.TechnologyDTO) error
}

type technologyService struct {
	repo repositories.TechnologyRepository
}

func NewTechnologyService(r repositories.TechnologyRepository) TechnologyService {
	return &technologyService{repo: r}
}

func (s *technologyService) CreateTechnology(techDTO *dto.TechnologyDTO) error {
	// fazer validações extras aqui

	newTechnology := &models.Technology{
		Name:   techDTO.Name,
		Slug:   techDTO.Slug,
		Status: models.StatusWaitingApproval,
	}

	err := s.repo.CreateTechnology(newTechnology)
	if err != nil {
		return err
	}

	return nil
}
