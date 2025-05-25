package services

import (
	"fmt"
	"log"

	"icomphub-api/dto"
	"icomphub-api/models"
	"icomphub-api/repositories"
)

type TechnologyService interface {
	GetTechnology() ([]dto.TechnologyGetDTO, error)
	CreateTechnology(techDTO *dto.TechnologyDTO) error
	DeleteTechnology(id uint64) error
	UpdateTechnology(techGetDTO *dto.TechnologyGetDTO) error
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

func (s *technologyService) DeleteTechnology(id uint64) error {
	log.Printf("Serviço: Tentando deletar tecnologia com ID: %d", id)

	err := s.repo.DeleteTechnology(id)
	if err != nil {
		log.Printf("Serviço: Erro ao deletar tecnologia ID %d no repositório: %v", id, err)
		return fmt.Errorf("serviço: falha ao deletar tecnologia com ID %d: %w", id, err)
	}

	log.Printf("Serviço: Tecnologia com ID %d marcada para deleção com sucesso.", id)
	return nil
}

func (s *technologyService) UpdateTechnology(techGetDTO *dto.TechnologyGetDTO) error {
	err := s.repo.UpdateTechnology(techGetDTO.ID, techGetDTO.Name, techGetDTO.Slug)
	if err != nil {
		return fmt.Errorf("serviço: falha ao atualizar tecnologia com ID %d: %w", techGetDTO.ID, err)
	}

	return nil
}

func (s *technologyService) GetTechnology() ([]dto.TechnologyGetDTO, error) {
	modelTechnologies, err := s.repo.GetTechnology()
	if err != nil {
		return nil, err
	}

	var dtoTechnologies []dto.TechnologyGetDTO

	for _, modelTech := range modelTechnologies {
		dtoTechnologies = append(dtoTechnologies, dto.TechnologyGetDTO{
			ID:   modelTech.ID,
			Name: modelTech.Name,
			Slug: modelTech.Slug,
		})
	}

	return dtoTechnologies, nil
}
