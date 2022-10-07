package repoimpl

import (
	"go-gin/models"
	"go-gin/repository"

	"gorm.io/gorm"
)

type BookRepoImpl struct {
	DB *gorm.DB
}

func NewBookRepo(db *gorm.DB) repository.BookRepo {
	return &BookRepoImpl{DB: db}
}

func (b *BookRepoImpl) FindAll() ([]*models.Book, error) {
	var books []*models.Book
	err := b.DB.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (u *BookRepoImpl) FindById(id int) (*models.Book, error) {
	var book models.Book
	err := u.DB.First(&book, id).Error
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (b *BookRepoImpl) Save(book *models.Book) (*models.Book, error) {
	err := b.DB.Create(book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (b *BookRepoImpl) Update(book *models.Book) (*models.Book, error) {

	err := b.DB.Updates(book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (b *BookRepoImpl) Delete(book *models.Book) error {

	err := b.DB.Delete(book).Error
	if err != nil {
		return err
	}
	return nil
}

// Language: go
