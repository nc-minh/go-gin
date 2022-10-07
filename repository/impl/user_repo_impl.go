package repoimpl

import (
	"go-gin/models"
	"go-gin/repository"

	"gorm.io/gorm"
)

type UserRepoImpl struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) repository.UserRepo {
	return &UserRepoImpl{DB: db}
}

func (u *UserRepoImpl) FindAll() ([]*models.User, error) {
	var users []*models.User
	err := u.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepoImpl) FindById(id int) (*models.User, error) {
	var user models.User
	err := u.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepoImpl) Save(user *models.User) (*models.User, error) {
	err := u.DB.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepoImpl) Update(user *models.User) (*models.User, error) {

	err := u.DB.Updates(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepoImpl) Delete(user *models.User) error {

	err := u.DB.Delete(user).Error
	if err != nil {
		return err
	}
	return nil
}

// Language: go
