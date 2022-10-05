package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `json:"Username" gorm:"index; not null"`
	Address     string `json:"Address"`
	Email       string `json:"Email"`
	PhoneNumber string `json:"PhoneNumber"`
}
