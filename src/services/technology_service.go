package services

import (
	"fmt"
	"log"

	"icomphub-api/dto"
	"icomphub-api/models"
	"icomphub-api/repositories"
)

type TechnologyService interface {
	CreateTechnology(techDTO *dto.TechnologyDTO) error
	DeleteTechnology(id uint64) error
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

	// Opcional: Adicionar lógica de negócio aqui antes de deletar, se necessário.
	// Por exemplo, verificar se a tecnologia pode ser deletada com base em seu status ou dependências.
	// Lembre-se que repo.DeleteTechnology já verifica a existência.

	err := s.repo.DeleteTechnology(id)
	if err != nil {
		log.Printf("Serviço: Erro ao deletar tecnologia ID %d no repositório: %v", id, err)
		return fmt.Errorf("serviço: falha ao deletar tecnologia com ID %d: %w", id, err)
	}

	log.Printf("Serviço: Tecnologia com ID %d marcada para deleção com sucesso.", id)
	return nil
}
