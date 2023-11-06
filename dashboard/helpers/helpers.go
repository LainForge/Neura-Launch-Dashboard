package helpers

import (
    "crypto/rand"
    "encoding/hex"

    log "github.com/sirupsen/logrus"
)

// CheckError handler
func CheckError(err error) {
    if err != nil {
        log.Panic(err)
    }
}

// CheckError with custom message
func CheckErrorWithMessage(err error, message string) {
    if err != nil {
        log.Panic(message)
    }
}

// GenerateRandomToken generates a random token of length
func GenerateProjectToken(length int) (string, error) {

    // If we want a hex token of n bytes, we need 2n characters
    // because each byte is represented by 2 hex characters
    // And having an odd length will result in fractional number of bytes
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
