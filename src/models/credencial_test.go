package models_test

import (
	"safePasswordApi/src/configs"
	enum "safePasswordApi/src/enum/geral"
	"safePasswordApi/src/models"
	"testing"
	"time"
)

var credentials = []models.Credencial{
	{
		Id:        1,
		UsuarioId: 1,
		Descricao: "Credential 1",
		SiteUrl:   "https://www.example.com",
		Login:     "user1",
		Senha:     "password1",
		CriadoEm:  time.Now(),
	},
	{
		Id:        2,
		UsuarioId: 1,
		Descricao: "Credential 2",
		SiteUrl:   "https://www.example.com",
		Login:     "user2",
		Senha:     "password2",
		CriadoEm:  time.Now(),
	},
	{
		Id:        3,
		UsuarioId: 2,
		Descricao: "Credential 3",
		SiteUrl:   "https://www.example.com",
		Login:     "user3",
		Senha:     "password3",
		CriadoEm:  time.Now(),
	},
	{
		Id:        4,
		UsuarioId: 2,
		Descricao: "Credential 4",
		SiteUrl:   "https://www.example.com",
		Login:     "user4",
		Senha:     "password4",
		CriadoEm:  time.Now(),
	},
	{
		Id:        5,
		UsuarioId: 3,
		Descricao: "Credential 5",
		SiteUrl:   "https://www.example.com",
		Login:     "user5",
		Senha:     "password5",
		CriadoEm:  time.Now(),
	},
	{
		Id:        6,
		UsuarioId: 3,
		Descricao: "Credential 6",
		SiteUrl:   "https://www.example.com",
		Login:     "user6",
		Senha:     "password6",
		CriadoEm:  time.Now(),
	},
	{
		Id:        7,
		UsuarioId: 4,
		Descricao: "Credential 7",
		SiteUrl:   "https://www.example.com",
		Login:     "user7",
		Senha:     "password7",
		CriadoEm:  time.Now(),
	},
	{
		Id:        8,
		UsuarioId: 4,
		Descricao: "Credential 8",
		SiteUrl:   "https://www.example.com",
		Login:     "user8",
		Senha:     "password8",
		CriadoEm:  time.Now(),
	},
	{
		Id:        9,
		UsuarioId: 5,
		Descricao: "Credential 9",
		SiteUrl:   "https://www.example.com",
		Login:     "user9",
		Senha:     "password9",
		CriadoEm:  time.Now(),
	},
	{
		Id:        10,
		UsuarioId: 5,
		Descricao: "Credential 10",
		SiteUrl:   "https://www.example.com",
		Login:     "user10",
		Senha:     "password10",
		CriadoEm:  time.Now(),
	},
}

func TestCredencial_Prepare(t *testing.T) {
	configs.InitializeConfigurations()
	for _, credential := range credentials {
		err := credential.Preparar(enum.TipoPreparacao_Cadastro)
		if err != nil {
			t.Fatalf("An error occurred while preparing the credential: %v", err)
		}
	}
}

func TestValidate(t *testing.T) {
	for _, credential := range credentials {
		err := credential.Validar()

		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}
	}
}

func TestValidate_UsuarioIdZero(t *testing.T) {
	var err error
	for _, credential := range credentials {
		if credential.Id%2 == 1 {
			credential.UsuarioId = 1
		} else {
			credential.UsuarioId = 0
		}

		err = credential.Validar()
		if err == nil {
			if credential.Id%2 == 0 {
				t.Error("Esperava o erro 'o usuário é obrigatório e não pode ficar em branco', mas nenhum retornou")
			}
		} else if err.Error() != "o usuário é obrigatório e não pode ficar em branco" {
			t.Errorf("Erro esperado: %s", "o usuário é obrigatório e não pode ficar em branco")
		}
	}
}

func TestValidate_SenhaVazia(t *testing.T) {
	var err error
	for _, credential := range credentials {
		if credential.Id%2 == 1 {
			credential.Senha = ""
		} else {
			credential.Senha = "Teste"
		}
		err = credential.Validar()
		if err == nil {
			if credential.Id%2 == 1 {
				t.Error("Esperava o erro 'a senha é obrigatória e não pode ficar em branco', mas nenhum retornou")
			}
		} else if err.Error() != "a senha é obrigatória e não pode ficar em branco" {
			t.Errorf("Erro esperado: %s", "a senha é obrigatória e não pode ficar em branco")
		}
	}
}

