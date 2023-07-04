package asymmetrical_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"math/big"
	"safePasswordApi/src/security/encrypt/asymmetrical"
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	privateKey, err := asymmetrical.GeneratePrivateKey(2048)
	if err != nil {
		t.Errorf("Error generating private key: %v", err)
	}

	if privateKey == nil {
		t.Error("Private key should not be nil")
	}

	// Verify if the private key is valid
	err = privateKey.Validate()
	if err != nil {
		t.Errorf("Invalid private key: %v", err)
	}
}

func TestExportPrivateKey(t *testing.T) {
	privateKey, err := asymmetrical.GeneratePrivateKey(2048)
	if err != nil {
		t.Fatalf("Error generating private key: %v", err)
	}

	privateKeyPEM, err := asymmetrical.ExportPrivateKey(privateKey)
	if err != nil {
		t.Errorf("Error exporting private key: %v", err)
	}

	if privateKeyPEM == "" {
		t.Error("Private key PEM should not be empty")
	}

	// Verify if it is possible to parse the private key PEM
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		t.Error("Failed to parse the private key PEM")
	}

	// Parse the private key
	parsedPrivateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		t.Errorf("Error parsing the private key: %v", err)
	}

	// Verify if the exported private key matches the original private key
	if parsedPrivateKey == nil || parsedPrivateKey.N.Cmp(privateKey.N) != 0 {
		t.Error("The exported private key does not match the original private key")
	}
}

func TestGeneratePublicKey(t *testing.T) {
	privateKey, err := asymmetrical.GeneratePrivateKey(2048)
	if err != nil {
		t.Fatalf("Error generating private key: %v", err)
	}

	publicKey := asymmetrical.GeneratePublicKey(privateKey)

	if publicKey == nil {
		t.Error("Public key should not be nil")
	}

	// Verify if the public key matches the original private key
	if publicKey == nil || privateKey.PublicKey.N.Cmp(publicKey.N) != 0 || privateKey.PublicKey.E != publicKey.E {
		t.Error("The generated public key does not match the original private key")
	}

	// Verify if the public key is valid (basic validation example)
	if publicKey.N.BitLen() < 2048 {
		t.Error("The generated public key is invalid")
	}
}

func TestExportPublicKey(t *testing.T) {
	privateKey, err := asymmetrical.GeneratePrivateKey(2048)
	if err != nil {
		t.Fatalf("Error generating private key: %v", err)
	}

	publicKey := asymmetrical.GeneratePublicKey(privateKey)

	publicKeyPEM, err := asymmetrical.ExportPublicKey(publicKey)
	if err != nil {
		t.Errorf("Error exporting public key: %v", err)
	}

	if publicKeyPEM == "" {
		t.Error("Public key PEM should not be empty")
	}

	// Verify if it is possible to parse the public key PEM
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil || block.Type != "PUBLIC KEY" {
		t.Error("Failed to parse the public key PEM")
	}

	// Parse the public key
	parsedPublicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		t.Errorf("Error parsing the public key: %v", err)
	}

	// Convert to *rsa.PublicKey type
	parsedRSAPublicKey, ok := parsedPublicKey.(*rsa.PublicKey)
	if !ok {
		t.Error("Failed to convert the public key to *rsa.PublicKey type")
	}

	// Verify if the exported public key matches the original public key
	if parsedRSAPublicKey == nil || parsedRSAPublicKey.N.Cmp(publicKey.N) != 0 {
		t.Error("The exported public key does not match the original public key")
	}
}

func TestEncryptAndDecryptRSA(t *testing.T) {
	privateKey, err := asymmetrical.GeneratePrivateKey(2048)
	if err != nil {
		t.Fatalf("Error generating private key: %v", err)
	}

	publicKey := asymmetrical.GeneratePublicKey(privateKey)

	data := []byte("secret data")

	cipherText, err := asymmetrical.EncryptRSA(data, publicKey)
	if err != nil {
		t.Errorf("Error encrypting: %v", err)
	}

	decryptedData, err := asymmetrical.DecryptRSA(cipherText, privateKey)
	if err != nil {
		t.Errorf("Error decrypting: %v", err)
	}

	if string(decryptedData) != string(data) {
		t.Error("The decrypted data does not match the original data")
	}
}

func TestValidatePrivateKey_ValidPrivateKey(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	err = asymmetrical.ValidatePrivateKey(string(privateKeyPEM))
	if err != nil {
		t.Errorf("Expected private key to be valid, but got error: %v", err)
	}
}

func TestValidatePrivateKey_InvalidPrivateKey(t *testing.T) {
	// Creating an invalid private key by explicitly setting a factor to zero
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	// Set a factor to zero to make the private key invalid
	privateKey.Primes[0] = big.NewInt(0)

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	err = asymmetrical.ValidatePrivateKey(string(privateKeyPEM))
	if err == nil {
		t.Error("Expected private key to be invalid, but it was considered valid")
	} else {
		t.Logf("Expected error: %v", err)
	}
}

func TestValidatePublicKey_ValidPublicKey(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	publicKey := &privateKey.PublicKey

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		t.Fatalf("Failed to marshal public key: %v", err)
	}

	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	err = asymmetrical.ValidatePublicKey(string(publicKeyPEM))
	if err != nil {
		t.Errorf("Expected public key to be valid, got error: %v", err)
	}
}

func TestValidatePublicKey_InvalidPublicKey(t *testing.T) {
	// Invalid public key with null modulus (N)
	publicKey := &rsa.PublicKey{
		N: nil,
		E: 65537, // Valid exponent (E)
	}

	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	})

	err := asymmetrical.ValidatePublicKey(string(publicKeyPEM))
	if err == nil {
		t.Error("Expected public key to be invalid, but it was considered valid")
	} else {
		t.Logf("Expected error: %v", err)
	}

	// Invalid public key with invalid exponent (E)
	publicKey = &rsa.PublicKey{
		N: big.NewInt(123456789), // Valid modulus (N)
		E: 1,                     // Invalid exponent (E)
	}

	publicKeyPEM = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	})

	err = asymmetrical.ValidatePublicKey(string(publicKeyPEM))
	if err == nil {
		t.Error("Expected public key to be invalid, but it was considered valid")
	} else {
		t.Logf("Expected error: %v", err)
	}

	// Invalid public key with negative modulus (N)
	publicKey = &rsa.PublicKey{
		N: big.NewInt(-123456789), // Invalid modulus (negative N)
		E: 65537,                  // Valid exponent (E)
	}

	publicKeyPEM = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	})

	err = asymmetrical.ValidatePublicKey(string(publicKeyPEM))
	if err == nil {
		t.Error("Expected public key to be invalid, but it was considered valid")
	} else {
		t.Logf("Expected error: %v", err)
	}
}
