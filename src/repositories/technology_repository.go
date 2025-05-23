package repositories

import (
	"gorm.io/gorm"
	"icomphub-api/models"
)

type TechnologyRepository interface {
	CreateTechnology(tech *models.Technology) error
}

type TechnologyRepository struct {
	db *gorm.DB
}

func NewTechnologyRepository(db *gorm.DB) TechnologyRepository {
	return &TechnologyRepository{db}
}

func (r *TechnologyRepository) CreateTechnology(tech *models.Technology) error {
	return r.db.Create(tech).Error
}
