package mappers

import (
	"icomphub-api/dtos"
	"icomphub-api/models"
)

func UserToDTO(user *models.User) *dtos.UserDTO {
	return &dtos.UserDTO{
		Id:               user.ID,
		Slug:             user.Slug,
		Nickname:         user.Nickname,
		PersonalEmail:    user.PersonalEmail,
		ProfilePictureId: user.ProfilePictureId,
		Role:             user.SystemRole,
	}
}

func CreateRequestDTOToUser(createDTO *dtos.UserCreateRequestDTO) *models.User {
	return &models.User{
		Slug:               createDTO.Slug,
		Nickname:           createDTO.Nickname,
		FullName:           createDTO.FullName,
		PersonalEmail:      createDTO.PersonalEmail,
		InstitutionalEmail: createDTO.InstitutionalEmail,
		Registration:       createDTO.Registration,
		Password:           createDTO.Password,
	}
}

func UsersToDTOs(users []models.User) []dtos.UserDTO {
	dtosList := make([]dtos.UserDTO, len(users))
	for i, user := range users {
		dtosList[i] = *UserToDTO(&user)
	}
	return dtosList
}
