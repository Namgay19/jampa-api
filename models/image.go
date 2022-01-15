package models

type Image struct {
	Model
	ImageUrl string `json:"image_url" gorm:"not null"`
	OwnerID int `json:"owner_id"`
	OwnerType string `json:"owner_type"`
}