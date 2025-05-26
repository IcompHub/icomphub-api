package models

import (
	"time"

	"icomphub-api/enums"
)

type Technology struct {
	ID        uint64           `json:"id" gorm:"primaryKey;autoIncrement"`
	Slug      string           `json:"slug" gorm:"unique;not null"`
	Name      string           `json:"name" gorm:"unique;not null"`
	Status    enums.StatusEnum `json:"status" gorm:"type:status_enum;default:'waiting_approval'"`
	CreatedAt time.Time        `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time        `json:"updated_at" gorm:"autoUpdateTime"`
}
