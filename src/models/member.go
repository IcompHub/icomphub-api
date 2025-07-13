package models

import (
	"time"

	"icomphub-api/enums"
)

type Member struct {
	ID        uint64           `json:"id" gorm:"primaryKey;autoIncrement"`
	Nickname  string           `json:"nickname"`
	Status    enums.StatusEnum `json:"status" gorm:"type:status_enum;default:'waiting_approval'"`
	ProjectID uint64           `json:"project_id" gorm:"not null;index"`
	UserId    *uint64          `json:"user_id" gorm:"index"`
	CreatedAt time.Time        `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time        `json:"updated_at" gorm:"autoUpdateTime"`

	Project Project `json:"-" gorm:"constraint:OnDelete:CASCADE"`
	User    *User   `json:"-" gorm:"constraint:OnDelete:CASCADE"`
	Roles   []Role  `gorm:"many2many:member_roles;"`
}
