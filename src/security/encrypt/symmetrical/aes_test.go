package symmetricEncryp_test

import (
	"safePasswordApi/src/configs"
	symmetricEncryp "safePasswordApi/src/security/encrypt/symmetrical"
	"testing"
)

var Path_DotEnv = "./../../../../.env"

// TestEncryptDataAES tests the EncryptDataAES function
func TestEncryptDataAES(t *testing.T) {
	configs.InitializeConfigurations(Path_DotEnv)
	data := "secretdata"

	ciphertext, err := symmetricEncryp.EncryptDataAES(data, configs.AESKey)
	if err != nil {
		t.Errorf("Error encrypting data with AES: %v", err)
	}

	if ciphertext == "" {
		t.Error("The encrypted ciphertext should not be empty")
	}
}

// TestDecryptDataAES_ValidData tests the DecryptDataAES function with valid ciphertext
func TestDecryptDataAES_ValidData(t *testing.T) {
	configs.InitializeConfigurations(Path_DotEnv)
	data := "secretdata"

	ciphertext, err := symmetricEncryp.EncryptDataAES(data, configs.AESKey)
	if err != nil {
		t.Fatalf("Error encrypting data with AES: %v", err)
	}

	decryptedData, err := symmetricEncryp.DecryptDataAES(ciphertext, configs.AESKey)
	if err != nil && err.Error() != "unencrypted data using AES 256" {
		t.Errorf("Error decrypting AES ciphertext: %v", err)
	}

	if decryptedData != data {
		t.Errorf("Decrypted data does not match expected data. Expected: %s, Got: %s", data, decryptedData)
	}
}

// TestDecryptDataAES_InvalidCiphertext tests the DecryptDataAES function with invalid ciphertext
func TestDecryptDataAES_InvalidCiphertext(t *testing.T) {
	configs.InitializeConfigurations(Path_DotEnv)

	_, err := symmetricEncryp.DecryptDataAES("invalidciphertext", configs.AESKey)
	if err == nil {
		t.Error("Expected error when decrypting invalid ciphertext, but got nil")
	}
}

// TestGenerateRandomAESKey tests the GenerateRandomAESKey function
func TestGenerateRandomAESKey(t *testing.T) {
	key, err := symmetricEncryp.GenerateRandomAESKey() // Key size in bytes (256 bits)
	if len(key) != 32 {
		t.Errorf("Error AES Key invalid size, size: %v", len(key))
	}
	if err != nil {
		t.Errorf("Error generating random AES key: %v", err)
	}
}
