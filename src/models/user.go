package models

import (
	"time"

	"icomphub-api/enums"
)

type User struct {
	ID                 uint64               `json:"id" gorm:"primaryKey;autoIncrement"`
	Slug               string               `json:"slug" gorm:"unique;not null"`
	Nickname           string               `json:"nickname" gorm:"unique;not null"`
	FullName           string               `json:"full_name" gorm:"not null"`
	PersonalEmail      string               `json:"personal_email" gorm:"unique;not null"`
	InstitutionalEmail *string              `json:"institutional_email,omitempty" gorm:"unique"`
	Registration       *string              `json:"registration,omitempty" gorm:"unique"`
	Password           string               `json:"-" gorm:"not null"` // do not show
	Status             enums.StatusEnum     `json:"status" gorm:"type:status_enum;default:'waiting_approval'"`
	SystemRole         enums.SystemRoleEnum `json:"system_role" gorm:"type:system_role_enum;default:'user'"`
	CreatedAt          time.Time            `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time            `json:"updated_at" gorm:"autoUpdateTime"`
}
