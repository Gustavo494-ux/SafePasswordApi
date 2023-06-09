package database

import (
	"safePasswordApi/src/configs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Conectar() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", configs.StringConexao)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestarConexao() (err error) {
	_, err = sqlx.Connect("mysql", configs.StringConexao)
	return
}