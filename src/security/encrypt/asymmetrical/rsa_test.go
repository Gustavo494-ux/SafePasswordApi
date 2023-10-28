package asymmetrical_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"math/big"
	"os"
	"safePasswordApi/src/modules/logger"
	"safePasswordApi/src/routines/inicializacao"
	"safePasswordApi/src/security/encrypt/asymmetrical"
	"testing"
)

// TestMain:Função executada antes das demais
func TestMain(m *testing.M) {
	inicializacao.CarregarDotEnv()
	inicializacao.InicializarEncriptacao()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestGeneratePrivateKey(t *testing.T) {
	privateKey, err := asymmetrical.GeneratePrivateKey(2048)
	if err != nil {
		t.Errorf("Error generating private key: %v", err)
		logger.Logger().Error("Ocorreu um erro ao gerar a chave privada RSA", err)
	}

	if privateKey == nil {
		t.Error("A chave privada não deve ser nula")
		logger.Logger().Error("A chave privada não deve ser nula", err)
	}

	// Verifica se a chave é valida
	err = privateKey.Validate()
	if err != nil {
		t.Errorf("Ocorreu um erro ao gerar a chave privada RSA: %v", err)
		logger.Logger().Error("Ocorreu um erro ao gerar a chave privada RSA", err)
	}
	logger.Logger().Info("Teste TestGeneratePrivateKey executado com sucesso!")
}

func TestExportPrivateKey(t *testing.T) {
	privateKey, err := asymmetrical.GeneratePrivateKey(2048)
	if err != nil {
		t.Fatalf("Erro ao gerar a chave privada: %v", err)
		logger.Logger().Error("Erro ao gerar chave privada", err)
	}

	privateKeyPEM, err := asymmetrical.ExportPrivateKey(privateKey)
	if err != nil {
		t.Errorf("Erro ao exportar a chave privada: %v", err)
		logger.Logger().Error("Erro ao exportar a chave privada", err)
	}

	if privateKeyPEM == "" {
		t.Error("A chave privada PEM não deve estar vazia")
		logger.Logger().Error("A chave privada PEM não deve estar vazia", err)
	}

	// Verifica se é possível analisar a chave privada PEM
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		t.Error("Falha ao decodificar a chave privada PEM")
		logger.Logger().Error("Falha ao decodificar a chave privada PEM", err)
	}

	// analisa a chave privada
	parsedPrivateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		t.Errorf("Erro ao converter a chave privada: %v", err)
		logger.Logger().Error("Erro ao converter a chave privada", err)
	}

	//Verifica se a chave privada exportada corresponde à chave privada original
	if parsedPrivateKey == nil || parsedPrivateKey.N.Cmp(privateKey.N) != 0 {
		t.Error("A chave privada exportada não corresponde à chave privada original")
		logger.Logger().Error("A chave privada exportada não corresponde à chave privada original", err)
	}

	logger.Logger().Info("Teste TestExportPrivateKey executado com sucesso!")
}

func TestGeneratePublicKey(t *testing.T) {
	privateKey, err := asymmetrical.GeneratePrivateKey(2048)
	if err != nil {
		t.Fatalf("Erro ao gerar a chave privada: %v", err)
		logger.Logger().Error("Erro ao gerar a chave privada", err)
	}

	RSAPrivateKey, err := asymmetrical.ExportPrivateKey(privateKey)
	if err != nil {
		t.Fatalf("Erro ao exportar chave privada: %v", err)
		logger.Logger().Error("Erro ao exportar chave privada", err)
	}

	publicKey, err := asymmetrical.GeneratePublicKey(RSAPrivateKey)
	if err != nil {
		t.Fatalf("Erro ao gerar chave publica: %v", err)
		logger.Logger().Error("Erro ao gerar chave publica", err)
	}

	if publicKey == nil {
		t.Error("A chave pública não deve ser nula")
		logger.Logger().Error("A chave pública não deve ser nula", err)
	}

	//Verifica se a chave pública corresponde à chave privada original
	if publicKey == nil || privateKey.PublicKey.N.Cmp(publicKey.N) != 0 || privateKey.PublicKey.E != publicKey.E {
		t.Error("A chave pública gerada não corresponde à chave privada original")
		logger.Logger().Error("A chave pública gerada não corresponde à chave privada original", err)
	}

	// Verifique se a chave pública é válida
	if publicKey.N.BitLen() < 2048 {
		t.Error("A chave pública gerada é inválida")
		logger.Logger().Error("A chave pública gerada é inválida", err)
	}
	logger.Logger().Info("Teste TestGeneratePublicKey executado com sucesso!")
}

