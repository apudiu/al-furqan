package model

type Media struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Name string `json:"name" gorm:"not null"`
	Path string `json:"path" gorm:"not null"`
	Type string `json:"type" gorm:"not null"`
	Size string `json:"size" gorm:"not null"`
	Cdn  bool   `json:"cdn" gorm:"default:true;"`

	OwnerID   uint   `json:"owner_id"`
	OwnerType string `json:"owner_type"`
}
