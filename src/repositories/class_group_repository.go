package repositories

import (
	"strings"

	"icomphub-api/dtos"
	"icomphub-api/models"

	"gorm.io/gorm"
)

type ClassGroupRepository interface {
	GetAll(req *dtos.ClassGroupRequestDTO) ([]models.ClassGroup, error)
	CountAll(req *dtos.ClassGroupRequestDTO) (uint64, error)
	Find(id uint64) (*models.ClassGroup, error)
	Create(classGroup *models.ClassGroup) error
	Delete(classGroup *models.ClassGroup) error
	Update(classGroup *models.ClassGroup) error
}

type classGroupRepository struct {
	db *gorm.DB
}

func NewClassGroupRepository(db *gorm.DB) ClassGroupRepository {
	return &classGroupRepository{db}
}

func (repository *classGroupRepository) makeSearchQuery(req *dtos.ClassGroupRequestDTO) *gorm.DB {
	query := repository.db.Model(&models.ClassGroup{})

	if strings.TrimSpace(req.Search) != "" {
		searchPattern := "%" + req.Search + "%"
		query = query.Where("name ILIKE ?", searchPattern)
	}

	return query
}

func (repository *classGroupRepository) GetAll(req *dtos.ClassGroupRequestDTO) ([]models.ClassGroup, error) {
	offset := (req.PageNumber - 1) * req.PageSize
	var classGroups []models.ClassGroup

	query := repository.makeSearchQuery(req)

	query = query.Offset(int(offset)).Limit(int(req.PageSize))

	err := query.Find(&classGroups).Error
	return classGroups, err
}

func (repository *classGroupRepository) CountAll(req *dtos.ClassGroupRequestDTO) (uint64, error) {
	var count int64 = 0

	query := repository.makeSearchQuery(req)

	err := query.Count(&count).Error
	return uint64(count), err
}

func (repository *classGroupRepository) Find(id uint64) (*models.ClassGroup, error) {
	var classGroup *models.ClassGroup
	err := repository.db.Model(&models.ClassGroup{}).First(&classGroup, id).Error

	return classGroup, err
}

func (repository *classGroupRepository) Create(classGroup *models.ClassGroup) error {
	return repository.db.Model(&models.ClassGroup{}).Create(&classGroup).Error
}

func (repository *classGroupRepository) Delete(classGroup *models.ClassGroup) error {
	return repository.db.Model(&models.ClassGroup{}).Where("id = ?", classGroup.ID).Delete(&classGroup).Error
}

func (repository *classGroupRepository) Update(classGroup *models.ClassGroup) error {
	return repository.db.Model(&models.ClassGroup{}).Where("id = ?", classGroup.ID).Save(&classGroup).Error
}
