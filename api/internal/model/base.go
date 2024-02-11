package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}
