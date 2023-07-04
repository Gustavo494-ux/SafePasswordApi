package hashEncrpt_test

import (
	hashEncrpt "safePasswordApi/src/security/encrypt/hash"
	"testing"
)

func TestGenerateSHA512(t *testing.T) {
	data := "password123"
	hash, err := hashEncrpt.GenerateSHA512(data)
	if err != nil {
		t.Errorf("Error generating SHA512 hash: %v", err)
	}

	if hash == "" {
		t.Error("The generated hash should not be empty")
	}
}

func TestCompareSHA512_ValidCredentials(t *testing.T) {
	hash := "bed4efa1d4fdbd954bd3705d6a2a78270ec9a52ecfbfb010c61862af5c76af1761ffeb1aef6aca1bf5d02b3781aa854fabd2b69c790de74e17ecfec3cb6ac4bf"
	decryptedPassword := "password123"

	err := hashEncrpt.CompareSHA512(hash, decryptedPassword)
	if err != nil {
		t.Errorf("Error comparing SHA512 hashes: %v", err)
	}
}

func TestCompareSHA512_InvalidCredentials(t *testing.T) {
	hash := "bed4efa1d4fdbd954bd3705d6a2a78270ec9a52ecfbfb010c61862af5c76af1761ffyb1aef6aca1bf5d02b3781aa854fabd2b69c790de74e17ecfec3cb6ac4bf"
	decryptedPassword := "password123"

	err := hashEncrpt.CompareSHA512(hash, decryptedPassword)
	if err == nil {
		t.Error("Expected error when comparing SHA512 hashes, but got nil")
	}
}
