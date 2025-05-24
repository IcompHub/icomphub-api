package repositories

import (
	"gorm.io/gorm"
	"icomphub-api/models"
)

type TechnologyRepository interface {
	CreateTechnology(tech *models.Technology) error
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
