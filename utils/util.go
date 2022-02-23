package utils

import (
	"fmt"
	"math/rand"
	"time"

	"com.ocbc.smb/constants"
	"github.com/dgrijalva/jwt-go"
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

func GenerateToken(userName string, userId int64) (interface{}, error) {
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := sign.Claims.(jwt.MapClaims)
	claims["user"] = userName
	claims["userId"] = fmt.Sprintf("%v", (userId))

	unixNano := time.Now().UnixNano()
	umillisec := unixNano / 1000000
	timeToString := fmt.Sprintf("%v", umillisec)
	claims["tokenCreated"] = timeToString

	token, err := sign.SignedString([]byte(constants.TokenSecretKey))
	if err != nil {
		return nil, err
	}
	return token, nil
}
