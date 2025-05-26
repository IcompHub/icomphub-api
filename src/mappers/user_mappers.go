package mappers

import (
	"icomphub-api/dtos"
	"icomphub-api/models"
)

func UserToDTO(user *models.User) dtos.UserDTO {
	return dtos.UserDTO{
		Id:            user.ID,
		Slug:          user.Slug,
		Nickname:      user.Nickname,
		PersonalEmail: user.PersonalEmail,
		Role:          user.SystemRole,
	}
}

func UsersToDTOs(users []models.User) []dtos.UserDTO {
	dtosList := make([]dtos.UserDTO, len(users))
	for i, user := range users {
		dtosList[i] = UserToDTO(&user)
	}
	return dtosList
}
