package repositories

import (
	"strings"

	"icomphub-api/dtos"
	"icomphub-api/models"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	GetAll(req *dtos.ProjectRequestDTO) ([]models.Project, error)
	CountAll(req *dtos.ProjectRequestDTO) (uint64, error)
	Find(id uint64) (*models.Project, error)
	Create(project *models.Project) error
	Delete(project *models.Project) error
	Update(project *models.Project) error
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{db}
}

func (repository *projectRepository) makeSearchQuery(req *dtos.ProjectRequestDTO) *gorm.DB {
	query := repository.db.Model(&models.Project{})

	if strings.TrimSpace(req.Search) != "" {
		searchPattern := "%" + req.Search + "%"
		query = query.Where("name ILIKE ?", searchPattern)
	}

	// if req.ClassGroupID != nil {
	// 	query = query.Where("class_group_id = ?", *req.ClassGroupID)
	// }

	return query
}

func (repository *projectRepository) GetAll(req *dtos.ProjectRequestDTO) ([]models.Project, error) {
	offset := (req.PageNumber - 1) * req.PageSize
	var projects []models.Project

	query := repository.makeSearchQuery(req)

	query = query.Offset(int(offset)).Limit(int(req.PageSize))

	err := query.Find(&projects).Error
	return projects, err
}

func (repository *projectRepository) CountAll(req *dtos.ProjectRequestDTO) (uint64, error) {
	var count int64 = 0

	query := repository.makeSearchQuery(req)

	err := query.Count(&count).Error
	return uint64(count), err
}

func (repository *projectRepository) Find(id uint64) (*models.Project, error) {
	var project *models.Project
	err := repository.db.Model(&models.Project{}).First(&project, id).Error

	return project, err
}

func (repository *projectRepository) Create(project *models.Project) error {
	return repository.db.Model(&models.Project{}).Create(&project).Error
}

func (repository *projectRepository) Delete(project *models.Project) error {
	return repository.db.Model(&models.Project{}).Where("id = ?", project.ID).Delete(&project).Error
}

func (repository *projectRepository) Update(project *models.Project) error {
	return repository.db.Model(&models.Project{}).Where("id = ?", project.ID).Save(&project).Error
}
