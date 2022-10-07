package repository

import (
	"go-gin/models"
)

type BookRepo interface {
	FindAll() ([]*models.Book, error)
	FindById(id int) (*models.Book, error)
	Save(book *models.Book) (*models.Book, error)
	Update(book *models.Book) (*models.Book, error)
	Delete(book *models.Book) error
}
