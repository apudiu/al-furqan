package model

import "time"

type Surah struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Serial uint16 `json:"serial" gorm:"type:smallint; not null;"`
	Name   string `json:"name" gorm:"type:varchar(30); not null"`

	TotalAyats uint16 `json:"total_ayats" gorm:"type:smallint(3); not null;"`
	Ayats      []Ayat `json:"ayats"`

	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}
