package handlers

import (
	db "go-gin/databases"
	repoimpl "go-gin/repository/impl"

	"gorm.io/gorm"
)

var connected = db.ConnnectPostgres().DB

var UserRepo = repoimpl.NewUserRepo(connected)
var BookRepo = repoimpl.NewBookRepo(connected)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}
