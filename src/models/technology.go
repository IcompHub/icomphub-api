package models

import "gorm.io/gorm"

type (
	StatusEnum string
)

const (
	StatusActive          StatusEnum = "active"
	StatusInactive        StatusEnum = "inactive"
	StatusWaitingApproval StatusEnum = "waiting_approval"
)

type Technology struct {
	ID        uint64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Slug      string     `json:"slug" gorm:"unique;not null"`
	Name      string     `json:"name" gorm:"unique;not null"`
	Status    StatusEnum `json:"status" gorm:"type:status_enum;default:'waiting_approval'"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}
