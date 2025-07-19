package repositories

import (
	"icomphub-api/models"

	"gorm.io/gorm"
)

type MemberRepository interface {
	FindByID(id uint64) (*models.Member, error)
	FindByProject(projectID uint64) ([]models.Member, error)
	Create(member *models.Member) error
	Update(member *models.Member) error
	Delete(member *models.Member) error
	ReplaceRoles(memberID uint64, roleIDs []uint64) error
	RemoveRoles(memberID uint64, roleIDs []uint64) error
}

type memberRepository struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) MemberRepository {
	return &memberRepository{db}
}

func (r *memberRepository) FindByID(id uint64) (*models.Member, error) {
	var member models.Member
	err := r.db.Preload("User").Preload("Roles").First(&member, id).Error
	return &member, err
}

func (r *memberRepository) FindByProject(projectID uint64) ([]models.Member, error) {
	var members []models.Member
	err := r.db.
		Preload("User").
		Preload("Roles").
		Where("project_id = ?", projectID).
		Find(&members).Error
	return members, err
}

func (r *memberRepository) Create(member *models.Member) error {
	return r.db.Create(&member).Error
}

func (r *memberRepository) Update(member *models.Member) error {
	return r.db.Save(&member).Error
}

func (r *memberRepository) Delete(member *models.Member) error {
	return r.db.Delete(&member).Error
}

func (r *memberRepository) ReplaceRoles(memberID uint64, roleIDs []uint64) error {
	var member models.Member
	if err := r.db.First(&member, memberID).Error; err != nil {
		return err
	}

	var roles []models.Role
	if err := r.db.Where("id IN ?", roleIDs).Find(&roles).Error; err != nil {
		return err
	}

	return r.db.Model(&member).Association("Roles").Replace(&roles)
}

func (r *memberRepository) RemoveRoles(memberID uint64, roleIDs []uint64) error {
	var member models.Member
	if err := r.db.First(&member, memberID).Error; err != nil {
		return err
	}

	var roles []models.Role
	if err := r.db.Where("id IN ?", roleIDs).Find(&roles).Error; err != nil {
		return err
	}

	return r.db.Model(&member).Association("Roles").Delete(&roles)
}
