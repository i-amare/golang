package utils

import (
	"errors"
	"fmt"
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

func VerifyAuthToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("Not ok!")
			return "", errors.New("Unexpected token signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return err
	}

	if !parsedToken.Valid {
		return errors.New("Authorisation token is invalid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("Token claims could not be parsed")
	}

	fmt.Println(claims["email"])
	fmt.Println(claims["userID"])
	fmt.Println(claims["exp"])

	return nil
}
