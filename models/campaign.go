package models

import (
	"time"
)

type Model struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"` 
}

type Campaign struct {
	Model
	Header string `json:"header"`
	SubHeader string `json:"sub_header"` 
	Description string `json:"description"`
	Category string `json:"category"`
	EndDate time.Time `json:"end_date"`
}
