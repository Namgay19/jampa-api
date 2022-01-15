package models

import (
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Model struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"` 
}
  
func ConnectDatabase() {
	dsn := "user=nenzin password=passwarps dbname=jampa_dev port=5432 sslmode=disable TimeZone=Asia/Thimphu"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
  
	if err != nil {
	  panic(err)
	}
  
	database.AutoMigrate(&Campaign{})
	database.AutoMigrate(&Image{})
	database.AutoMigrate(&Donation{})
  
	DB = database
}

func Paginate(page, pageSize string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(page)
		pageSize, _ := strconv.Atoi(pageSize)

		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 15:
			pageSize = 15
		case pageSize <= 10:
			pageSize = 10
		}
		
		offset := (page - 1) * pageSize
    	return db.Offset(offset).Limit(pageSize)
	}
}