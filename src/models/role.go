package models

import "time"

type Role struct {
	ID        uint64 `gorm:"primaryKey"`
	Slug      string `gorm:"unique"`
	Name      string `gorm:"unique"`
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
