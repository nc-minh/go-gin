package repository

import (
	"go-gin/models"
)

type UserRepo interface {
	FindAll() ([]*models.User, error)
	FindById(id int) (*models.User, error)
	Save(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(user *models.User) error
}
