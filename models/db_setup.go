package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
  
func ConnectDatabase() {
	dsn := "user=nenzin password=passwarps dbname=jampa_dev port=5432 sslmode=disable TimeZone=Asia/Thimphu"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  
	if err != nil {
	  panic(err)
	}
  
	database.AutoMigrate(&Campaign{})
  
	DB = database
}
