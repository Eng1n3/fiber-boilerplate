package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	ID          uuid.UUID `gorm:"primaryKey;not null" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `gorm:"null" json:"-"`
	Path        string    `gorm:"not null" json:"path"`
	Method      string    `gorm:"not null" json:"method"`
	CreatedAt   time.Time `gorm:"autoCreateTime:milli" json:"-"`
	UpdatedAt   time.Time `gorm:"autoCreateTime:milli;autoUpdateTime:milli" json:"-"`
	DeletedAt   time.Time `gorm:"index" json:"-"`
}

func (permission *Permission) BeforeCreate(_ *gorm.DB) error {
	permission.ID = uuid.New() // Generate UUID before create
	return nil
}

func (permission *Permission) TableName() string {
	return "permissions"
}
