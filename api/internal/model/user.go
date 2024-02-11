package model

import "time"

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Name     string `json:"name" gorm:"type:varchar(30); not null"`
	Email    string `json:"email" gorm:"type:varchar(40); not null"`
	Password string `json:"password" gorm:"type:varchar(100); not null"`

	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}
