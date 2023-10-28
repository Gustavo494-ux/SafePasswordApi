package configuracoes

import (
	"log"
	"os"
	"safePasswordApi/src/modules/tipo"
	symmetricEncryp "safePasswordApi/src/security/encrypt/symmetrical"
)

type AES struct {
	Chave   *string
	Caminho string
}

const (
	chave256 = 32
)

// ConfigurarChavesAES: Configura as chaves AES
func (aes *AES) ConfigurarChavesAES() {
	aes.carregarCaminhosChavesAES()
	aes.formatarCaminhosAES()

	aes.configurarChaveAES256()
}

// configurarChaveAES256 realiza realiza toda a configuração da chaves AES 256
func (aes *AES) configurarChaveAES256() {
	configurarChave(aes.Caminho, aes.Chave, chave256)
}

// configurarChave: configura uma chave AES a partir do caminho e tamnho
func configurarChave(caminho string, chave *string, tamanho int) {
	var err error
	CarregarChaveDeArquivo(caminho, chave)
	switch {
	case len(*chave) == 0:
		{
			*chave, err = symmetricEncryp.GenerateRandomAESKey()
			if err != nil {
				log.Fatal("Erro ao gerar uma chave AES", err)
			}

		}
	case len(*chave) == tamanho:
	case len(*chave) >= tamanho:
		*chave = string(*chave)[:tamanho]
	default:
		log.Fatal("Chave AES inválida")
	}
}

// carregarCaminhosChavesAES: carrega os caminhos relacionados a AES do env
func (aes *AES) carregarCaminhosChavesAES() {
	aes.Caminho = tipo.Coalesce().Str(os.Getenv("AES_KEY_PATH"), "ES_KEY.key")
}

// formatarCaminhosAES: Formata os caminhos das chaves
func (aes *AES) formatarCaminhosAES() {
	formatarCaminho(&aes.Caminho)
}
