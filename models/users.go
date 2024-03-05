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
	Avatar            *Attachment  `gorm:"polymorphic:Owner;"`
	About             *string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	CompanyLevelId    *int32
	Address           *string
	Phone             *string
	Gender            *enums.UserGenderType
	Birthday          time.Time
	SlackId           string
	State             enums.UserState
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

type UsersQuery struct {
	NameCont     *string
	FullNameCont *string
	EmailCont    *string
	SlackIdCont  *string
	StateEq      *string
}
