package service

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type AuthCustomClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func ValidateToken(tokenString string) (string, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	claims := &AuthCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", errors.New("invalid token")
	}
	return claims.ID, nil
}
