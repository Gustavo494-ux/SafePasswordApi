package models_test

import (
	"fmt"
	"os"
	enum "safePasswordApi/src/enum/geral"
	"safePasswordApi/src/models"
	"safePasswordApi/src/modules/logger"
	"safePasswordApi/src/routines/inicializacao"
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

// TestMain:Função executada antes das demais
func TestMain(m *testing.M) {
	inicializacao.CarregarDotEnv()
	inicializacao.InicializarEncriptacao()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestCredencial_Prepare(t *testing.T) {
	for _, credential := range credentials {
		err := credential.Preparar(enum.TipoPreparacao_Cadastro)
		if err != nil {
			logger.Logger().Error("Ocorreu um erro ao realizar o teste TestCredencial_Prepare", err, credential)
			t.Errorf("Ocorreu um erro ao realizar o teste TestCredencial_Prepare: %s", err.Error())
		}
	}
	logger.Logger().Info("Teste TestCredencial_Prepare executado com sucesso!")
}

func TestValidate(t *testing.T) {
	for _, credential := range credentials {
		err := credential.Validar()

		if err != nil {
			logger.Logger().Error("Ocorreu um erro inesperado ao executar o teste TestValidate", err, credential)
			t.Errorf("Unexpected error: %s", err.Error())
		}
	}
	logger.Logger().Info("Teste TestValidate executado com sucesso!")
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
				logger.Logger().Error("Esperava o erro 'o usuário é obrigatório e não pode ficar em branco', mas nenhum retornou", err, credential)
				t.Error("Esperava o erro 'o usuário é obrigatório e não pode ficar em branco', mas nenhum retornou")
			}
		} else if err.Error() != "o usuário é obrigatório e não pode ficar em branco" {
			logger.Logger().Error(fmt.Sprintf("Erro esperado: usuário é obrigatório e não pode ficar em branco,mas retornou : %s", err), err, credential)
			t.Errorf("Erro esperado: usuário é obrigatório e não pode ficar em branco,mas retornou : %s", err)
		}
	}
	logger.Logger().Info("Teste TestValidate_UsuarioIdZero executado com sucesso!")
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
				logger.Logger().Error("Esperava o erro 'a senha é obrigatória e não pode ficar em branco', mas nenhum retornou", err, credential)
				t.Error("Esperava o erro 'a senha é obrigatória e não pode ficar em branco', mas nenhum retornou")
			}
		} else if err.Error() != "a senha é obrigatória e não pode ficar em branco" {
			logger.Logger().Error(fmt.Sprintf("Erro esperado: a senha é obrigatória e não pode ficar em branco,mas retornou : %s", err), err, credential)
			t.Errorf("Erro esperado: a senha é obrigatória e não pode ficar em branco,mas retornou : %s", err)
		}
	}
	logger.Logger().Info("Teste TestValidate_SenhaVazia executado com sucesso!")
}

func TestFormat_SaveData(t *testing.T) {
	var err error
	for _, credential := range credentials {
		err = credential.Formatar(enum.TipoFormatacao_Cadastro)

		if err != nil {
			t.Errorf("Erro inesperado: %s", err.Error())
			logger.Logger().Error("Erro inesperado", err, credential)
		}
	}
	logger.Logger().Info("Teste TestFormat_SaveData executado com sucesso!")
}

func TestFormat_RetrieveData(t *testing.T) {
	var err error
	for _, credential := range credentials {
		err = credential.Preparar(enum.TipoPreparacao_Consulta)

		if err.Error() != "data not encrypted with RSA" {
			t.Errorf("Erro inesperado: %v", err)
			logger.Logger().Error("Erro inesperado", err, credential)

		}
	}
	logger.Logger().Info("Teste TestFormat_RetrieveData executado com sucesso!")
}

func TestEncryptAES(t *testing.T) {
	var err error
	credentialsWithEmptyPassword := credentials
	for _, credential := range credentialsWithEmptyPassword {
		err = credential.CriptografarAES()
		if err != nil {
			t.Errorf("Erro inesperado: %s", err.Error())
			logger.Logger().Error("Erro inesperado ao criptografar", err, credential)
		}
	}
	logger.Logger().Info("Teste TestEncryptAES executado com sucesso!")
}

