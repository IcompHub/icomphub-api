package dtos

type LoginRequestDTO struct {
	Email    string `json:"email" binding:"required,email" example:"nelson.dev@test.com"`
	Password string `json:"password" binding:"required" example:"123456"`
}
