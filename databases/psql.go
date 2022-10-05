package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	models "go-gin/models"
)

func Init() *gorm.DB {
	dbURL := "postgres://mars:mars@localhost:5432/test"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.User{})

	return db
}
