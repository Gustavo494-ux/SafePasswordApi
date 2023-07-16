package models_test

import (
	"safePasswordApi/src/configs"
	"safePasswordApi/src/models"
	"testing"
	"time"
)

var UserPath_DotEnv = "./../../.env"
var usuarioCredentials = []models.User{
	{
		ID:         1,
		Name:       "User 1",
		Email:      "usuario1@example.com",
		Password:   "password1",
		Created_at: time.Now(),
	},
	{
		ID:         2,
		Name:       "User 2",
		Email:      "usuario2@example.com",
		Password:   "password2",
		Created_at: time.Now(),
	},
	{
		ID:         3,
		Name:       "User 3",
		Email:      "usuario3@example.com",
		Password:   "password3",
		Created_at: time.Now(),
	},
}

func TestUser_Preparar(t *testing.T) {
	configs.InitializeConfigurations(UserPath_DotEnv)
	for _, usuario := range usuarioCredentials {
		err := usuario.Prepare("signup")
		if err != nil {
			t.Fatalf("An error occurred while preparing the usuario for registration: %v", err)
		}

		err = usuario.Prepare("query")
		if err != nil {
			t.Fatalf("An error occurred while preparing the usuario for query: %v", err)
		}

	}
}

func TestUser_Validar(t *testing.T) {
	for _, usuario := range usuarioCredentials {
		err := usuario.Validate("signup")

		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}
	}
}

func TestUser_Validar_EmptyName(t *testing.T) {
	var err error
	for _, usuario := range usuarioCredentials {
		usuario.Name = ""
		err = usuario.Validate("signup")
		if err == nil {
			t.Error("Expected an error, but none returned")
		} else if err.Error() != "name is required and cannot be blank" {
			t.Errorf("Expected error: %s", "name is required and cannot be blank")
		}
	}
}

func TestUser_Validar_InvalidEmail(t *testing.T) {
	var err error
	for _, usuario := range usuarioCredentials {
		usuario.Email = "invalidemail"
		err = usuario.Validate("signup")
		if err == nil {
			t.Error("Expected an error, but none returned")
		} else if err.Error() != "invalid email format" {
			t.Errorf("Expected error: %s", "invalid email address")
		}
	}
}

func TestUser_Validar_EmptyPassword(t *testing.T) {
	var err error
	for _, usuario := range usuarioCredentials {
		usuario.Password = ""
		err = usuario.Validate("signup")
		if err != nil && err.Error() != "password is required and cannot be blank" {
			t.Errorf("Unexpected error: %s", err.Error())
		}
	}
}

func TestUser_Formatar_Signup(t *testing.T) {
	configs.InitializeConfigurations(UserPath_DotEnv)
	var err error
	for _, usuario := range usuarioCredentials {
		err = usuario.Format("signup")

		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}
	}
}

func TestUser_Encrypt(t *testing.T) {
	configs.InitializeConfigurations(UserPath_DotEnv)
	var err error
	for _, usuario := range usuarioCredentials {
		err = usuario.Encrypt()
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}
	}
}

func TestUser_Decrypt(t *testing.T) {
	configs.InitializeConfigurations(UserPath_DotEnv)
	var err error
	for _, usuario := range usuarioCredentials {
		err = usuario.Decrypt()
		if err != nil && err.Error() != "data not encrypted with RSA" {
			t.Errorf("Unexpected error: %s", err.Error())
		}
	}
}

func TestUser_EncryptAES(t *testing.T) {
	configs.InitializeConfigurations(UserPath_DotEnv)
	var err error
	for _, usuario := range usuarioCredentials {
		err = usuario.EncryptAES()
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}
	}
}

func TestUser_DecryptAES(t *testing.T) {
	configs.InitializeConfigurations(UserPath_DotEnv)
	var err error
	for _, usuario := range usuarioCredentials {
		err = usuario.DecryptAES()
		if err != nil && err.Error() != "unencrypted data using AES 256" {
			t.Errorf("Unexpected error: %s", err.Error())
		}
	}
}

func TestUser_EncryptDecryptAES(t *testing.T) {
	configs.InitializeConfigurations(UserPath_DotEnv)
	var err error

	// Encrypt usuarios
	usuariosToEncrypt := usuarioCredentials

	for i := range usuariosToEncrypt {
		err = usuariosToEncrypt[i].EncryptAES()
		if err != nil {
			t.Errorf("Unexpected error while encrypting: %s", err.Error())
		}
	}

	//Decrypt usuarios
	for i := range usuariosToEncrypt {
		err = usuariosToEncrypt[i].DecryptAES()
		if err != nil {
			t.Errorf("Unexpected error while decrypting: %s", err.Error())
		}
	}

	//Compare original usuarios with decrypted ones
	for i := range usuarioCredentials {
		if usuariosToEncrypt[i].Name != usuarioCredentials[i].Name ||
			usuariosToEncrypt[i].Email != usuarioCredentials[i].Email ||
			usuariosToEncrypt[i].Password != usuarioCredentials[i].Password {
			t.Errorf("Decrypted usuario does not match the original")
		}
	}
}

func TestUser_EncryptRSA(t *testing.T) {
	configs.InitializeConfigurations(UserPath_DotEnv)
	var err error
	for _, usuario := range usuarioCredentials {
		err = usuario.EncryptRSA()
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}
	}
}

func TestUser_DecryptRSA(t *testing.T) {
	configs.InitializeConfigurations(UserPath_DotEnv)
	var err error
	for _, usuario := range usuarioCredentials {
		err = usuario.DecryptRSA()
		if err != nil && err.Error() != "data not encrypted with RSA" {
			t.Errorf("Unexpected error: %s", err.Error())
		}
	}
}

func TestUser_EncryptDecryptRSA(t *testing.T) {
	configs.InitializeConfigurations(UserPath_DotEnv)
	var err error

	// Encrypt usuarios
	usuariosToEncrypt := usuarioCredentials

	for i := range usuariosToEncrypt {
		err = usuariosToEncrypt[i].EncryptRSA()
		if err != nil {
			t.Errorf("Unexpected error while encrypting: %s", err.Error())
		}
	}

	// Decrypt usuarios
	for i := range usuariosToEncrypt {
		err = usuariosToEncrypt[i].DecryptRSA()
		if err != nil {
			t.Errorf("Unexpected error while decrypting: %s", err.Error())
		}
	}

	// Compare original usuarios with decrypted ones
	for i := range usuarioCredentials {
		if usuariosToEncrypt[i].Name != usuarioCredentials[i].Name ||
			usuariosToEncrypt[i].Email != usuarioCredentials[i].Email ||
			usuariosToEncrypt[i].Password != usuarioCredentials[i].Password {
			t.Errorf("Decrypted usuario does not match the original")
		}
	}
}
