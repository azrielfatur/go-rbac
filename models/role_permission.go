package models

import (
	"time"

	"gorm.io/gorm"
)

type RolePermission struct {
	ID           int16 `gorm:"not null; primary_key"`
	RoleID       int16 `gorm:"not null; index"`
	PermissionID int16 `gorm:"not null; index"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

type TablerRolePermission interface {
	TableName() string
}

func (RolePermission) TableName() string {
	return "tr_role_permissions"
}
