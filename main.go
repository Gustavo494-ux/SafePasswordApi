package main

import (
	"fmt"
	"safePasswordApi/src/configs"
	"safePasswordApi/src/database"
	"safePasswordApi/src/router"
	"safePasswordApi/src/routines/inicializacao"
	"time"
)

func init() {
	inicializacao.Inicializar()
	database.TestarConexao()
}

func main() {
	r := router.Gerar()
	r.Server.WriteTimeout = 30 * time.Second
	r.Logger.Fatal(r.Start(fmt.Sprintf(":%d", configs.Port)))
}
