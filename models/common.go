package models

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Thimphu", user, password, dbname, port)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Campaign{})
	database.AutoMigrate(&Image{})
	database.AutoMigrate(&Donation{})
	database.AutoMigrate(&User{})

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
		case pageSize > 10:
			pageSize = 10
		case pageSize <= 3:
			pageSize = 3
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
