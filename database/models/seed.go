package model

import (
	"time"
)

type Seed struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Timestamp time.Time `gorm:"autoCreateTime:milli" json:"timestamp"`
}
