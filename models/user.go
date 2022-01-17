package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              int16     `gorm:"not null; primary_key"`
	Username        string    `gorm:"size:100;not null"`
	Name            string    `gorm:"size:100;not null"`
	Email           string    `gorm:"size:100;not null"`
	Password        string    `gorm:"size:100;not null"`
	Role            []Role    `gorm:"many2many:tr_user_roles"`
	RememberToken   string    `gorm:"default: NULL"`
	EmailVerifiedAt time.Time `gorm:"default: NULL"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

type TablerUser interface {
	TableName() string
}

func (User) TableName() string {
	return "ms_users"
}
