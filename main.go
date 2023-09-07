package main

import (
	"fmt"
	_ "safePasswordApi/docs"
	"safePasswordApi/src/configs"
	"safePasswordApi/src/router"
	"time"

	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
)

func init() {
	configs.InitializeConfigurations()
}

// @title SafePasswordApi
// @version 0.1
// @description SafePasswordApi é uma API de gerenciamento de senhas com criptografia integral,
// @description proporcionando armazenamento e acesso seguros para dados sensíveis de usuários,
// @description por meio de métodos eficientes de manipulação de senhas.

// @contact.name Gustavo Gama
// @contact.url https://www.linkedin.com/in/gustavo-gama-966b341b8/
// @contact.email gustavogama494@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	r := router.Gerar()
	r.Server.WriteTimeout = 30 * time.Second

	r.GET("/swagger/*", echoSwagger.WrapHandler)
	r.Logger.Fatal(r.Start(fmt.Sprintf(":%d", configs.Port)))
}
