package models

import (
	"errors"
	// "fmt"
	// "safePasswordApi/src/configs"
	// "safePasswordApi/src/security/encrypt/asymmetrical"
	// hashEncrpt "safePasswordApi/src/security/encrypt/hash"
	// symmetricEncrypt "safePasswordApi/src/security/encrypt/symmetrical"
	// "strconv"
	//"strings"

	"github.com/badoux/checkmail"
)

type UserRegister struct {
	Name     string `json:"name,omitempty" db:"name"`
	Email    string `json:"email,omitempty" db:"email"`
	Password string `json:"password,omitempty" db:"safepassword,omitempty"`
}

// Prepare will call methods to validate and format the received user based on the given stage.
func (user UserRegister) Prepare(stage string) error {
	if err := user.Validate(stage); err != nil {
		return err
	}
	return nil
}

// Validate checks if the user fields are valid based on the given stage.
func (user UserRegister) Validate(stage string) error {
	if user.Name == "" {
		return errors.New("name is required and cannot be blank")
	}

	if user.Email == "" {
		return errors.New("email is required and cannot be blank")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("invalid email format")
	}

	if user.Password == "" {
		return errors.New("password is required and cannot be blank")
	}

	return nil
}
