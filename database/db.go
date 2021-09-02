package database

import (
	"log"

	"backend.com/m/v2/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Setup() {
	dsn := "postgres://postgres@localhost/go_backend"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Person{}, &models.Job{})

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
