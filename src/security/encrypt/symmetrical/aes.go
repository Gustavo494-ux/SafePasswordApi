package symmetricEncryp

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	hashEncrpt "safePasswordApi/src/security/encrypt/hash"
)

// EncryptDataAES encrypts the data using the AES algorithm
func EncryptDataAES(Data string, Key string) (string, error) {
	keyBytes := reduceKey([]byte(Key), 32)
	iv := reduceKey([]byte(Key), aes.BlockSize)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(Data))
	stream.XORKeyStream(ciphertext, []byte(Data))

	ciphertext = append(iv, ciphertext...)
	return fmt.Sprintf("%x", ciphertext), nil
}

// DecryptDataAES decrypts data using the AES algorithm
func DecryptDataAES(Data string, Key string) (string, error) {
	keyBytes := reduceKey([]byte(Key), 32)
	iv := reduceKey([]byte(Key), aes.BlockSize)

	ciphertextBytes, err := hex.DecodeString(Data)
	if err != nil {
		return "", err
	}

	if len(ciphertextBytes) < aes.BlockSize {
		return "", errors.New("invalid ciphertext")
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	ciphertextBytes = ciphertextBytes[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertextBytes, ciphertextBytes)

	return string(ciphertextBytes), nil
}

func reduceKey(key []byte, newLength int) []byte {
	newKey, _ := hashEncrpt.GenerateSHA512(hex.EncodeToString(key))
	if newLength >= len(newKey) {
		return key
	}
	return append([]byte(nil), newKey[:newLength]...)
}

func GenerateRandomAESKey() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	keyString := hex.EncodeToString(key)
	return keyString, nil
}
