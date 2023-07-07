package symmetricEncryp_test

import (
	"crypto/rand"
	"encoding/hex"
	symmetricEncryp "safePasswordApi/src/security/encrypt/symmetrical"
	"testing"
)

func TestEncryptDataAES(t *testing.T) {
	data := "secretdata"

	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		t.Fatalf("Error generating AES key: %v", err)
	}

	ciphertext, err := symmetricEncryp.EncryptDataAES(data, hex.EncodeToString(key))
	if err != nil {
		t.Errorf("Error encrypting data with AES: %v", err)
	}

	if ciphertext == "" {
		t.Error("The encrypted ciphertext should not be empty")
	}
}

func TestDecryptDataAES_ValidData(t *testing.T) {
	data := "secretdata"

	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		t.Fatalf("Error generating AES key: %v", err)
	}

	ciphertext, err := symmetricEncryp.EncryptDataAES(data, hex.EncodeToString(key))
	if err != nil {
		t.Fatalf("Error encrypting data with AES: %v", err)
	}

	decryptedData, err := symmetricEncryp.DecryptDataAES(ciphertext, hex.EncodeToString(key))
	if err != nil {
		t.Errorf("Error decrypting AES ciphertext: %v", err)
	}

	if decryptedData != data {
		t.Errorf("Decrypted data does not match expected data. Expected: %s, Got: %s", data, decryptedData)
	}
}

func TestDecryptDataAES_InvalidCiphertext(t *testing.T) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		t.Fatalf("Error generating AES key: %v", err)
	}

	_, err = symmetricEncryp.DecryptDataAES("invalidciphertext", hex.EncodeToString(key))
	if err == nil {
		t.Error("Expected error when decrypting invalid ciphertext, but got nil")
	}
}

func TestGenerateRandomAESKey(t *testing.T) {
	_, err := symmetricEncryp.GenerateRandomAESKey() // Key size in bytes (256 bits)
	if err != nil {
		t.Errorf("Error generating random AES key: %v", err)
		return
	}
}
