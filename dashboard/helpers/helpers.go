package helpers

import (
	"crypto/rand"
	"encoding/hex"
)

// CheckError handler
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// CheckError with custom message
func CheckErrorWithMessage(err error, message string) {
	if err != nil {
		panic(message)
	}
}

// GenerateRandomToken generates a random token of length
func GenerateProjectToken(length int) (string, error) {
	if length%2 != 0 {
		return "", nil
	}

	bytes := make([]byte, length/2)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}
