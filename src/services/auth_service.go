package services

import (
	"errors"

	"icomphub-api/auth"
	"icomphub-api/codes"
	"icomphub-api/dtos"
	"icomphub-api/repositories"
)

type AuthService interface {
	Login(req *dtos.LoginRequestDTO) (string, codes.Code, error)
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{userRepo}
}

func (s *authService) Login(req *dtos.LoginRequestDTO) (string, codes.Code, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil || user == nil {
		return "", codes.LoginFailed, errors.New("invalid credentials")
	}

	if err := auth.CheckPasswordHash(user.Password, req.Password); err != nil {
		return "", codes.LoginFailed, errors.New("invalid credentials")
	}

	token, err := auth.GenerateJWT(*user)
	if err != nil {
		return "", codes.LoginFailed, err
	}

	return token, codes.LoginSuccess, nil
}
