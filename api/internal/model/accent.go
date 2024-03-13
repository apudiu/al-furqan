package model

import "time"

type Accent struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Name string `json:"name"`

	LanguageID uint `json:"language_id"`
	Language   Ayat `json:"language"`

	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}
