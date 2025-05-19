package repositories

import (
	"icomphub-api/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FindAll() ([]models.User, error){
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}