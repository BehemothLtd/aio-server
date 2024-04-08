package models

import (
	"aio-server/enums"
	"aio-server/pkg/constants"
	"aio-server/pkg/utilities"
	"fmt"
	"os"
	"time"

	"gorm.io/gorm"
)

type LeaveDayRequest struct {
	Id           int32  `gorm:"not null;autoIncrement"`
	UserId       int32  `gorm:"not null;type:bigint;default:null"`
	ApproverId   *int32 `gorm:"not null;type:bigint;default:null"`
	User         User   `gorm:"foreignKey:UserId"`
	Approver     *User  `gorm:"foreignKey:ApproverId"`
	From         time.Time
	To           time.Time
	TimeOff      float64
	RequestType  enums.RequestType      `gorm:"not null;"`
	RequestState enums.RequestStateType `gorm:"not null;"`
	Reason       *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	LockVersion  int32 `gorm:"not null;autoIncrement;default:0"`
}

func (r *LeaveDayRequest) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed() {
		tx.Statement.SetColumn("lock_version", r.LockVersion+1)
	}

	return
}

func (request *LeaveDayRequest) GetMessage(db *gorm.DB, mentions *[]*string) string {
	user := User{Id: request.UserId}
	err := db.Table("users").Where(&user).First(&user).Error
	if err != nil {
		return err.Error()
	}

	requestType := utilities.SnakeCaseToHumanize(request.RequestType.String())
	from := request.From.Format(constants.DDMMYYY_HHMM_DateFormat)
	to := request.To.Format(constants.DDMMYYY_HHMM_DateFormat)

	groupId := os.Getenv("SLACK_GROUP_VN_MEMBER_ID")
	insightFrontDomain := os.Getenv("MM_FRONT_DOMAIN")

	if insightFrontDomain == "" {
		insightFrontDomain = "https://insight.behemoth.vn"
	}
	requetLink := fmt.Sprintf("%+s/leave_day_requests?id=%+d", insightFrontDomain, request.Id)

	message := fmt.Sprintf("<!subteam^%+s>\n<@%+v> requested %+s.\nFrom: %+s to: %+s\n%+s", groupId, *user.SlackId, requestType, from, to, requetLink)
	if request.Reason != nil {
		message += fmt.Sprintf("\nReason: %s", *request.Reason)
	}
	if mentions != nil {
		message += MentionText(mentions)
	}

	return message
}

func MentionText(mentions *[]*string) string {
	mentionText := "\n"
	for _, value := range *mentions {
		mentionText += fmt.Sprintf("<@%s>", *value)
	}

	return mentionText
}
