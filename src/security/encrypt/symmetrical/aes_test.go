package symmetricEncryp_test

import (
	"fmt"
	"os"
	"safePasswordApi/src/configs"
	"safePasswordApi/src/modules/logger"
	"safePasswordApi/src/routines/inicializacao"
	symmetricEncryp "safePasswordApi/src/security/encrypt/symmetrical"
	"testing"
)

// TestMain:Função executada antes das demais
func TestMain(m *testing.M) {
	inicializacao.CarregarDotEnv()
	inicializacao.InicializarEncriptacao()
	exitCode := m.Run()
	os.Exit(exitCode)
}

// TestEncryptDataAES tests the EncryptDataAES function
func TestEncryptDataAES(t *testing.T) {
	data := "secretdata"

	ciphertext, err := symmetricEncryp.EncryptDataAES(data, configs.AESKey)
	if err != nil {
		t.Errorf("Erro ao criptografar dados com AES: %v", err)
		logger.Logger().Error("Erro ao criptografar dados com AES", err)
	}

	if ciphertext == "" {
		t.Error("O texto cifrado criptografado não deve estar vazio")
		logger.Logger().Error("O texto cifrado criptografado não deve estar vazio", err)
	}
	logger.Logger().Info("Teste TestEncryptDataAES executado com sucesso!")
}

// TestDecryptDataAES_ValidData tests the DecryptDataAES function with valid ciphertext
func TestDecryptDataAES_ValidData(t *testing.T) {
	data := "secretdata"

	ciphertext, err := symmetricEncryp.EncryptDataAES(data, configs.AESKey)
	if err != nil {
		t.Errorf("Erro ao criptografar dados com AES: %v", err)
		logger.Logger().Error("Erro ao criptografar dados com AES", err)
	}

	decryptedData, err := symmetricEncryp.DecryptDataAES(ciphertext, configs.AESKey)
	if err != nil && err.Error() != "unencrypted data using AES 256" {
		t.Errorf("Error decrypting AES ciphertext: %v", err)
		logger.Logger().Error("Error decrypting AES ciphertext", err)
	}

	if decryptedData != data {
		t.Errorf("Os dados descriptografados não correspondem aos dados esperados. Esperado: %s, Recebido: %s", data, decryptedData)
		logger.Logger().Error("Os dados descriptografados não correspondem aos dados esperados", err)
	}
	logger.Logger().Info("Teste TestDecryptDataAES_ValidData executado com sucesso!")
}

// TestDecryptDataAES_InvalidCiphertext tests the DecryptDataAES function with invalid ciphertext
func TestDecryptDataAES_InvalidCiphertext(t *testing.T) {
	_, err := symmetricEncryp.DecryptDataAES("invalidciphertext", configs.AESKey)
	if err == nil {
		t.Error("Erro esperado ao descriptografar texto cifrado inválido, mas obteve zero")
		logger.Logger().Error("Erro esperado ao descriptografar texto cifrado inválido, mas obteve zero", err)
	}
	logger.Logger().Info("Teste TestDecryptDataAES_InvalidCiphertext executado com sucesso!")
}

// TestGenerateRandomAESKey tests the GenerateRandomAESKey function
func TestGenerateRandomAESKey(t *testing.T) {
	key, err := symmetricEncryp.GenerateRandomAESKey() // Key size in bytes (256 bits)
	if len(key) != 32 {
		t.Errorf("Erro tamanho inválido da chave AES, tamanho: %v", len(key))
		logger.Logger().Error(fmt.Sprintf("Erro tamanho inválido da chave AES, tamanho: %v", len(key)), err)
	}
	if err != nil {
		t.Errorf("Erro ao gerar chave AES aleatória: %v", err)
		logger.Logger().Error("Erro ao gerar chave AES aleatória", err)
	}
	logger.Logger().Info("Teste TestGenerateRandomAESKey executado com sucesso!")
}
