package model

import "time"

type Language struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Name     string `json:"name" gorm:"type:varchar(30); not null;"`
	Complete bool   `json:"complete" gorm:"default:false"`

	Surahs []Surah `json:"surahs" gorm:"many2many:language_surah"`

	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}