func TestFormat_SaveData(t *testing.T) {
	configs.InitializeConfigurations()
	var err error
	for _, credential := range credentials {
		err = credential.Formatar(enum.TipoFormatacao_Cadastro)

		if err != nil {
			t.Errorf("Erro inesperado: %s", err.Error())
		}
	}
}

func TestFormat_RetrieveData(t *testing.T) {
	configs.InitializeConfigurations()
	var err error
	for _, credential := range credentials {
		err = credential.Preparar(enum.TipoPreparacao_Consulta)

		if err.Error() != "data not encrypted with RSA" {
			t.Errorf("Error: %v", err)
		}
	}
}

func TestEncryptAES(t *testing.T) {
	configs.InitializeConfigurations()
	var err error
	credentialsWithEmptyPassword := credentials
	for _, credential := range credentialsWithEmptyPassword {
		err = credential.CriptografarAES()
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}
	}
}

func TestDecryptAES(t *testing.T) {
	configs.InitializeConfigurations()
	var err error
	for _, credential := range credentials {
		err = credential.DescriptografarAES()
		if err != nil && err.Error() != "unencrypted data using AES 256" {
			t.Errorf("Unexpected error: %s", err.Error())
		}
	}
}

func TestEncryptDecryptAES(t *testing.T) {
	configs.InitializeConfigurations()
	var err error

	// Encrypt credentials
	credentialsToEncrypt := credentials

	for i := range credentialsToEncrypt {
		err = credentialsToEncrypt[i].CriptografarAES()
		if err != nil {
			t.Errorf("Unexpected error while encrypting: %s", err.Error())
		}
	}

	// Decrypt credentials
	for i := range credentialsToEncrypt {
		err = credentialsToEncrypt[i].DescriptografarAES()
		if err != nil {
			t.Errorf("Unexpected error while decrypting: %s", err.Error())
		}
	}

	// Compare original credentials with decrypted ones
	for i := range credentials {
		if credentialsToEncrypt[i].Descricao != credentials[i].Descricao ||
			credentialsToEncrypt[i].SiteUrl != credentials[i].SiteUrl ||
			credentialsToEncrypt[i].Login != credentials[i].Login ||
			credentialsToEncrypt[i].Senha != credentials[i].Senha {
			t.Errorf("Decrypted credential does not match the original")
		}
	}
}

func TestEncryptRSA(t *testing.T) {
	configs.InitializeConfigurations()
	var err error
	for _, credential := range credentials {
		err = credential.CriptografarRSA()
		if err != nil {
			t.Errorf("Unexpected error: %s", err.Error())
		}
	}
}

func TestDecryptRSA(t *testing.T) {
	configs.InitializeConfigurations()
	var err error
	for _, credential := range credentials {
		err = credential.DescriptografarRSA()
		if err != nil && err.Error() != "data not encrypted with RSA" {
			t.Errorf("Unexpected error: %s", err.Error())
		}
	}
}

func TestEncryptDecryptRSA(t *testing.T) {
	configs.InitializeConfigurations()
	var err error

	// Encrypt credentials
	credentialsToEncrypt := credentials

	for i := range credentialsToEncrypt {
		err = credentialsToEncrypt[i].CriptografarRSA()
		if err != nil {
			t.Errorf("Unexpected error while encrypting: %s", err.Error())
		}
	}

	// Decrypt credentials
	for i := range credentialsToEncrypt {
		err = credentialsToEncrypt[i].DescriptografarRSA()
		if err != nil {
			t.Errorf("Unexpected error while decrypting: %s", err.Error())
		}
	}

	// Compare original credentials with decrypted ones
	for i := range credentials {
		if credentialsToEncrypt[i].Descricao != credentials[i].Descricao ||
			credentialsToEncrypt[i].SiteUrl != credentials[i].SiteUrl ||
			credentialsToEncrypt[i].Login != credentials[i].Login ||
			credentialsToEncrypt[i].Senha != credentials[i].Senha {
			t.Errorf("Decrypted credential does not match the original")
		}
	}
}

