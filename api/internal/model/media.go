package model

type Media struct {
	baseModel
	Name string `json:"name" gorm:"not null"`
	Path string `json:"path" gorm:"not null"`
}
