package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID          int16        `gorm:"not null; primary_key"`
	Name        string       `gorm:"size:100; not null"`
	Description string       `gorm:"size:100; default: NULL"`
	Active      bool         `gorm:"not null" default:"true"`
	Permissions []Permission `gorm:"many2many:tr_role_permissions;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type TablerRole interface {
	TableName() string
}

func (Role) TableName() string {
	return "ms_roles"
}
