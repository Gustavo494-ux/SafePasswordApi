package inicializacao

import (
	"safePasswordApi/src/configs"
	configuracoes "safePasswordApi/src/routines/Configuracoes"
)

// inicializarEncriptacao: realiza as configurações necessárias para utilizar cada função de encriptação
func inicializarEncriptacao() {
	go inicializarRSA()
	go inicializarAES()
}

// inicializarRSA: realiza as configurações necessárias para utilizar o RSA
func inicializarRSA() {
	RSA := configuracoes.RSA{
		ChavePrivada:        &configs.RSAPrivateKey,
		ChavePublica:        &configs.RSAPublicKey,
		CaminhoChavePrivada: configs.RSAPrivateKeyPath,
		CaminhoChavePublica: configs.RSAPublicKey,
	}
	RSA.ConfigurarChavesRSA()
}

// inicializarAES: realiza as configurações necessárias para utilizar o AES
func inicializarAES() {
	aes := configuracoes.AES{
		Chave:   &configs.AESKey,
		Caminho: configs.AESKeyPath,
	}
	aes.ConfigurarChavesAES()
}
