package db

import (
	"log"

	"github.com/fatih/color"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	models "go-gin/models"
)

type Postgres struct {
	*gorm.DB
}

var Psql = &Postgres{}

func ConnnectPostgres() *Postgres {
	dbURL := "postgres://mars:mars@localhost:5432/test"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.User{})

	Psql.DB = db

	color.Green("Database connected")
	return Psql
}
