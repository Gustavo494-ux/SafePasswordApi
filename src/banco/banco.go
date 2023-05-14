package banco

import (
	"api/src/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//    _ "github.com/go-sql-driver/mysql"

func Conectar() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", config.StringConexao)
	if err != nil {
		return nil, err
	}

	return db, nil
}
