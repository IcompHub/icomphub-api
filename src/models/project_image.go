package models

import (
	"time"
)

type ProjectImage struct {
	ID        uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	ImageID   string    `json:"image_id" gorm:"not null"`
	ProjectID uint64    `json:"project_id" gorm:"not null;index"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Project Project `json:"-" gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE"`
}
