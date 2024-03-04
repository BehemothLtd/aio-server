package models

import (
	"aio-server/enums"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Id                int32
	Email             string
	EncryptedPassword string
	FullName          string
	LockVersion       int
	Name              string
	FavoritedSnippets []*Snippet   `gorm:"many2many:snippets_favorites"`
	UserGroups        []*UserGroup `gorm:"many2many:users_user_groups"`
	About             string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	CompanyLevelId    *int32
	Address           *string
	Phone             *string
	SlackId           *string
	Gender            *enums.UserGenderType
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
