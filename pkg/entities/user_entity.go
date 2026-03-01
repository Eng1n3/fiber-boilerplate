package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;not null" json:"id"`
	Username  string    `gorm:"not null" json:"username"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli" json:"-"`
	UpdatedAt time.Time `gorm:"autoCreateTime:milli;autoUpdateTime:milli" json:"-"`
	DeletedAt time.Time `gorm:"index" json:"-"`
}

func (user *User) BeforeCreate(_ *gorm.DB) error {
	user.ID = uuid.New() // Generate UUID before create
	return nil
}

func (user *User) TableName() string {
	return "users"
}
