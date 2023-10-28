package inicializacao

import (
	"fmt"
	"safePasswordApi/src/configs"
	"safePasswordApi/src/modules/logger"
	"safePasswordApi/src/utility/fileHandler"

	"github.com/joho/godotenv"
)

func CarregarDotEnv() {
	caminho, err := fileHandler.GetSourceDirectory(configs.RootDirectory)
	if err != nil {
		logger.Logger().Fatal("Ocorreu um erro ao montar o caminho do diretorio raiz", err, caminho)
	}
	err = godotenv.Load(fmt.Sprintf("%s/.env", caminho))
	if err != nil {
		logger.Logger().Fatal("Ocorreu um erro ao carregar o arquivo env", err, caminho)
	}
}