func TestExportPublicKey(t *testing.T) {
	privateKey, err := asymmetrical.GeneratePrivateKey(2048)
	if err != nil {
		t.Fatalf("Erro ao gerar a chave privada: %v", err)
		logger.Logger().Error("Erro ao gerar a chave privada", err)
	}

	RSAPrivateKey, err := asymmetrical.ExportPrivateKey(privateKey)
	if err != nil {
		t.Fatalf("Erro ao exportar chave privada: %v", err)
		logger.Logger().Error("Erro ao exportar chave privada", err)
	}

	publicKey, err := asymmetrical.GeneratePublicKey(RSAPrivateKey)
	if err != nil {
		t.Fatalf("Erro ao gerar a chave pública: %v", err)
		logger.Logger().Error("Erro ao gerar a chave pública", err)
	}

	publicKeyPEM, err := asymmetrical.ExportPublicKey(publicKey)
	if err != nil {
		t.Errorf("Erro ao exportar a chave pública: %v", err)
		logger.Logger().Error("Erro ao exportar a chave pública", err)
	}

	if publicKeyPEM == "" {
		t.Error("A chave pública PEM não deve estar vazia")
		logger.Logger().Error("A chave pública PEM não deve estar vazia", err)
	}

	// Verifica se é possível analisar a chave pública PEM
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil || block.Type != "PUBLIC KEY" {
		t.Error("Falha ao decodificar a chave pública PEM")
		logger.Logger().Error("Falha ao decodificar a chave pública PEM", err)
	}

	// Analise a chave pública
	parsedPublicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		t.Errorf("Erro ao converter a chave pública: %v", err)
		logger.Logger().Error("Erro ao converter a chave pública", err)
	}

	// Converte para o tipo *rsa.PublicKey
	parsedRSAPublicKey, ok := parsedPublicKey.(*rsa.PublicKey)
	if !ok {
		t.Error("Falha ao converter a chave pública para o tipo *rsa.PublicKey")
		logger.Logger().Error("Falha ao converter a chave pública para o tipo *rsa.PublicKey", err)
	}

	// Verifica se a chave pública exportada corresponde à chave pública original
	if parsedRSAPublicKey == nil || parsedRSAPublicKey.N.Cmp(publicKey.N) != 0 {
		t.Error("A chave pública exportada não corresponde à chave pública original:")
		logger.Logger().Error("A chave pública exportada não corresponde à chave pública original", err)
	}
	logger.Logger().Info("Teste TestExportPublicKey executado com sucesso!")
}

func TestEncryptAndDecryptRSA(t *testing.T) {
	privateKey, err := asymmetrical.GeneratePrivateKey(2048)
	if err != nil {
		t.Fatalf("Erro ao gerar a chave privada: %v", err)
		logger.Logger().Error("Erro ao gerar a chave privada", err)
	}

	RSAPrivateKey, err := asymmetrical.ExportPrivateKey(privateKey)
	if err != nil {
		t.Fatalf("Erro ao exportar a chave privada: %v", err)
		logger.Logger().Error("Erro ao exportar a chave privada", err)
	}

	publicKey, err := asymmetrical.GeneratePublicKey(RSAPrivateKey)
	if err != nil {
		t.Fatalf("Erro ao gerar a chave pública: %v", err)
		logger.Logger().Error("Erro ao gerar a chave pública", err)
	}

	data := "secret data"

	cipherText, err := asymmetrical.EncryptRSA(data, publicKey)
	if err != nil {
		t.Errorf("Erro ao encriptar RSA: %v", err)
		logger.Logger().Error("Erro ao encriptar RSA", err)
	}

	decryptedData, err := asymmetrical.DecryptRSA(cipherText, privateKey)
	if err != nil {
		t.Errorf("Erro ao descriptografar RSA: %v", err)
		logger.Logger().Error("Erro ao descriptografar RSA", err)
	}

	if string(decryptedData) != string(data) {
		t.Error("Os dados descriptografados e originais são diferentes")
		logger.Logger().Error("Os dados descriptografados e originais são diferentes", err)
	}

	logger.Logger().Info("Teste TestEncryptAndDecryptRSA executado com sucesso!")
}

