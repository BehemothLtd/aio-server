package models

import (
	"aio-server/enums"
	"aio-server/pkg/constants"
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
	SlackId           *string
	State             enums.UserState
	Timing            *UserTiming
	WorkingTimelogs   []WorkingTimelog
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

type ThisMonthWorkingHours struct {
	Hours                     float64
	PercentCompareToLastMonth float64
	UpFromLastMonth           bool
	TimeGraphOnProjects       TimeGraphOnProjects
}

type TimeGraphOnProjects struct {
	Labels []string
	Series []float64
}

type ProjectsWorkingHours struct {
	Hours float64
	Name  string
}

func (user *User) IsBod() bool {
	result := false
	for _, ug := range user.UserGroups {
		if ug.Title == constants.BODGroup {
			result = true

			return result
		}
	}
	return result
}
