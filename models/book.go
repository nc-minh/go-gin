package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"Title" gorm:"index; not null"`
	Author string `json:"Author" gorm:"index; not null"`
	Desc   string `json:"Desc"`
}
