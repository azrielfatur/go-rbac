package models

import (
	"time"

	"gorm.io/gorm"
)

type UserRole struct {
	ID        int16 `gorm:"not null; primary_key"`
	UserID    int16 `gorm:"not null; index"`
	RoleID    int16 `gorm:"not null; index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type TablerUserRole interface {
	TableName() string
}

func (UserRole) TableName() string {
	return "tr_user_roles"
}
