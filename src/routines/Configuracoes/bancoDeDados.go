package configuracoes

import (
	"fmt"
	"os"
	"safePasswordApi/src/modules/logger"
	"safePasswordApi/src/modules/tipo"
	"strconv"
)

type BancoDeDados struct {
	StringConexao *string
	Nome          string
	Usuario       string
	Senha         string
	Host          string
	Porta         int
}

// ConfigurarBancoDeDados: realiza as configurações necessárias para o funcionametno do banco de dados
func (bd *BancoDeDados) ConfigurarBancoDeDados() {
	bd.carregarVariaveis()
	bd.definirStringConexao()
}

// carregarVariaveis: carrega as variaveis necessarias
func (bd *BancoDeDados) carregarVariaveis() {
	var err error
	bd.Usuario = tipo.Coalesce().Str(os.Getenv("DB_USER"), "root")
	bd.Senha = tipo.Coalesce().Str(os.Getenv("DB_PASSWORD"), "")
	bd.Host = tipo.Coalesce().Str(os.Getenv("DB_HOST"), "localhost")
	bd.Nome = tipo.Coalesce().Str(os.Getenv("DB_NAME"), "SafePasswordApi")
	bd.Porta, err = strconv.Atoi(tipo.Coalesce().Str(os.Getenv("DB_PORT"), "3306"))
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao converter o DB_PORT do .env para inteiro", err, fmt.Sprintf("DB_PORT:%s ", os.Getenv("DB_PORT")))
	}
}

// definirStringConexao: define a string de conexão
func (bd *BancoDeDados) definirStringConexao() {
	*bd.StringConexao = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		bd.Usuario,
		bd.Senha,
		bd.Host,
		bd.Porta,
		bd.Nome,
	)
}
