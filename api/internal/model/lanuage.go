package model

import "time"

type Language struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Name string `json:"name" gorm:"type:varchar(30); not null;"`

	Accents []Accent `json:"accents"`

	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}