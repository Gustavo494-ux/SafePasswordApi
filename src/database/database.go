package database

import (
	"safePasswordApi/src/configs"
	"safePasswordApi/src/logsCatalogados"
	"safePasswordApi/src/modules/logger"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

// Conectar: Conecta ao banco de dados
func Conectar() (db *sqlx.DB, err error) {
	db, err = sqlx.Open("mysql", configs.StringConnection)
	if err != nil {
		logger.Logger().Error(logsCatalogados.ErroConexaoBancoDeDados, err, configs.StringConnection)
	}

	logger.Logger().Rastreamento(logsCatalogados.ConexaoBandoDeDadosEstabelecida)
	return
}

// TestarConexao: realiza o teste de conexão
func TestarConexao() {
	_, err := Conectar()
	if err != nil {
		logger.Logger().Fatal(logsCatalogados.ErroConexaoBancoDeDados, err, configs.StringConnection)
		return
	}
	logger.Logger().Rastreamento(logsCatalogados.TesteConexaoRealizado)
}
