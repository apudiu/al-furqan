package model

import "time"

type Accent struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Name string `json:"name" gorm:"type:varchar(30); not null;index:idx_accent_language,unique;"`

	LanguageID uint     `json:"language_id" gorm:"not null;index:idx_accent_language,unique;"`
	Language   Language `json:"language"`

	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}
