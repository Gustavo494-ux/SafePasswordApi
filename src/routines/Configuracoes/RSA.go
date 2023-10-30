package configuracoes

import (
	"crypto/rsa"
	"fmt"
	"os"

	"safePasswordApi/src/modules/logger"
	"safePasswordApi/src/security/encrypt/asymmetrical"

	"safePasswordApi/src/modules/tipo"
)

type RSA struct {
	ChavePrivada *string
	ChavePublica *string

	CaminhoChavePrivada string
	CaminhoChavePublica string
}

// ConfigurarChavesRSA: Configura as chaves RSA
func (varRSA *RSA) ConfigurarChavesRSA() {
	varRSA.CarregarCaminhosChavesRSA()
	varRSA.FormatarCaminhosRSA()

	varRSA.ConfigurarChavePrivadaRSA()
	varRSA.ConfigurarChavePublicaRSA()
}

// ConfigurarChavePublicaRSA realiza realiza toda a configuração da chave pública
func (varRSA *RSA) ConfigurarChavePublicaRSA() {
	CarregarChaveDeArquivo(varRSA.CaminhoChavePublica, varRSA.ChavePublica)
	err := asymmetrical.ValidatePublicKey(*varRSA.ChavePublica)
	if err != nil {
		err = asymmetrical.ValidatePrivateKey(*varRSA.ChavePrivada)
		if err != nil {
			logger.Logger().Fatal(fmt.Sprintf("Chave privada RSA inválida, verifique: %s", varRSA.CaminhoChavePrivada), err)
		}

		PublicKey, err := asymmetrical.GeneratePublicKey(*varRSA.ChavePrivada)
		if err != nil {
			logger.Logger().Fatal(fmt.Sprintf("Ocorreu um erro ao gerar a chave pública RSA, verifique: %s", varRSA.CaminhoChavePublica), err)
		}

		*varRSA.ChavePublica, err = asymmetrical.ExportPublicKey(PublicKey)
		if err != nil {
			logger.Logger().Fatal(fmt.Sprintf("Ocorreu um erro ao exportar a chave pública RSA, verifique: %s", varRSA.CaminhoChavePublica), err)
		}
		ExportarChaveParaArquivo(varRSA.CaminhoChavePublica, *varRSA.ChavePublica)
	}
}

// ConfigurarChavePrivada realiza realiza toda a configuração da chave privada
func (varRSA *RSA) ConfigurarChavePrivadaRSA() {
	var chave *rsa.PrivateKey

	CarregarChaveDeArquivo(varRSA.CaminhoChavePrivada, varRSA.ChavePrivada)
	err := asymmetrical.ValidatePrivateKey(*varRSA.ChavePrivada)
	if err != nil {
		chave = GerarChavePrivadaRSA()
		varRSA.ExportarChaveRSAPrivadaString(chave)
		ExportarChaveParaArquivo(varRSA.CaminhoChavePrivada, *varRSA.ChavePrivada)
	}
}

// GerarChavePrivada: gera uma chave RSA privada
func GerarChavePrivadaRSA() (chavePrivada *rsa.PrivateKey) {
	chavePrivada, err := asymmetrical.GeneratePrivateKey(2048)
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao gerar a chave privada RSA", err)
	}
	return
}

// ExportarChavePrivadaString: exporta a chave privada para string
func (varRSA *RSA) ExportarChaveRSAPrivadaString(chavePrivada *rsa.PrivateKey) {
	var err error
	*varRSA.ChavePrivada, err = asymmetrical.ExportPrivateKey(chavePrivada)
	if err != nil {
		logger.Logger().Fatal("Erro ao exportar a chave privada RSA", err)
	}
}

// CarregarCaminhosChavesRSA: carrega os caminhos relacionados a RSA do env
func (varRSA *RSA) CarregarCaminhosChavesRSA() {
	varRSA.CaminhoChavePrivada = tipo.Coalesce().Str(os.Getenv("RSA_PRIVATE_KEY_PATH"), "")
	varRSA.CaminhoChavePublica = tipo.Coalesce().Str(os.Getenv("RSA_PUBLIC_KEY_PATH"), "")
}

// FormatarCaminhos: Formata os caminhos das chaves
func (varRSA *RSA) FormatarCaminhosRSA() {
	formatarCaminho(&varRSA.CaminhoChavePrivada)
	formatarCaminho(&varRSA.CaminhoChavePublica)
}
