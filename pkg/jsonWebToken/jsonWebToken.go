package jsonwebtoken

import (
	"aio-server/models"
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJwtToken generates a JWT token with the provided claims.
func GenerateJwtToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// DecodeJwtToken decodes and validates the JWT token.
func DecodeJwtToken(tokenString string, userClaim *models.UserClaims) error {
	token, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("invalid token")
	}
	return nil
}
