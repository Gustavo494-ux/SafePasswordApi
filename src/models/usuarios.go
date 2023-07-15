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

type Usuario struct {
	ID       uint64    `json:"id,omitempty" db:"id"`
	Nome     string    `json:"nome,omitempty" db:"nome"`
	Email    string    `json:"email,omitempty" db:"email"`
	Senha    string    `json:"senha,omitempty" db:"senha"`
	CriadoEm time.Time `json:"criadoEm,omitempty" db:"criadoem"`
}

// Prepare will call methods to validate and format the received user based on the given stage.
func (user *Usuario) Prepare(stage string) error {
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
func (user *Usuario) Validate(stage string) error {
	if user.Nome == "" {
		return errors.New("name is required and cannot be blank")
	}

	if user.Email == "" {
		return errors.New("email is required and cannot be blank")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("invalid email format")
	}

	if user.Senha == "" && stage == "signup" {
		return errors.New("password is required and cannot be blank")
	}

	return nil
}

// Format trims the leading and trailing spaces of the user's Name and Email fields and applies additional formatting based on the given stage.
func (user *Usuario) Format(stage string) error {
	user.Nome = strings.TrimSpace(user.Nome)
	user.Email = strings.TrimSpace(user.Email)

	switch stage {
	case "signup":
		{
			senhaHash, err := hashEncrpt.GenerateSHA512(user.Senha)
			if err != nil {
				return err
			}
			user.Senha = senhaHash
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
func (user *Usuario) GenerateSymmetricEncryptionKey() ([]byte, error) {
	idHash, erro := hashEncrpt.GenerateSHA512(strconv.FormatUint(user.ID, 10))
	if erro != nil {
		return []byte{}, erro
	}

	var senhaHash string
	if len(user.Senha) == 128 {
		senhaHash = user.Senha
	} else {
		senhaHash, erro = hashEncrpt.GenerateSHA512(user.Senha)
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
func (user *Usuario) EncryptAES() error {
	var err error
	if user.Nome, err = symmetricEncrypt.EncryptDataAES(user.Nome, configs.AESKey); err != nil {
		return err
	}

	if user.Email, err = symmetricEncrypt.EncryptDataAES(user.Email, configs.AESKey); err != nil {
		return err
	}

	return nil
}

// DecryptAES decrypts the user's Name and Email using AES decryption.
func (user *Usuario) DecryptAES() error {
	var err error
	if user.Nome, err = symmetricEncrypt.DecryptDataAES(user.Nome, configs.AESKey); err != nil {
		return err
	}

	if user.Email, err = symmetricEncrypt.DecryptDataAES(user.Email, configs.AESKey); err != nil {
		return err
	}

	return nil
}

// EncryptRSA encrypts the user's Name and Email using RSA encryption.
func (user *Usuario) EncryptRSA() error {
	var err error
	publicKey, err := asymmetrical.ParseRSAPublicKey(configs.RSAPublicKey)
	if err != nil {
		return err
	}

	if user.Nome, err = asymmetrical.EncryptRSA(user.Nome, publicKey); err != nil {
		return err
	}

	if user.Email, err = asymmetrical.EncryptRSA(user.Email, publicKey); err != nil {
		return err
	}
	return nil
}

// DecryptRSA decrypts the user's Name and Email using RSA decryption.
func (user *Usuario) DecryptRSA() error {
	var err error
	privateKey, err := asymmetrical.ParseRSAPrivateKey(configs.RSAPrivateKey)
	if err != nil {
		return err
	}

	if user.Nome, err = asymmetrical.DecryptRSA(user.Nome, privateKey); err != nil {
		return err
	}

	if user.Email, err = asymmetrical.DecryptRSA(user.Email, privateKey); err != nil {
		return err
	}
	return nil
}

// Encrypt encrypts the user's data using both AES and RSA encryption.
func (user *Usuario) Encrypt() error {
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
func (user *Usuario) Decrypt() error {
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
