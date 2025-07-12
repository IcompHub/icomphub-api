package repositories

import (
	"icomphub-api/dtos"
	"icomphub-api/models"

	"gorm.io/gorm"
)

type ProjectImageRepository interface {
	GetAll(req *dtos.ProjectImageRequestDTO) ([]models.ProjectImage, error)
	CountAll(req *dtos.ProjectImageRequestDTO) (uint64, error)
	Find(id uint64) (*models.ProjectImage, error)
	Create(projectImage *models.ProjectImage) error
	Delete(projectImage *models.ProjectImage) error
	Update(projectImage *models.ProjectImage) error
}

type projectImageRepository struct {
	db *gorm.DB
}

func NewProjectImageRepository(db *gorm.DB) ProjectImageRepository {
	return &projectImageRepository{db}
}

func (repository *projectImageRepository) makeSearchQuery(req *dtos.ProjectImageRequestDTO) *gorm.DB {
	query := repository.db.Model(&models.ProjectImage{})

	return query
}

func (repository *projectImageRepository) GetAll(req *dtos.ProjectImageRequestDTO) ([]models.ProjectImage, error) {
	offset := (req.PageNumber - 1) * req.PageSize
	var projectImages []models.ProjectImage

	query := repository.makeSearchQuery(req)

	query = query.Offset(int(offset)).Limit(int(req.PageSize))

	err := query.Preload("Technologies").Find(&projectImages).Error
	return projectImages, err
}

func (repository *projectImageRepository) CountAll(req *dtos.ProjectImageRequestDTO) (uint64, error) {
	var count int64 = 0

	query := repository.makeSearchQuery(req)

	err := query.Count(&count).Error
	return uint64(count), err
}

func (repository *projectImageRepository) Find(id uint64) (*models.ProjectImage, error) {
	var projectImage *models.ProjectImage
	err := repository.db.Model(&models.ProjectImage{}).Preload("Technologies").First(&projectImage, id).Error

	return projectImage, err
}

func (repository *projectImageRepository) Create(projectImage *models.ProjectImage) error {
	return repository.db.Model(&models.ProjectImage{}).Create(&projectImage).Error
}

func (repository *projectImageRepository) Delete(projectImage *models.ProjectImage) error {
	return repository.db.Model(&models.ProjectImage{}).Where("id = ?", projectImage.ID).Delete(&projectImage).Error
}

func (repository *projectImageRepository) Update(projectImage *models.ProjectImage) error {
	return repository.db.Model(&models.ProjectImage{}).Where("id = ?", projectImage.ID).Save(&projectImage).Error
}
