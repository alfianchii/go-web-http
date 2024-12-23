package utils

import (
	"fmt"
	"time"
	"web-http/config"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtToken = []byte(config.GetENV("JWT_TOKEN_NAME"))

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateJWT(username string) (string, error) {
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.TokenDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtToken)
}

func ValidateJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtToken, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}
	return claims, nil
}

func GetBearerToken(authHeader string) (string, error) {
	unique := len("Bearer ")
	
	if authHeader == "" {
		return "", fmt.Errorf("unauthorized; authorization header is empty")
	}

	if len(authHeader) < unique || authHeader[:unique] != "Bearer " {
		return "", fmt.Errorf("unauthorized; uthorization header is not a Bearer token")
	}

	return authHeader[unique:], nil
}