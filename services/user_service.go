package services

import (
	"icomphub-api/models"
	"icomphub-api/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository


}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

//posso adicionar lógica futuramente
func (s *UserService) GetAllUsers() ([]models.User, error){
	return s.userRepo.FindAll()
}