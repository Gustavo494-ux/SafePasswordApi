package asymmetrical_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	privateKey, err := GeneratePrivateKey(2048)
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
	privateKey, err := GeneratePrivateKey(2048)
	if err != nil {
		t.Fatalf("Error generating private key: %v", err)
	}

	privateKeyPEM, err := ExportPrivateKey(privateKey)
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
	privateKey, err := GeneratePrivateKey(2048)
	if err != nil {
		t.Fatalf("Error generating private key: %v", err)
	}

	publicKey := GeneratePublicKey(privateKey)

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
	privateKey, err := GeneratePrivateKey(2048)
	if err != nil {
		t.Fatalf("Error generating private key: %v", err)
	}

	publicKey := GeneratePublicKey(privateKey)

	publicKeyPEM, err := ExportPublicKey(publicKey)
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
	privateKey, err := GeneratePrivateKey(2048)
	if err != nil {
		t.Fatalf("Error generating private key: %v", err)
	}

	publicKey := GeneratePublicKey(privateKey)

	data := []byte("secret data")

	cipherText, err := EncryptRSA(data, publicKey)
	if err != nil {
		t.Errorf("Error encrypting: %v", err)
	}

	decryptedData, err := DecryptRSA(cipherText, privateKey)
	if err != nil {
		t.Errorf("Error decrypting: %v", err)
	}

	if string(decryptedData) != string(data) {
		t.Error("The decrypted data does not match the original data")
	}
}

func GeneratePrivateKey(bits int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func ExportPrivateKey(privateKey *rsa.PrivateKey) (string, error) {
	if privateKey == nil {
		return "", errors.New("private key is nil")
	}

	derBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derBytes,
	}

	privateKeyPEM := pem.EncodeToMemory(block)
	return string(privateKeyPEM), nil
}

func ExportPublicKey(publicKey *rsa.PublicKey) (string, error) {
	if publicKey == nil {
		return "", errors.New("public key is nil")
	}

	derBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", err
	}

	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derBytes,
	}

	publicKeyPEM := pem.EncodeToMemory(block)
	return string(publicKeyPEM), nil
}

func GeneratePublicKey(privateKey *rsa.PrivateKey) *rsa.PublicKey {
	if privateKey == nil {
		return nil
	}

	publicKey := &privateKey.PublicKey
	return publicKey
}

func EncryptRSA(data []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	if publicKey == nil {
		return nil, errors.New("public key is nil")
	}

	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, data)
	if err != nil {
		return nil, err
	}

	return cipherText, nil
}

func DecryptRSA(cipherText []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	if privateKey == nil {
		return nil, errors.New("private key is nil")
	}

	data, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	if err != nil {
		return nil, err
	}

	return data, nil
}