/*
package models_test

import (
	"safePasswordApi/src/configs"
	"safePasswordApi/src/models"
	"testing"
	"time"
)

var  = "./../../.env"
var credenciais = []models.Credencial{
	{
		Id:        1,
		UsuarioId: 1,
		Descricao: "Credencial 1",
		SiteUrl:   "https://www.example.com",
		Login:     "user1",
		Senha:     "senha1",
		CriadoEm:  time.Now(),
	},
	{
		Id:        2,
		UsuarioId: 1,
		Descricao: "Credencial 2",
		SiteUrl:   "https://www.example.com",
		Login:     "user2",
		Senha:     "senha2",
		CriadoEm:  time.Now(),
	},
	{
		Id:        3,
		UsuarioId: 2,
		Descricao: "Credencial 3",
		SiteUrl:   "https://www.example.com",
		Login:     "user3",
		Senha:     "senha3",
		CriadoEm:  time.Now(),
	},
	{
		Id:        4,
		UsuarioId: 2,
		Descricao: "Credencial 4",
		SiteUrl:   "https://www.example.com",
		Login:     "user4",
		Senha:     "senha4",
		CriadoEm:  time.Now(),
	},
	{
		Id:        5,
		UsuarioId: 3,
		Descricao: "Credencial 5",
		SiteUrl:   "https://www.example.com",
		Login:     "user5",
		Senha:     "senha5",
		CriadoEm:  time.Now(),
	},
	{
		Id:        6,
		UsuarioId: 3,
		Descricao: "Credencial 6",
		SiteUrl:   "https://www.example.com",
		Login:     "user6",
		Senha:     "senha6",
		CriadoEm:  time.Now(),
	},
	{
		Id:        7,
		UsuarioId: 4,
		Descricao: "Credencial 7",
		SiteUrl:   "https://www.example.com",
		Login:     "user7",
		Senha:     "senha7",
		CriadoEm:  time.Now(),
	},
	{
		Id:        8,
		UsuarioId: 4,
		Descricao: "Credencial 8",
		SiteUrl:   "https://www.example.com",
		Login:     "user8",
		Senha:     "senha8",
		CriadoEm:  time.Now(),
	},
	{
		Id:        9,
		UsuarioId: 5,
		Descricao: "Credencial 9",
		SiteUrl:   "https://www.example.com",
		Login:     "user9",
		Senha:     "senha9",
		CriadoEm:  time.Now(),
	},
	{
		Id:        10,
		UsuarioId: 5,
		Descricao: "Credencial 10",
		SiteUrl:   "https://www.example.com",
		Login:     "user10",
		Senha:     "senha10",
		CriadoEm:  time.Now(),
	},
}

func TestCredencial_Preparar(t *testing.T) {
	configs.InitializeConfigurations()
	for _, credencial := range credenciais {
		err := credencial.Prepare("salvarDados", "")
		if err != nil {
			t.Fatalf("ocorreu um erro ao realizar a preparação da credencial, error: %v", err)
		}
	}
}

func TestValidar(t *testing.T) {
	for _, credencial := range credenciais {
		err := credencial.Validate()

		if err != nil {
			t.Errorf("Erro inesperado: %s", err.Error())
		}
	}
}

func TestValidar_UsuarioIdZero(t *testing.T) {
	var err error
	for _, credencial := range credenciais {
		if credencial.Id%2 == 1 {
			credencial.UsuarioId = 1
		} else {
			credencial.UsuarioId = 0
		}

		err = credencial.Validate()
		if err == nil {
			if credencial.Id%2 == 0 {
				t.Error("Esperava-se um erro, mas nenhum foi retornado")
			}
		} else if err.Error() != "usuário é obrigatório e não pode estar em branco" {
			t.Errorf("Erro esperado: %s", "usuário é obrigatório e não pode estar em branco")
		}
	}
}

func TestValidar_SenhaVazia(t *testing.T) {
	var err error
	for _, credencial := range credenciais {
		if credencial.Id%2 == 1 {
			credencial.Senha = ""
		} else {
			credencial.Senha = "Teste"
		}
		err = credencial.Validate()
		if err == nil {
			if credencial.Id%2 == 1 {
				t.Error("Esperava-se um erro, mas nenhum foi retornado")
			}
		} else if err.Error() != "a senha é obrigatória e não pode estar em branco" {
			t.Errorf("Erro esperado: %s", "a senha é obrigatória e não pode estar em branco")
		}
	}
}

func TestFormatar_SalvarDados(t *testing.T) {
	configs.InitializeConfigurations()
	var err error
	for _, credencial := range credenciais {
		err = credencial.Format("salvarDados")

		if err != nil {
			t.Errorf("Erro inesperado: %s", err.Error())
		}
	}
}

func TestFormatar_ConsultarDados(t *testing.T) {
	configs.InitializeConfigurations()
	var err error
	for _, credencial := range credenciais {
		err = credencial.Format("consultarDados")

		if err.Error() != "data not encrypted with RSA" {
			t.Errorf("error: %v", err)
		}
	}
}

func TestCriptografarAES(t *testing.T) {
	configs.InitializeConfigurations()
	var err error
	credenciaisSenhaVazia := credenciais
	for _, credencial := range credenciaisSenhaVazia {
		err = credencial.EncryptAES()
		if err != nil {
			t.Errorf("Erro inesperado: %s", err.Error())
		}
	}
}

func TestDescriptografarAES(t *testing.T) {
	configs.InitializeConfigurations()
	var err error
	for _, credencial := range credenciais {
		err = credencial.DecryptAES()
		if err != nil && err.Error() != "unencrypted data using AES 256" {
			t.Errorf("Erro inesperado: %s", err.Error())
		}
	}
}

func TestEncryptDecryptAES(t *testing.T) {
	configs.InitializeConfigurations()
	var err error

	// Criptografar as credenciais
	credenciaisParaCriptografar := credenciais

	for i := range credenciaisParaCriptografar {
		err = credenciaisParaCriptografar[i].EncryptAES()
		if err != nil {
			t.Errorf("Erro inesperado ao criptografar: %s", err.Error())
		}
	}

	// Descriptografar as credenciais
	for i := range credenciaisParaCriptografar {
		err = credenciaisParaCriptografar[i].DecryptAES()
		if err != nil {
			t.Errorf("Erro inesperado ao descriptografar: %s", err.Error())
		}
	}

	// Comparar as credenciais originais com as descriptografadas
	for i := range credenciais {
		if credenciaisParaCriptografar[i].Descricao != credenciais[i].Descricao ||
			credenciaisParaCriptografar[i].SiteUrl != credenciais[i].SiteUrl ||
			credenciaisParaCriptografar[i].Login != credenciais[i].Login ||
			credenciaisParaCriptografar[i].Senha != credenciais[i].Senha {
			t.Errorf("Credencial descriptografada não corresponde à original")
		}
	}
}

func TestCriptografarRSA(t *testing.T) {
	configs.InitializeConfigurations()
	var err error
	for _, credencial := range credenciais {
		err = credencial.EncryptRSA()
		if err != nil {
			t.Errorf("Erro inesperado: %s", err.Error())
		}
	}
}

func TestDescriptografarRSA(t *testing.T) {
	configs.InitializeConfigurations()
	var err error
	for _, credencial := range credenciais {
		err = credencial.DecryptRSA()
		if err != nil && err.Error() != "data not encrypted with RSA" {
			t.Errorf("Erro inesperado: %s", err.Error())
		}
	}
}

func TestEncryptDecryptRSA(t *testing.T) {
	configs.InitializeConfigurations()
	var err error

	// Criptografar as credenciais
	credenciaisParaCriptografar := credenciais

	for i := range credenciaisParaCriptografar {
		err = credenciaisParaCriptografar[i].EncryptRSA()
		if err != nil {
			t.Errorf("Erro inesperado ao criptografar: %s", err.Error())
		}
	}

	// Descriptografar as credenciais
	for i := range credenciaisParaCriptografar {
		err = credenciaisParaCriptografar[i].DecryptRSA()
		if err != nil {
			t.Errorf("Erro inesperado ao descriptografar: %s", err.Error())
		}
	}

	// Comparar as credenciais originais com as descriptografadas
	for i := range credenciais {
		if credenciaisParaCriptografar[i].Descricao != credenciais[i].Descricao ||
			credenciaisParaCriptografar[i].SiteUrl != credenciais[i].SiteUrl ||
			credenciaisParaCriptografar[i].Login != credenciais[i].Login ||
			credenciaisParaCriptografar[i].Senha != credenciais[i].Senha {
			t.Errorf("Credencial descriptografada não corresponde à original")
		}
	}
}

*/
