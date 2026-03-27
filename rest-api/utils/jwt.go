package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRET_KEY = "$2a$14$fP5BuSxkQP/nFAaLtEXRwuwy5a.TR7X7mtb/BjrorLLOaznFB/oPO"
const ADMIN_ID = 1

func GenerateAuthToken(email string, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	return token.SignedString([]byte(SECRET_KEY))
}

func VerifyAuthToken(token string) (int64, error) {
	claims, err := ParseAuthTokenClaims(token)
	if err != nil {
		return -1, err
	}

	userID := int64(claims["userID"].(float64))

	return userID, nil
}

func ParseAuthTokenClaims(token string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("Not ok!")
			return "", errors.New("Unexpected token signing method")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return jwt.MapClaims{}, err
	}

	if !parsedToken.Valid {
		return jwt.MapClaims{}, errors.New("Authorisation token is invalid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return jwt.MapClaims{}, errors.New("Token claims could not be parsed")
	}

	return claims, nil
}
