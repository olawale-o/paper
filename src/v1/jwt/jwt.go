package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func getRole(username string) string {
	if username == "senior" {
		return "senior"
	}
	return "employee"
}

func signToken(claims *jwt.Token) (string, error) {
	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func createToken(payload string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": payload,
		"iss": "go-simple-rest",
		"aud": getRole(payload),
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})
	fmt.Printf("Token claims added: %+v\n", claims)
	tokenString, err := signToken(claims)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
