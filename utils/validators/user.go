package validators

import (
	"errors"
	models "go-gin/models"
)

func Validate(user *models.User) error {
	if user.Email == "" {

		err := errors.New("email is required")
		return err
	}

	if user.Address == "" {
		err := errors.New("address is required")
		return err
	}

	if user.PhoneNumber == "" {
		err := errors.New("phone number is required")
		return err
	}

	if user.Username == "" {
		err := errors.New("username is required")
		return err
	}

	if len(user.Username) < 6 {
		err := errors.New("username must be at least 6 characters")
		return err
	}

	return nil
}
