package services

import (
	"icomphub-api/codes"
	"icomphub-api/dtos"
	"icomphub-api/mappers"
	"icomphub-api/models"
	"icomphub-api/repositories"
)

type MemberService interface {
	GetByID(id uint64) (*dtos.MemberDTO, codes.Code, error)
	GetByProject(projectID uint64) ([]dtos.MemberDTO, codes.Code, error)
	Create(dto *dtos.MemberCreateRequestDTO) (*dtos.MemberDTO, codes.Code, error)
	Update(id uint64, dto *dtos.MemberUpdateRequestDTO) (*dtos.MemberDTO, codes.Code, error)
	Delete(id uint64) (codes.Code, error)
}

type memberService struct {
	repository  repositories.MemberRepository
	roleService RoleService
}

func NewMemberService(repo repositories.MemberRepository, roleService RoleService) MemberService {
	return &memberService{repository: repo, roleService: roleService}
}

func (s *memberService) GetByID(id uint64) (*dtos.MemberDTO, codes.Code, error) {
	member, err := s.repository.FindByID(id)
	if err != nil {
		return nil, codes.ErrorFindingMember, err
	}
	dto := mappers.MemberToDTO(member)
	return dto, codes.FindMember, nil
}

func (s *memberService) GetByProject(projectID uint64) ([]dtos.MemberDTO, codes.Code, error) {
	members, err := s.repository.FindByProject(projectID)
	if err != nil {
		return nil, codes.ErrorFindingMember, err
	}
	dtos := mappers.MembersToDTOs(members)
	return dtos, codes.FindMember, nil
}

func (s *memberService) Create(dto *dtos.MemberCreateRequestDTO) (*dtos.MemberDTO, codes.Code, error) {
	member := &models.Member{
		Nickname:  dto.Nickname,
		ProjectID: dto.ProjectID,
		UserId:    dto.UserId,
	}

	if err := s.repository.Create(member); err != nil {
		return nil, codes.ErrorCreatingMember, err
	}

	var validRoleIDs []uint64
	for _, roleID := range dto.RoleIDs {
		role, code, err := s.roleService.GetByID(roleID)
		if err != nil {
			return nil, code, err
		}
		validRoleIDs = append(validRoleIDs, role.Id)
	}

	if err := s.repository.ReplaceRoles(member.ID, validRoleIDs); err != nil {
		return nil, codes.ErrorCreatingMember, err
	}

	fullMember, err := s.repository.FindByID(member.ID)
	if err != nil {
		return nil, codes.ErrorCreatingMember, err
	}

	return mappers.MemberToDTO(fullMember), codes.CreateMember, nil
}

func (s *memberService) Update(id uint64, dto *dtos.MemberUpdateRequestDTO) (*dtos.MemberDTO, codes.Code, error) {
	member, err := s.repository.FindByID(id)
	if err != nil {
		return nil, codes.ErrorFindingMember, err
	}

	if dto.Nickname != nil {
		member.Nickname = *dto.Nickname
	}

	err = s.repository.Update(member)
	if err != nil {
		return nil, codes.ErrorUpdatingMember, err
	}

	return mappers.MemberToDTO(member), codes.UpdateMember, nil
}

func (s *memberService) Delete(id uint64) (codes.Code, error) {
	member, err := s.repository.FindByID(id)
	if err != nil {
		return codes.ErrorFindingMember, err
	}

	err = s.repository.Delete(member)
	if err != nil {
		return codes.ErrorDeletingMember, err
	}

	return codes.DeleteMember, nil
}
