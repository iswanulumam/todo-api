package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthRepository struct{}

func New() *AuthRepository {
	return &AuthRepository{}
}

func (a *AuthRepository) Login(username string, password string) (string, error) {
	if username == "admin" && password == "admin" {
		claims := jwt.MapClaims{}
		claims["authorized"] = true
		claims["id"] = 1
		claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
		claims["name"] = username

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token.SignedString([]byte("R4HASIA"))
	}
	return "", errors.New("failed login!")
}
