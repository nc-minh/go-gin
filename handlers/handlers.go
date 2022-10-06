package handlers

import (
	db "go-gin/databases"
	repoimpl "go-gin/repository/impl"

	"gorm.io/gorm"
)

var UserRepo = repoimpl.NewUserRepo(db.ConnnectPostgres().DB)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}
