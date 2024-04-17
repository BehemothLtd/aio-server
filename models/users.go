package models

import (
	"aio-server/enums"
	"aio-server/pkg/constants"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
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
	Avatar            *Attachment  `gorm:"polymorphic:Owner;polymorphicValue:User"`
	About             *string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	CompanyLevelId    *int32
	Address           *string
	Phone             *string
	Gender            *enums.UserGenderType
	Birthday          *time.Time
	SlackId           *string
	State             enums.UserState
	Timing            *UserTiming
	WorkingTimelogs   []WorkingTimelog
	ProjectAssignees  []ProjectAssignee
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

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed() {
		tx.Statement.SetColumn("lock_version", u.LockVersion+1)
	}

	return
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.assignDefaultData()
	return
}

func (u *User) assignDefaultData() (err error) {
	re := regexp.MustCompile(`(.*)@`)

	matches := re.FindStringSubmatch(u.Email)
	if len(matches) >= 2 {
		u.Name = matches[1]
	}

	timing := UserTiming{
		ActiveAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	u.Timing = &timing
	return
}
