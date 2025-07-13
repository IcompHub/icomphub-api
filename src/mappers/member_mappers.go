package mappers

import (
	"icomphub-api/dtos"
	"icomphub-api/models"
)

func MemberToDTO(Member *models.Member) *dtos.MemberDTO {
	var userDTO *dtos.UserDTO = nil

	if Member.User != nil {
		userDTO = UserToDTO(Member.User)
	}

	return &dtos.MemberDTO{
		ID:        Member.ID,
		Nickname:  Member.Nickname,
		Status:    string(Member.Status),
		ProjectID: Member.ProjectID,
		UserId:    Member.UserId,
		User:      userDTO,
		Project:   *ProjectToShortDTO(&Member.Project),
		Roles:     RolesToDTOs(Member.Roles),
	}
}

func MemberToShortDTO(Member *models.Member) *dtos.MemberShortDTO {
	return &dtos.MemberShortDTO{
		ID:       Member.ID,
		Nickname: Member.Nickname,
		Status:   string(Member.Status),
		Roles:    RolesToDTOs(Member.Roles),
	}
}

func CreateRequestDTOToMember(createDTO *dtos.MemberCreateRequestDTO) *models.Member {
	return &models.Member{
		Nickname:  createDTO.Nickname,
		ProjectID: createDTO.ProjectID,
		UserId:    createDTO.UserId,
	}
}

func MembersToDTOs(members []models.Member) []dtos.MemberDTO {
	dtosList := make([]dtos.MemberDTO, len(members))
	for i, member := range members {
		dtosList[i] = *MemberToDTO(&member)
	}
	return dtosList
}

func MembersToShortDTOs(members []models.Member) []dtos.MemberShortDTO {
	dtosList := make([]dtos.MemberShortDTO, len(members))
	for i, member := range members {
		dtosList[i] = *MemberToShortDTO(&member)
	}
	return dtosList
}
