package configuracoes

import (
	"fmt"
	"os"
	"safePasswordApi/src/modules/logger"
	"safePasswordApi/src/modules/tipo"
	hashEncrpt "safePasswordApi/src/security/encrypt/hash"
	symmetricEncryp "safePasswordApi/src/security/encrypt/symmetrical"
	"strings"
	"time"
)

type JWt struct {
	Chave   *string
	Caminho string
}

// ConfigurarChavesAES: Configura as chaves AES
func (jwt *JWt) ConfigurarJWT() {
	jwt.carregarCaminhoChaveJwt()
	jwt.formatarCaminhoJWT()

	jwt.configurarChaveJWT()
}

// configurarChaveJWT256: Configura a chave  JWT
func (jwt *JWt) configurarChaveJWT() {
	CarregarChaveDeArquivo(jwt.Caminho, jwt.Chave)
	if len(strings.Trim(*jwt.Chave, " ")) == 0 {
		ChaveAESAleatoria, err := symmetricEncryp.GenerateRandomAESKey()
		if err != nil {
			logger.Logger().Fatal("Ocorreu um erro ao gerar a chave secreta do JWT", err)
		}
		ChaveAESAleatoria, err = hashEncrpt.GenerateSHA512(ChaveAESAleatoria)
		if err != nil {
			logger.Logger().Fatal("Ocorreu um erro ao gerar a chave secreta do JWT", err)
		}

		*jwt.Chave, err = hashEncrpt.GenerateSHA512(fmt.Sprintf("%s,%s,%d", ChaveAESAleatoria, ChaveAESAleatoria, time.Now().Unix()))
		if err != nil {
			logger.Logger().Fatal("Ocorreu um erro ao gerar a chave secreta do JWT", err)
		}

		ExportarChaveParaArquivo(jwt.Caminho, *jwt.Chave)
	}
}

// carregarCaminhoChaveJwt: carrega os caminhos relacionados a JWT do env
func (jwt *JWt) carregarCaminhoChaveJwt() {
	jwt.Caminho = tipo.Coalesce().Str(os.Getenv("SECRET_KEY_JWT_PATH"), "SECRET_KEY_JWT.key")
}

// formatarCaminhoJWT: Formata os caminhos da chave
func (jwt *JWt) formatarCaminhoJWT() {
	formatarCaminho(&jwt.Caminho)
}
