package main

import (
	"fmt"
	"safePasswordApi/src/configs"
	"safePasswordApi/src/router"
	"time"
)

func init() {
	configs.InitializeConfigurations()
}

func main() {

	r := router.Gerar()
	r.Server.WriteTimeout = 30 * time.Second
	r.Logger.Fatal(r.Start(fmt.Sprintf(":%d", configs.Port)))
}
