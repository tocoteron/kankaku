package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTCustomClaims struct {
	jwt.StandardClaims
}

// Generate token to authenticate user
func GenerateToken(secret []byte) (string, error) {
	// Set custom claims
	claims := &JWTCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Encode token
	t, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return t, nil
}

// Authenticate user by id and password
func Authenticate(id string, password string) bool {
	fmt.Println(id, password)
	if id == "test" && password == "password" {
		return true
	}
	return false
}
