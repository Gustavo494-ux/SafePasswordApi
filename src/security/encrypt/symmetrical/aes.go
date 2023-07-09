package symmetricEncryp

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

// EncryptDataAES encrypts the data using the AES-256 algorithm
func EncryptDataAES(Data string, Key string) (string, error) {
	keyBytes := []byte(Key)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// Generate a unique initialization vector (IV)
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(Data))

	// Copy the IV to the beginning of the ciphertext
	copy(ciphertext[:aes.BlockSize], iv)

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(Data))

	return hex.EncodeToString(ciphertext), nil
}

// DecryptDataAES decrypts data using the AES-256 algorithm
func DecryptDataAES(Data string, Key string) (string, error) {
	if !IsTextEncryptedAES(Data) {
		return "", errors.New("unencrypted data using AES 256")
	}

	keyBytes := []byte(Key)
	ciphertext, err := hex.DecodeString(Data)
	if err != nil {
		return "", err
	}

	// Retrieve the IV from the beginning of the ciphertext
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}

// GenerateRandomAESKey generates a random AES-256 key
func GenerateRandomAESKey() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	//return hex.EncodeToString(key), nil
	return string(key), nil
}

// IsTextEncryptedAES checks if the text is encrypted with AES
func IsTextEncryptedAES(text string) bool {
	return len([]byte(text)) > aes.BlockSize
}
