package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "$2a$14$fP5BuSxkQP/nFAaLtEXRwuwy5a.TR7X7mtb/BjrorLLOaznFB/oPO"

func GenerateAuthToken(email string, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
