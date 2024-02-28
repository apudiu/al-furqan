package model

import "time"

type Ayat struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Serial uint16 `json:"serial" gorm:"type:smallint; not null;index:idx_ayat_surah_accent,unique;"`
	Ruku   bool   `json:"ruku" gorm:"type:tinyint(1); not null;"`
	Sijdah bool   `json:"sijdah" gorm:"type:tinyint(1); not null;"`

	SurahID uint  `json:"surah_id" gorm:"not null;index:idx_ayat_surah_accent,unique;"`
	Surah   Surah `json:"surah"`

	AccentID uint   `json:"accent_id" gorm:"not null;index:idx_ayat_surah_accent,unique;"`
	Accent   Accent `json:"accent"`

	File  Media  `json:"file" gorm:"polymorphic:Owner;"`
	Start uint16 `json:"start" gorm:"not null;"`
	End   uint16 `json:"end" gorm:"not null;"`

	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}
