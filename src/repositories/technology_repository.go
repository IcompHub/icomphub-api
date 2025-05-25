package repositories

import (
	"errors"
	"fmt"

	"icomphub-api/models"

	"gorm.io/gorm"
)

type TechnologyRepository interface {
	GetTechnology() ([]models.Technology, error)
	CreateTechnology(tech *models.Technology) error
	DeleteTechnology(id uint64) error
	UpdateTechnology(id uint64, name string, slug string) error
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

func (r *technologyRepo) UpdateTechnology(id uint64, name string, slug string) error {
	var technology models.Technology
	err := r.db.First(&technology, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("tecnologia com ID %d não encontrada: %w", id, gorm.ErrRecordNotFound)
		}

		return fmt.Errorf("erro ao buscar tecnologia com ID %d: %w", id, err)
	}

	technology.Name = name
	technology.Slug = slug

	err = r.db.Save(technology).Error

	return err
}

func (r *technologyRepo) GetTechnology() ([]models.Technology, error) {
	var technologies []models.Technology
	err := r.db.Find(&technologies).Error
	return technologies, err
}
