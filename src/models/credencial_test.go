package models_test

import (
	"safePasswordApi/src/configs"
	"safePasswordApi/src/models"
	"testing"
	"time"
)

var Path_DotEnv = "./../../.env"
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
	configs.InitializeConfigurations(Path_DotEnv)
	for _, credencial := range credenciais {
		err := credencial.Preparar("salvarDados", "")
		if err != nil {
			t.Fatalf("ocorreu um erro ao realizar a preparação da credencial, error: %v", err)
		}
	}
}

func TestValidar(t *testing.T) {
	for _, credencial := range credenciais {
		err := credencial.Validar()

		if err != nil {
			t.Errorf("Erro inesperado: %s", err.Error())
		}
	}
}

func TestValidar_UsuarioIdZero(t *testing.T) {
	var err error
	var credenciaisUsuarioIdZeroUm []models.Credencial
	credenciaisUsuarioIdZeroUm = credenciais
	//err := credencial.validar()
	for _, credencial := range credenciaisUsuarioIdZeroUm {
		if credencial.Id%2 == 1 {
			credencial.UsuarioId = 1
		} else {
			credencial.UsuarioId = 0
		}

		err = credencial.Validar()
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
	var credenciaisSenhaVazia []models.Credencial
	credenciaisSenhaVazia = credenciais
	for _, credencial := range credenciaisSenhaVazia {
		if credencial.Id%2 == 1 {
			credencial.Senha = ""
		} else {
			credencial.Senha = "Teste"
		}
		err = credencial.Validar()
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
	configs.InitializeConfigurations(Path_DotEnv)
	var err error
	var credenciaisSenhaVazia []models.Credencial
	credenciaisSenhaVazia = credenciais
	for _, credencial := range credenciaisSenhaVazia {
		err = credencial.Formatar("salvarDados")

		if err != nil {
			t.Errorf("Erro inesperado: %s", err.Error())
		}
	}
}

func TestFormatar_ConsultarDados(t *testing.T) {
	configs.InitializeConfigurations(Path_DotEnv)
	var err error
	var credenciaisSenhaVazia []models.Credencial
	credenciaisSenhaVazia = credenciais
	for _, credencial := range credenciaisSenhaVazia {
		err = credencial.Formatar("consultarDados")

		if err.Error() != "data not encrypted with RSA" {
			t.Errorf("error: %v", err)
		}
	}
}

/*
func TestCriptografarAES(t *testing.T) {
	credencial := Credencial{
		Id:        1,
		UsuarioId: 1,
		Descricao: "Credencial",
		SiteUrl:   "https://www.example.com",
		Login:     "user",
		Senha:     "senha",
		CriadoEm:  time.Now(),
	}

	err := credencial.criptografarAES()

	if err != nil {
		t.Errorf("Erro inesperado: %s", err.Error())
	}

	// Verifique se os campos foram criptografados corretamente
	// (implementar asserções adequadas para verificar os valores criptografados)
}

/*
func TestDescriptografarAES(t *testing.T) {
	credencial := Credencial{
		Id:        1,
		UsuarioId: 1,
		Descricao: "Credencial",
		SiteUrl:   "https://www.example.com",
		Login:     "user",
		Senha:     "senha",
		CriadoEm:  time.Now(),
	}

	err := credencial.descriptografarAES()

	if err != nil {
		t.Errorf("Erro inesperado: %s", err.Error())
	}

	// Verifique se os campos foram descriptografados corretamente
	// (implementar asserções adequadas para verificar os valores descriptografados)
}

func TestCriptografarRSA(t *testing.T) {
	credencial := Credencial{
		Id:        1,
		UsuarioId: 1,
		Descricao: "Credencial",
		SiteUrl:   "https://www.example.com",
		Login:     "user",
		Senha:     "senha",
		CriadoEm:  time.Now(),
	}

	err := credencial.criptografarRSA()

	if err != nil {
		t.Errorf("Erro inesperado: %s", err.Error())
	}

	// Verifique se os campos foram criptografados corretamente
	// (implementar asserções adequadas para verificar os valores criptografados)
}

func TestDescriptografarRSA(t *testing.T) {
	credencial := Credencial{
		Id:        1,
		UsuarioId: 1,
		Descricao: "Credencial",
		SiteUrl:   "https://www.example.com",
		Login:     "user",
		Senha:     "senha",
		CriadoEm:  time.Now(),
	}

	err := credencial.descriptografarRSA()

	if err != nil {
		t.Errorf("Erro inesperado: %s", err.Error())
	}

	// Verifique se os campos foram descriptografados corretamente
	// (implementar asserções adequadas para verificar os valores descriptografados)
}
*/
