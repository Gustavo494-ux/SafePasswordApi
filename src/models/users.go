package models

import (
	"errors"
	"fmt"
	"safePasswordApi/src/configs"
	"safePasswordApi/src/security/encrypt/asymmetrical"
	hashEncrpt "safePasswordApi/src/security/encrypt/hash"
	symmetricEncrypt "safePasswordApi/src/security/encrypt/symmetrical"
	"strconv"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID         uint64    `json:"id,omitempty" db:"id"`
	Name       string    `json:"name,omitempty" db:"name"`
	Email      string    `json:"email,omitempty" db:"email"`
	Email_Hash string    `json:"email_hash,omitempty" db:"email_hash"`
	Password   string    `json:"password,omitempty" db:"safepassword,omitempty"`
	Created_at time.Time `json:"created_at,omitempty" db:"created_at"`
}

// Prepare will call methods to validate and format the received user based on the given stage.
func (user *User) Prepare(stage string) error {
	switch stage {
	case "signup":
		if err := user.Validate(stage); err != nil {
			return err
		}
	}

	if err := user.Format(stage); err != nil {
		return err
	}

	return nil
}

// Validate checks if the user fields are valid based on the given stage.
func (user *User) Validate(stage string) error {
	if user.Name == "" {
		return errors.New("name is required and cannot be blank")
	}

	if user.Email == "" {
		return errors.New("email is required and cannot be blank")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("invalid email format")
	}

	if user.Password == "" && stage == "signup" {
		return errors.New("password is required and cannot be blank")
	}

	return nil
}

// Format trims the leading and trailing spaces of the user's Name and Email fields and applies additional formatting based on the given stage.
func (user *User) Format(stage string) error {
	var err error
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)

	switch stage {
	case "signup":
		{
			user.Password, err = hashEncrpt.GenerateSHA512(user.Password)
			if err != nil {
				return err
			}

			user.Email_Hash, err = hashEncrpt.GenerateSHA512(user.Email)
			if err != nil {
				return err
			}

			if err = user.Encrypt(); err != nil {
				return err
			}
		}
	case "query":
		if err := user.Decrypt(); err != nil {
			return err
		}
	}

	return nil
}

// GenerateSymmetricEncryptionKey generates a symmetric encryption key based on the user's ID and Password.
func (user *User) GenerateSymmetricEncryptionKey() ([]byte, error) {
	idHash, erro := hashEncrpt.GenerateSHA512(strconv.FormatUint(user.ID, 10))
	if erro != nil {
		return []byte{}, erro
	}

	var senhaHash string
	if len(user.Password) == 128 {
		senhaHash = user.Password
	} else {
		senhaHash, erro = hashEncrpt.GenerateSHA512(user.Password)
		if erro != nil {
			return []byte{}, erro
		}
	}

	chaveDeCodificacao, erro := hashEncrpt.GenerateSHA512(fmt.Sprintf("%s%d%s", idHash, user.ID, senhaHash))
	if erro != nil {
		return []byte{}, erro
	}
	return []byte(chaveDeCodificacao), nil
}

// EncryptAES encrypts the user's Name and Email using AES encryption.
func (user *User) EncryptAES() error {
	var err error
	if user.Name, err = symmetricEncrypt.EncryptDataAES(user.Name, configs.AESKey); err != nil {
		return err
	}

	if user.Email, err = symmetricEncrypt.EncryptDataAES(user.Email, configs.AESKey); err != nil {
		return err
	}

	return nil
}

// DecryptAES decrypts the user's Name and Email using AES decryption.
func (user *User) DecryptAES() error {
	var err error
	if user.Name, err = symmetricEncrypt.DecryptDataAES(user.Name, configs.AESKey); err != nil {
		return err
	}

	if user.Email, err = symmetricEncrypt.DecryptDataAES(user.Email, configs.AESKey); err != nil {
		return err
	}

	return nil
}

// EncryptRSA encrypts the user's Name and Email using RSA encryption.
func (user *User) EncryptRSA() error {
	var err error
	publicKey, err := asymmetrical.ParseRSAPublicKey(configs.RSAPublicKey)
	if err != nil {
		return err
	}

	if user.Name, err = asymmetrical.EncryptRSA(user.Name, publicKey); err != nil {
		return err
	}

	if user.Email, err = asymmetrical.EncryptRSA(user.Email, publicKey); err != nil {
		return err
	}
	return nil
}

// DecryptRSA decrypts the user's Name and Email using RSA decryption.
func (user *User) DecryptRSA() error {
	var err error
	privateKey, err := asymmetrical.ParseRSAPrivateKey(configs.RSAPrivateKey)
	if err != nil {
		return err
	}

	if user.Name, err = asymmetrical.DecryptRSA(user.Name, privateKey); err != nil {
		return err
	}

	if user.Email, err = asymmetrical.DecryptRSA(user.Email, privateKey); err != nil {
		return err
	}
	return nil
}

// Encrypt encrypts the user's data using both AES and RSA encryption.
func (user *User) Encrypt() error {
	err := user.EncryptAES()
	if err != nil {
		return err
	}

	err = user.EncryptRSA()
	if err != nil {
		return err
	}

	return nil
}

// Decrypt decrypts the user's data using both AES and RSA decryption.
func (user *User) Decrypt() error {
	err := user.DecryptRSA()
	if err != nil {
		return err
	}

	err = user.DecryptAES()
	if err != nil {
		return err
	}

	return nil
}