func TestDecryptAES(t *testing.T) {
	var err error
	for _, credential := range credentials {
		err = credential.DescriptografarAES()
		if err != nil && err.Error() != "unencrypted data using AES 256" {
			t.Errorf("Erro inesperado: %s", err.Error())
			logger.Logger().Error("Erro inesperado ao descriptografar", err, credential)
		}
	}
	logger.Logger().Info("Teste TestDecryptAES executado com sucesso!")
}

func TestEncryptDecryptAES(t *testing.T) {
	var err error

	// Encrypt credentials
	credentialsToEncrypt := credentials

	for i := range credentialsToEncrypt {
		err = credentialsToEncrypt[i].CriptografarAES()
		if err != nil {
			t.Errorf("Erro Inesperado: %s", err.Error())
			logger.Logger().Error("Erro inesperado ao criptografar", err, credentialsToEncrypt[i])
		}
	}

	// Decrypt credentials
	for i := range credentialsToEncrypt {
		err = credentialsToEncrypt[i].DescriptografarAES()
		if err != nil {
			t.Errorf("Erro inesperado ao descriptografar: %s", err.Error())
			logger.Logger().Error("Erro inesperado ao descriptografar", err, credentialsToEncrypt[i])
		}
	}

	// Compare original credentials with decrypted ones
	for i := range credentials {
		if credentialsToEncrypt[i].Descricao != credentials[i].Descricao ||
			credentialsToEncrypt[i].SiteUrl != credentials[i].SiteUrl ||
			credentialsToEncrypt[i].Login != credentials[i].Login ||
			credentialsToEncrypt[i].Senha != credentials[i].Senha {
			t.Errorf("A credencial descriptografada não corresponde à original")
			logger.Logger().Error("A credencial descriptografada não corresponde à original", err, credentialsToEncrypt[i])
		}
	}
	logger.Logger().Info("Teste TestEncryptDecryptAES executado com sucesso!")
}

func TestEncryptRSA(t *testing.T) {
	var err error
	for _, credential := range credentials {
		err = credential.CriptografarRSA()
		if err != nil {
			t.Errorf("Erro inesperado ao criptografar: %s", err.Error())
			logger.Logger().Error("Erro inesperado", err, credential)
		}
	}
	logger.Logger().Info("Teste TestEncryptRSA executado com sucesso!")
}

func TestDecryptRSA(t *testing.T) {
	var err error
	for _, credential := range credentials {
		err = credential.DescriptografarRSA()
		if err != nil && err.Error() != "data not encrypted with RSA" {
			t.Errorf("Erro inesperado: %s", err.Error())
			logger.Logger().Error("Erro inesperado ao descriptografar", err, credential)
		}
	}
	logger.Logger().Info("Teste TestDecryptRSA executado com sucesso!")
}

func TestEncryptDecryptRSA(t *testing.T) {
	var err error

	// Encrypt credentials
	credentialsToEncrypt := credentials

	for i := range credentialsToEncrypt {
		err = credentialsToEncrypt[i].CriptografarRSA()
		if err != nil {
			t.Errorf("Unexpected error while encrypting: %s", err.Error())
			logger.Logger().Error("Erro inesperado ao criptografar", err, credentialsToEncrypt[i])
		}
	}

	// Decrypt credentials
	for i := range credentialsToEncrypt {
		err = credentialsToEncrypt[i].DescriptografarRSA()
		if err != nil {
			t.Errorf("Erro inesperado ao descriptografar: %s", err.Error())
			logger.Logger().Error("Erro inesperado ao descriptografar", err, credentialsToEncrypt[i])
		}
	}

	// Compare original credentials with decrypted ones
	for i := range credentials {
		if credentialsToEncrypt[i].Descricao != credentials[i].Descricao ||
			credentialsToEncrypt[i].SiteUrl != credentials[i].SiteUrl ||
			credentialsToEncrypt[i].Login != credentials[i].Login ||
			credentialsToEncrypt[i].Senha != credentials[i].Senha {
			logger.Logger().Error("A credencial descriptografada não corresponde à original", err, credentialsToEncrypt[i])
		}
	}
	logger.Logger().Info("Teste TestEncryptDecryptRSA executado com sucesso!")
}
