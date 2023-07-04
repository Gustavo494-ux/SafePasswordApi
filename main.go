package main

import (
	"fmt"
	"log"
	"net/http"
	"safePasswordApi/src/configs"
	"safePasswordApi/src/router"
)

func main() {
	configs.InitializeConfigurations()
	fmt.Println("Rodando API!")

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configs.Port), r))
}
