package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) error {
	fmt.Println("before create", time.Now())

	tx.Statement.SetColumn("CreatedAt", time.Now())
	tx.Statement.SetColumn("UpdatedAt", time.Now())
	return nil
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.SetColumn("UpdatedAt", time.Now())
	return nil
}
