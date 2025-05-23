package service

import (
	"icomphub-api/models"
	"icomphub-api/repositories"
)

type TechnologyService interface {
	CreateTechnology(tech *models.Technology) error
}

type technologyService struct {
	repo repositories.TechnologyRepository
}

func NewTechnologyService(r repositories.TechnologyRepository) TechnologyService {
	return &technologyService{repo: r}
}

func (s *technologyService) CreateTechnology(tech *models.Technology) error {
	// fazer validações extras aqui
	return s.repo.CreateTechnology(tech)
}
