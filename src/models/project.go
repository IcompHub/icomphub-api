package models

import (
	"time"

	"icomphub-api/enums"

	"gorm.io/datatypes"
)

type Project struct {
	ID           uint64           `json:"id" gorm:"primaryKey;autoIncrement"`
	Slug         string           `json:"slug" gorm:"unique;not null"`
	Name         string           `json:"name" gorm:"unique;not null"`
	Status       enums.StatusEnum `json:"status" gorm:"type:status_enum;default:'waiting_approval'"`
	ThumbnailID  *string          `json:"thumbnail_id"`
	Data         datatypes.JSON   `json:"data" gorm:"type:jsonb;not null"`
	ClassGroupID uint64           `json:"class_group_id" gorm:"not null;index"`
	CreatedAt    time.Time        `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time        `json:"updated_at" gorm:"autoUpdateTime"`

	ClassGroup    ClassGroup     `json:"-" gorm:"constraint:OnDelete:CASCADE"`
	Members       []Member       `json:"members" gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE"`
	Technologies  []Technology   `gorm:"many2many:project_technologies;"`
	ProjectImages []ProjectImage `json:"project_images" gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE"`
}
