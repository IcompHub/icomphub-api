package repositories

import (
	"strings"

	"icomphub-api/dtos"
	"icomphub-api/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll(req *dtos.UserRequestDTO) ([]models.User, error)
	CountAll(req *dtos.UserRequestDTO) (uint64, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (repository *userRepository) makeSearchQuery(req *dtos.UserRequestDTO) *gorm.DB {
	var query *gorm.DB = repository.db.Model(&models.User{})

	if strings.TrimSpace(req.Search) != "" {
		searchPattern := "%" + req.Search + "%"
		query = query.Where("nickname ILIKE ?", searchPattern)
	}

	return query
}

func (repository *userRepository) GetAll(req *dtos.UserRequestDTO) ([]models.User, error) {
	offset := (req.PageNumber - 1) * req.PageSize
	var users []models.User

	var query *gorm.DB = repository.makeSearchQuery(req)

	query = query.Offset(int(offset)).Limit(int(req.PageSize))

	err := query.Find(&users).Error
	return users, err
}

func (repository *userRepository) CountAll(req *dtos.UserRequestDTO) (uint64, error) {
	var count int64 = 0

	var query *gorm.DB = repository.makeSearchQuery(req)

	err := query.Count(&count).Error
	return uint64(count), err
}
