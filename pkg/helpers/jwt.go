package helpers

import (
	"aio-server/models"
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(claims jwt.Claims) (token string, err error) {
	var t *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = t.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		return "", err
	}

	return token, nil
}

func DecodeJwtToken(tokenString string, userClaim *models.UserClaims) (err error) {
	token, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return err
	}

	// check token validity, for example token might have been expired
	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}
