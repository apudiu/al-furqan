package model

import "time"

type Ayat struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Serial uint16 `json:"serial" gorm:"type:smallint; not null;"`
	Ruku   uint8  `json:"ruku" gorm:"type:tinyint(1); not null;"`
	Sijdah uint8  `json:"sijdah" gorm:"type:tinyint(1); not null;"`

	SurahID   uint       `json:"surah_id"`
	Surah     Surah      `json:"surah"`
	Languages []Language `json:"languages" gorm:"many2many:language_surah"`

	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}
