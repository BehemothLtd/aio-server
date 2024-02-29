package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Id                int32
	Email             string
	EncryptedPassword string
	FullName          string
	LockVersion       int32
	Name              string
	FavoritedSnippets []*Snippet `gorm:"many2many:snippets_favorites"`
}

type UserClaims struct {
	Sub int32
	jwt.RegisteredClaims
}

func (user *User) GenerateJwtClaims() (claims jwt.Claims) {
	return UserClaims{
		user.Id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}
