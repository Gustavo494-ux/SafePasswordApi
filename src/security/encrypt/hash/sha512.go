package hashEncrpt

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
)

// GenerateSHA512 generates a hash using the Sha512 algorithm for the provided data
func GenerateSHA512(Data string) (string, error) {
	h := sha512.New()
	_, err := h.Write([]byte(Data))
	if err != nil {
		return "", err
	}
	encryptedPassword := h.Sum(nil)
	hashHex := hex.EncodeToString(encryptedPassword)
	return hashHex, nil
}

// CompareSHA512 checks if the provided text matches the hash
func CompareSHA512(encryptedPassword, decryptedPassword string) error {
	decryptedHash, err := GenerateSHA512(decryptedPassword)
	if err != nil {
		return err
	}
	encryptedPasswordBytes := []byte(encryptedPassword)

	if fmt.Sprintf("%x", decryptedHash) != fmt.Sprintf("%x", encryptedPasswordBytes) {
		return errors.New("invalid login credentials")
	}

	return nil
}
