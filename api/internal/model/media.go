package model

type Media struct {
	BaseModel
	Name string `json:"name" gorm:"not null"`
	Path string `json:"path" gorm:"not null"`
}
