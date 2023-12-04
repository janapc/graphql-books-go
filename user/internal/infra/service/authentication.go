package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthCustomClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(userId, userEmail string) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &AuthCustomClaims{
		userId,
		userEmail,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})
	tokenString, err := token.SignedString(secret)
	return tokenString, err
}
