package helper

import (
	"math/rand"
)

// GenerateRandomString generates random length of string. Possible characters: [a-Z:0-9]
func GenerateRandomString(length int) string {
	const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	token := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(characters)) // Generate a random index
		token[i] = characters[randomIndex]        // Assign a random character
	}

	return string(token)
}
