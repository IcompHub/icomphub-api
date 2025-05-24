package repositories

import (
	"errors"
	"fmt"

	"icomphub-api/models"

	"gorm.io/gorm"
)

type TechnologyRepository interface {
	CreateTechnology(tech *models.Technology) error
	DeleteTechnology(id uint64) error

}

type technologyRepo struct {
	db *gorm.DB
}

func NewTechnologyRepository(db *gorm.DB) TechnologyRepository {
	return &technologyRepo{db: db}
}

func (r *technologyRepo) CreateTechnology(tech *models.Technology) error {
	return r.db.Create(tech).Error
}

func (r *technologyRepo) DeleteTechnology(id uint64) error {
	var technology models.Technology
	err := r.db.First(&technology, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("tecnologia com ID %d não encontrada: %w", id, gorm.ErrRecordNotFound)
		}

		return fmt.Errorf("erro ao buscar tecnologia com ID %d: %w", id, err)
	}

	if err := r.db.Delete(&technology).Error; err != nil {
		return fmt.Errorf("erro ao deletar tecnologia com ID %d: %w", id, err)
	}

	return nil
}