func TestValidatePrivateKey_ValidPrivateKey(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Erro ao gerar a chave privada: %v", err)
		logger.Logger().Error("Erro ao gerar a chave privada", err)
	}

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	err = asymmetrical.ValidatePrivateKey(string(privateKeyPEM))
	if err != nil {
		t.Errorf("Esperava-se que a chave privada fosse válida, mas ocorreu um erro: %v", err)
		logger.Logger().Error("Esperava-se que a chave privada fosse válida, mas ocorreu um erro", err)
	}
	logger.Logger().Info("Teste TestValidatePrivateKey_ValidPrivateKey executado com sucesso!")
}

func TestValidatePrivateKey_InvalidPrivateKey(t *testing.T) {
	// Criando uma chave privada inválida definindo explicitamente um fator como zero
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Erro ao exportar a chave privada: %v", err)
		logger.Logger().Error("Erro ao exportar a chave privada", err)
	}

	// Defina um fator como zero para tornar a chave privada inválida
	privateKey.Primes[0] = big.NewInt(0)

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	err = asymmetrical.ValidatePrivateKey(string(privateKeyPEM))
	if err == nil {
		t.Error("Esperava-se que a chave privada fosse inválida, mas foi considerada válida")
		logger.Logger().Error("Esperava-se que a chave privada fosse inválida, mas foi considerada válida", err)
	}
	logger.Logger().Info("Teste TestValidatePrivateKey_InvalidPrivateKey executado com sucesso!")
}

func TestValidatePublicKey_ValidPublicKey(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Erro ao gerar a chave privada: %v", err)
		logger.Logger().Error("Erro ao gerar a chave privada", err)
	}

	publicKey := &privateKey.PublicKey

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		t.Fatalf("Erro ao empacotar a chave pública: %v", err)
		logger.Logger().Error("Erro ao empacotar a chave pública", err)
	}

	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	err = asymmetrical.ValidatePublicKey(string(publicKeyPEM))
	if err != nil {
		t.Errorf("Esperava-se que a chave pública fosse válida, ocorreu um erro: %v", err)
		logger.Logger().Error("Esperava-se que a chave pública fosse válida, ocorreu um erro", err)
	}
	logger.Logger().Info("Teste TestValidatePublicKey_ValidPublicKey executado com sucesso!")
}
func TestValidatePublicKey_InvalidPublicKey(t *testing.T) {
	// Chave pública inválida com módulo nulo (N)
	publicKey := &rsa.PublicKey{
		N: nil,
		E: 65537, // Exponent válido (E)
	}

	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	})

	err := asymmetrical.ValidatePublicKey(string(publicKeyPEM))
	if err == nil {
		t.Error("Esperava-se que a chave pública fosse inválida, mas foi considerada válida")
		logger.Logger().Error("Esperava-se que a chave pública fosse inválida, mas foi considerada válida", err)
	}

	// Chave pública inválida com expoente inválido (E)
	publicKey = &rsa.PublicKey{
		N: big.NewInt(123456789), // Módulo válido (N)
		E: 1,                     // Exponente inválido (E)
	}

	publicKeyPEM = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	})

	err = asymmetrical.ValidatePublicKey(string(publicKeyPEM))
	if err == nil {
		t.Error("Esperava-se que a chave pública fosse inválida, mas foi considerada válida")
		logger.Logger().Error("Esperava-se que a chave pública fosse inválida, mas foi considerada válida", err)
	}

	// Chave pública inválida com módulo negativo (N)
	publicKey = &rsa.PublicKey{
		N: big.NewInt(-123456789), // Módulo inválido (N negativo)
		E: 65537,                  // Exponente válido (E)
	}

	publicKeyPEM = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	})

	err = asymmetrical.ValidatePublicKey(string(publicKeyPEM))
	if err == nil {
		t.Error("Esperava-se que a chave pública fosse inválida, mas foi considerada válida")
		logger.Logger().Error("Esperava-se que a chave pública fosse inválida, mas foi considerada válida", err)
	}
	logger.Logger().Info("Teste TestValidatePublicKey_InvalidPublicKey executado com sucesso!")
}
