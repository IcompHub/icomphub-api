package models

import (
	"time"

	"icomphub-api/enums"

	"gorm.io/datatypes"
)

type Member struct {
	ID        uint64 `gorm:"primaryKey"`
	Nickname  string
	Status    string
	ProjectID uint64
	UserID    uint64
	RoleID    uint64
	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"foreignKey:UserID"`
	Role Role `gorm:"foreignKey:RoleID"`
}

type Project struct {
	ID           uint64 `gorm:"primaryKey"`
	Slug         string `gorm:"unique"`
	Name         string `gorm:"unique"`
	Status       string
	Data         datatypes.JSON
	ClassGroupID uint64
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Members []Member `gorm:"foreignKey:ProjectID"`

	Technologies []Technology `gorm:"many2many:projects_technologies;"`
}
