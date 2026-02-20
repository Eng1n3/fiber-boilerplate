package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID          uuid.UUID `gorm:"primaryKey;not null" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Email       string    `gorm:"uniqueIndex;not null" json:"email"`
	Description string    `gorm:"null" json:"-"`
	// VerifiedEmail bool      `gorm:"default:false;not null" json:"verified_email"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli" json:"-"`
	UpdatedAt time.Time `gorm:"autoCreateTime:milli;autoUpdateTime:milli" json:"-"`
	DeletedAt time.Time `gorm:"index" json:"-"`
	// Token         []Token   `gorm:"foreignKey:role_id;references:id" json:"-"`
}

func (role *Role) BeforeCreate(_ *gorm.DB) error {
	role.ID = uuid.New() // Generate UUID before create
	return nil
}

func (role *Role) TableName() string {
	return "roles"
}
