package database

import (
	"safePasswordApi/src/configs"
	"safePasswordApi/src/logsCatalogados"
	"safePasswordApi/src/modules/log"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (db *sqlx.DB, err error) {
	db, err = sqlx.Open("mysql", configs.StringConnection)
	if err != nil {
		log.Fatal(logsCatalogados.ErroConexaoBancoDeDados, err, configs.StringConnection)
	}

	log.Info(logsCatalogados.ConexaoBandoDeDadosEstabelecida)
	return
}

func TestarConexao() (err error) {
	_, err = sqlx.Connect("mysql", configs.StringConnection)
	if err != nil {
		log.Fatal(logsCatalogados.ErroConexaoBancoDeDados, err, configs.StringConnection)
		return
	}
	log.Info(logsCatalogados.TesteConexaoRealizado)
	return
}
