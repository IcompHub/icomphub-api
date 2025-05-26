package repositories

import (
	"strings"

	"icomphub-api/dtos"
	"icomphub-api/models"

	"gorm.io/gorm"
)

type TechnologyRepository interface {
	GetAll(req *dtos.TechnologyRequestDTO) ([]models.Technology, error)
	CountAll(req *dtos.TechnologyRequestDTO) (uint64, error)
	Find(id uint64) (*models.Technology, error)
	Create(technology *models.Technology) error
	Delete(technology *models.Technology) error
	Update(technology *models.Technology) error
}

type technologyRepository struct {
	db *gorm.DB
}

func NewTechnologyRepository(db *gorm.DB) TechnologyRepository {
	return &technologyRepository{db: db}
}

func (repository *technologyRepository) makeSearchQuery(req *dtos.TechnologyRequestDTO) *gorm.DB {
	query := repository.db.Model(&models.Technology{})

	if strings.TrimSpace(req.Search) != "" {
		searchPattern := "%" + req.Search + "%"
		query = query.Where("name ILIKE ?", searchPattern)
	}

	return query
}

func (repository *technologyRepository) GetAll(req *dtos.TechnologyRequestDTO) ([]models.Technology, error) {
	offset := (req.PageNumber - 1) * req.PageSize
	var technologies []models.Technology

	query := repository.makeSearchQuery(req)

	query = query.Offset(int(offset)).Limit(int(req.PageSize))

	err := query.Find(&technologies).Error
	return technologies, err
}

func (repository *technologyRepository) CountAll(req *dtos.TechnologyRequestDTO) (uint64, error) {
	var count int64 = 0

	query := repository.makeSearchQuery(req)

	err := query.Count(&count).Error
	return uint64(count), err
}

func (repository *technologyRepository) Find(id uint64) (*models.Technology, error) {
	var technology *models.Technology
	err := repository.db.First(&technology, id).Error

	return technology, err
}

func (repository *technologyRepository) Create(technology *models.Technology) error {
	return repository.db.Create(&technology).Error
}

func (repository *technologyRepository) Delete(technology *models.Technology) error {
	return repository.db.Delete(&technology).Error
}

func (repository *technologyRepository) Update(technology *models.Technology) error {
	return repository.db.Save(&technology).Error
}
