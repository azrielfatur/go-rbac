package models

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	ID        int16  `gorm:"not null; primary_key"`
	Name      string `gorm:"not null; size:100"`
	Url       string `gorm:"not null; size:100"`
	Icon      string `gorm:"default: NULL"`
	Order     int    `gorm:"default: NULL"`
	ParentID  int    `gorm:"default: NULL"`
	Childs    []Menu `gorm:"foreignKey:ParentID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type TablerMenu interface {
	TableName() string
}

func (Menu) TableName() string {
	return "ms_menus"
}
