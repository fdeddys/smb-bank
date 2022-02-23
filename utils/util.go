package utils

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var randomizer *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// GeneratePassword ...
func GeneratePassword(length int) string {
	randomChar := make([]byte, length)
	for i := range randomChar {
		randomChar[i] = charset[randomizer.Intn(len(charset))]
	}
	return string(randomChar)
}
