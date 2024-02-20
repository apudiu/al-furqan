package model

import "time"

type AudioFile struct {
	ID uint `json:"id" gorm:"primaryKey"`

	FilePath string `json:"file_path"`

	AyatID   uint   `json:"ayat_id"`
	Ayat     Ayat   `json:"ayat"`
	AccentID uint   `json:"accent_id"`
	Accent   Accent `json:"accent"`

	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}
