package base62

import (
	"crypto/rand"
	"math/big"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const aliasLength = 7

// GenerateRandomAlias creates a cryptographically secure 7-character string
func GenerateRandomAlias() (string, error) {
	result := make([]byte, aliasLength)
	base := big.NewInt(int64(len(alphabet)))

	for i := 0; i < aliasLength; i++ {
		// Securely generate a random number between 0 and 61
		num, err := rand.Int(rand.Reader, base)
		if err != nil {
			return "", err
		}

		// Map the random number to a character in our alphabet
		result[i] = alphabet[num.Int64()]
	}

	return string(result), nil
}
