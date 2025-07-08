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
	Find(id uint64) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Create(user *models.User) error
	Delete(user *models.User) error
	Update(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (repository *userRepository) makeSearchQuery(req *dtos.UserRequestDTO) *gorm.DB {
	query := repository.db.Model(&models.User{})

	if strings.TrimSpace(req.Search) != "" {
		searchPattern := "%" + req.Search + "%"
		query = query.Where("nickname ILIKE ?", searchPattern)
	}

	return query
}

func (repository *userRepository) GetAll(req *dtos.UserRequestDTO) ([]models.User, error) {
	offset := (req.PageNumber - 1) * req.PageSize
	var users []models.User

	query := repository.makeSearchQuery(req)

	query = query.Offset(int(offset)).Limit(int(req.PageSize))

	err := query.Find(&users).Error
	return users, err
}

func (repository *userRepository) CountAll(req *dtos.UserRequestDTO) (uint64, error) {
	var count int64 = 0

	query := repository.makeSearchQuery(req)

	err := query.Count(&count).Error
	return uint64(count), err
}

func (repository *userRepository) Find(id uint64) (*models.User, error) {
	var user *models.User
	err := repository.db.Model(&models.User{}).First(&user, id).Error

	return user, err
}

func (repository *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := repository.db.Model(&models.User{}).Where("personal_email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *userRepository) Create(user *models.User) error {
	return repository.db.Model(&models.User{}).Create(&user).Error
}

func (repository *userRepository) Delete(user *models.User) error {
	return repository.db.Model(&models.User{}).Where("id = ?", user.ID).Delete(&user).Error
}

func (repository *userRepository) Update(user *models.User) error {
	return repository.db.Model(&models.User{}).Where("id = ?", user.ID).Save(&user).Error
}
