package models

import (
	"errors"
	"log"
	hashEncrpt "safePasswordApi/src/security/encrypt/hash"
	"strings"

	"github.com/badoux/checkmail"
)

type UserLogin struct {
	Email          string `json:"email,omitempty" db:"email"`
	Password_Hash  string `json:"-" db:"safepassword,omitempty"`
	Password_Plain string `json:"password,omitempty" db:"-"`
	Email_Hash     string `json:"-" db:"email_hash"`
}

// Prepare: Format login data
func (user *UserLogin) Prepare() (err error) {
	if err = user.Format(); err != nil {
		return
	}

	if err = user.Validate(); err != nil {
		return
	}

	if err = user.HashEncrypt(); err != nil {
		return
	}

	return
}

// Validate: checks that login data is in a valid format
func (user *UserLogin) Validate() error {
	if len(strings.TrimSpace(user.Email)) == 0 {
		return errors.New("email is required and cannot be blank")
	}

	if err := checkmail.ValidateFormat(strings.TrimSpace(user.Email)); err != nil {
		return errors.New("invalid email format")
	}

	if len(strings.TrimSpace(user.Email)) == 0 {
		return errors.New("password is required and cannot be blank")
	}

	return nil
}

// Format: Format login data
func (user *UserLogin) Format() (err error) {
	user.Email = strings.TrimSpace(user.Email)
	return
}

// HashEncrypt: Encrypt hashed login data
func (user *UserLogin) HashEncrypt() (err error) {
	user.Password_Hash, err = hashEncrpt.GenerateSHA512(user.Password_Plain)
	if err != nil {
		log.Fatal(err)
	}

	user.Email_Hash, err = hashEncrpt.GenerateSHA512(user.Email)
	if err != nil {
		log.Fatal(err)
	}

	return
}
