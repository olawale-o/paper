package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func GetRole(role string) string {
	if role == "user" {
		return "user"
	}
	return "admin"
}

func signToken(claims *jwt.Token) (string, error) {
	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CreateToken(payload map[string]string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": payload["username"],
		"iss": "go-simple-rest",
		"aud": GetRole(payload["role"]),
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})
	tokenString, err := signToken(claims)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	sub, err := token.Claims.GetSubject()

	if err != nil {
		return "", fmt.Errorf("Unable to parse claim")
	}

	return sub, nil
}
