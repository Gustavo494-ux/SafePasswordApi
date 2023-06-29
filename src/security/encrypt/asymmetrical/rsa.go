package asymmetrical

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
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
