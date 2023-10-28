package configuracoes

import (
	"path/filepath"
	"safePasswordApi/src/modules/logger"
	"safePasswordApi/src/utility/fileHandler"
)

// CarregarChaveDeArquivo: carrega a chave de um arquivo
func CarregarChaveDeArquivo(caminho string, chave *string) {
	var err error

	err = fileHandler.CreateDirectoryOrFileIfNotExists(caminho)
	if err != nil {
		logger.Logger().Fatal("ocorreu um erro ao criar o arquivo da chave, caso ele não exista", err, caminho)
	}

	diretorio, nomeArquivo := filepath.Split(caminho)
	*chave, err = fileHandler.OpenFile(diretorio, nomeArquivo)
	if err != nil {
		logger.Logger().Fatal("ocorreu um erro ao carregar uma das chaves RSA", err, caminho)
	}
}

// ExportarChaveParaArquivo: Exporta uma chave para arquivo
func ExportarChaveParaArquivo(Caminho, chave string) {
	diretorio, nomeArquivo := filepath.Split(Caminho)
	if err := fileHandler.WriteFile(diretorio, nomeArquivo, chave); err != nil {
		logger.Logger().Fatal("Ocorreu um erro ao exportar uma das chaves RSA", err, Caminho)
	}
}

func formatarCaminho(Caminho *string) {
	err := fileHandler.RetornarCaminhoAbsolutoConcatenaRelativo(Caminho)
	if err != nil {
		logger.Logger().Fatal("Ocorreu um erro ao formatar o caminho", err, Caminho)
	}
}
