package model

import "time"

type Surah struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Name string `json:"name" gorm:"type:varchar(30); not null"`

	Ayats []Ayat `json:"ayats"`

	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}
