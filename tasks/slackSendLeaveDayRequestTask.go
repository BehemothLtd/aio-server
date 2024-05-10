package tasks

import (
	"aio-server/database"
	"aio-server/models"
	"aio-server/repository"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/hibiken/asynq"
)

type SlackSendLeaveDayRequestTaskPayload struct {
	RequestId int32
	Mentions  *[]*string
}

func NewSlackSendLeaveDayRequestTask(requestId int32, mentions *[]*string) (*asynq.Task, error) {
	if payload, err := json.Marshal(SlackSendLeaveDayRequestTaskPayload{RequestId: requestId, Mentions: mentions}); err != nil {
		return nil, err
	} else {
		return asynq.NewTask(SlackSendLeaveDayRequestMessageTask, payload), nil
	}
}

func HandleSlackSendLeaveDayRequestTask(ctx context.Context, t *asynq.Task) error {
	db := database.Db

	var p SlackSendLeaveDayRequestTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	// Get message content
	request := models.LeaveDayRequest{Id: p.RequestId}
	db.Model(&request).First(&request)
	message := request.GetMessage(db, p.Mentions)

	// Send message to slack channel
	slackClient := models.NewSlackClient()
	slackMessage, err := slackClient.SendMessage(message, os.Getenv("SLACK_LEAVE_WFH_REQUEST_CHANNEL_ID"), nil, nil)
	if err != nil {
		return err
	}

	// Create request message
	requestMessage := models.Message{
		ParentId:  request.Id,
		Content:   &message,
		Timestamp: &slackMessage.Ts,
	}
	messageRepo := repository.NewMessageRepository(&ctx, db)
	messageRepo.Create(&requestMessage)

	return nil
}
