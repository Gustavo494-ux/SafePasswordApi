package configs

import (
	"os"
)

const (
	FormatoDataHora   = "02/01/2006 15:04:05"
	CaminhoArquivoLog = "./Log_safepassword.log"
	RootDirectory     = "SafePasswordApi"
)

var (
	StringConnection  = ""
	Port              = 0
	SecretKeyJWT      []byte
	RSAPrivateKey     string
	RSAPublicKey      string
	AESKey            string
	RSAPrivateKeyPath string
	RSAPublicKeyPath  string
	AESKeyPath        string
	SecretKeyJWTPath  string
	ArquivoLog        os.File
)
