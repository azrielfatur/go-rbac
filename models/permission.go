package models

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	ID          int16  `gorm:"not null; primary_key"`
	Name        string `gorm:"size:100; not null"`
	Description string `gorm:"size:100; default: NULL"`
	Active      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type TablerPermission interface {
	TableName() string
}

func (Permission) TableName() string {
	return "ms_permissions"
}
