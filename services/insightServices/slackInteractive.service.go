package insightServices

import (
	"aio-server/exceptions"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/tasks"
	"fmt"
	"regexp"

	"gorm.io/gorm"
)

type SlackInteractiveService struct {
	Db   *gorm.DB
	Args models.SlackInteractivePayload
}

func (sis *SlackInteractiveService) Excecute() (*models.SlackInteractiveResponse, error) {
	callbackId := sis.Args.CallbackId

	switch callbackId {
	case constants.SlackChangeStateRq:
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
			return nil, err
		}

		info, err := tasks.AsynqClient.Enqueue(task)
		if err != nil {
			return nil, err
		}
		fmt.Printf("Task ID: %+v - completed at: %+v\n", info.ID, info.CompletedAt)

		action := payload.Action[0].Value
		text := payload.OriginalMessage.Text

		re := regexp.MustCompile(`<@([^>]+)>`)
		matchId := re.FindStringSubmatch(text)

		if len(matchId) > 1 {
			requesterId := matchId[1]
			text += fmt.Sprintf("\n%s by <@%s>", action, slackId)

			// Reply to request thread
			slackClient := models.NewSlackClient()
			replyText := fmt.Sprintf("<@%s> %s <@%s>'s request", *user.SlackId, action, requesterId)
			slackClient.SendMessage(replyText, payload.Channel.Id, nil, &payload.OriginalMessage.Ts)

			// Response to BOD's decision
			result.ResponseType = "in_channel"
			result.ReplaceOriginal = true
			result.Text = text
		} else {
			return nil, exceptions.NewBadRequestError("Requester not found!")
		}
	}

	return &result, nil
}
