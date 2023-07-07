package asymmetrical

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
)

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

func GeneratePublicKey(privateKeyString string) (*rsa.PublicKey, error) {
	// Decode the private key from PEM string
	block, _ := pem.Decode([]byte(privateKeyString))
	if block == nil {
		return nil, errors.New("failed to decode private key")
	}

	// Parse the private key
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	// Get the public key from the private key
	publicKey := &privateKey.PublicKey
	return publicKey, nil
}

func EncryptRSA(data string, publicKey *rsa.PublicKey) (string, error) {
	if publicKey == nil {
		return "", errors.New("public key is nil")
	}

	// Convert a string to bytes
	plainText := []byte(data)

	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		return "", err
	}

	// Convert the cipher text bytes to a string
	cipherTextString := base64.StdEncoding.EncodeToString(cipherText)

	return cipherTextString, nil
}

func DecryptRSA(cipherText string, privateKey *rsa.PrivateKey) (string, error) {
	if privateKey == nil {
		return "", errors.New("private key is nil")
	}

	// Decode the cipher text string to bytes
	cipherTextBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	data, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherTextBytes)
	if err != nil {
		return "", err
	}

	// Convert the decrypted data bytes to a string
	dataString := string(data)

	return dataString, nil
}

func ValidatePrivateKey(privateKeyString string) error {
	privateKeyBytes := []byte(privateKeyString)

	// Parse the private key from the string
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		return errors.New("failed to parse private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %v", err)
	}

	if privateKey == nil {
		return errors.New("private key is nil")
	}

	err = privateKey.Validate()
	if err != nil {
		return err
	}

	return nil
}

func ValidatePublicKey(publicKeyString string) error {
	publicKeyBytes := []byte(publicKeyString)

	// Parse the public key from the string
	block, _ := pem.Decode(publicKeyBytes)
	if block == nil {
		return errors.New("failed to parse public key")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return fmt.Errorf("failed to parse public key: %v", err)
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return errors.New("invalid public key type")
	}

	// Check if the exponent is valid
	if rsaPublicKey.E < 2 {
		return errors.New("invalid public key exponent")
	}

	// Check if the modulus is valid
	if rsaPublicKey.N == nil || rsaPublicKey.N.Sign() != 1 {
		return errors.New("invalid public key modulus")
	}

	// Check if the modulus is odd
	if rsaPublicKey.N.Bit(0) != 1 {
		return errors.New("invalid public key modulus (not odd)")
	}

	// Check if the modulus is greater than the exponent
	if rsaPublicKey.N.Cmp(big.NewInt(int64(rsaPublicKey.E))) <= 0 {
		return errors.New("invalid public key modulus (less than or equal to the exponent)")
	}

	return nil
}
