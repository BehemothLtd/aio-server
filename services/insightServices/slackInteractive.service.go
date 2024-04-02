package insightServices

import (
	"aio-server/models"
	"fmt"

	"gorm.io/gorm"
)

type SlackInteractiveService struct {
	Db   *gorm.DB
	Args models.SlackInteractivePayload
}

func (sis *SlackInteractiveService) Excecute() (*models.SlackInteractiveResponse, error) {
	callbackId := sis.Args.CallbackId

	switch callbackId {
	case "change_state_rq":
		return sis.ChangeStateRequestResponse()
	default:
		return nil, nil
	}
}

func (sis *SlackInteractiveService) ChangeStateRequestResponse() (*models.SlackInteractiveResponse, error) {
	payload := sis.Args
	slackId := payload.User.Id
	user := models.User{SlackId: &slackId}
	dbTable := sis.Db.Table("users").Preload("UserGroups")

	err := dbTable.Where(&user).First(&user).Error

	if err != nil {
		return nil, err
	}

	result := models.SlackInteractiveResponse{
		ResponseType:    "ephemeral",
		ReplaceOriginal: false,
		Text:            "You have no permission to execute this action",
	}

	if user.IsBod() {
		// approverId := user.Id
		// TODO : trigger request - change state job

		action := payload.Action[0].Value
		text := payload.OriginalMessage.Text
		text += fmt.Sprintf("\n%s by <@%s>", action, slackId)

		result.ResponseType = "in_channel"
		result.ReplaceOriginal = true
		result.Text = text
	}

	fmt.Printf(">>>>>>>>> result object %+v\n", result)

	return &result, nil
}
