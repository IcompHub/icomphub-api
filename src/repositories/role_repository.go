package repositories

import (
	"icomphub-api/models"

	"gorm.io/gorm"
)

type RoleRepository interface {
	FindAll() ([]models.Role, error)
	FindByID(id uint64) (*models.Role, error)
	FindBySlug(slug string) (*models.Role, error)
	Create(role *models.Role) error
	Update(role *models.Role) error
	Delete(role *models.Role) error
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db}
}

func (r *roleRepository) FindAll() ([]models.Role, error) {
	var roles []models.Role
	err := r.db.Find(&roles).Error
	return roles, err
}

func (r *roleRepository) FindByID(id uint64) (*models.Role, error) {
	var role models.Role
	err := r.db.First(&role, id).Error
	return &role, err
}

func (r *roleRepository) FindBySlug(slug string) (*models.Role, error) {
	var role models.Role
	err := r.db.Where("slug = ?", slug).First(&role).Error
	return &role, err
}

func (r *roleRepository) Create(role *models.Role) error {
	return r.db.Create(&role).Error
}

func (r *roleRepository) Update(role *models.Role) error {
	return r.db.Save(&role).Error
}

func (r *roleRepository) Delete(role *models.Role) error {
	return r.db.Delete(&role).Error
}
