package inicializacao

import (
	"safePasswordApi/src/configs"
	configuracoes "safePasswordApi/src/routines/Configuracoes"
	"sync"
)

// InicializarEncriptacao: realiza as configurações necessárias para utilizar cada função de encriptação
func InicializarEncriptacao() {
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer wg.Done()
		inicializarRSA()
	}()

	go func() {
		defer wg.Done()
		inicializarAES()
	}()

	go func() {
		defer wg.Done()
		inicializarJwt()
	}()

	wg.Wait()
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

// inicializarJwt: realiza as configurações necessárias para utilizar o Jwt
func inicializarJwt() {
	chaveString := string(configs.SecretKeyJWT)
	jwt := configuracoes.JWt{
		Chave:   &chaveString,
		Caminho: configs.SecretKeyJWTPath,
	}
	jwt.ConfigurarJWT()
	configs.SecretKeyJWT = []byte(chaveString)
}
