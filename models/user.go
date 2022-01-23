package models

type User struct {
	Model
	FirstName string `json:"first_name" gorm:"not null"`
	LastName string `json:"last_name"`
	Email string `json:"email" gorm:"not null"`
	Password string `json:"-" gorm:"not null"`
	Phone string `json:"phone"`
}
