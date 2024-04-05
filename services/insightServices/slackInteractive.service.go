package insightServices

import (
	"aio-server/models"
	"aio-server/tasks"
	"fmt"
	"log"

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
		// Update request state job
		task, err := tasks.NewSlackUpdateLeaveDayRequestStateTask(payload, user)
		if err != nil {
			log.Fatalf("could not create task: %v", err)
		}

		info, err := tasks.AsynqClient.Enqueue(task)
		if err != nil {
			log.Fatalf("could not enqueue task: %v", err)
		}
		fmt.Print(info)

		action := payload.Action[0].Value
		text := payload.OriginalMessage.Text
		text += fmt.Sprintf("\n%s by <@%s>", action, slackId)

		result.ResponseType = "in_channel"
		result.ReplaceOriginal = true
		result.Text = text
	}

	return &result, nil
}
