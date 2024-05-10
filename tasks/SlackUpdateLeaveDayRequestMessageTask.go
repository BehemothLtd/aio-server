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

type SlackUpdateLeaveDayRequestMessageTaskPayload struct {
	Request models.LeaveDayRequest
}

func NewSlackUpdateLeaveDayRequestMessageTask(request models.LeaveDayRequest) (*asynq.Task, error) {
	if payload, err := json.Marshal(SlackUpdateLeaveDayRequestMessageTaskPayload{Request: request}); err != nil {
		return nil, err
	} else {
		return asynq.NewTask(SlackUpdateLeaveDayRequestMessageTask, payload), nil
	}
}

func HandleSlackUpdateLeaveDayRequestMessageTask(ctx context.Context, t *asynq.Task) error {
	db := database.Db

	var p SlackUpdateLeaveDayRequestMessageTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	// Get message content
	request := p.Request
	db.Model(&request).First(&request)
	messageText := request.GetMessage(db, nil)

	// Get request message
	message := models.Message{ParentId: request.Id}
	messageRepo := repository.NewMessageRepository(&ctx, db)
	err := messageRepo.Find(&message)

	if err != nil {
		return err
	}

	// Update slack message
	slackClient := models.NewSlackClient()
	slackMessage, slackErr := slackClient.UpdateMessage(os.Getenv("SLACK_LEAVE_WFH_REQUEST_CHANNEL_ID"), *message.Timestamp, messageText, nil)
	if slackErr != nil {
		return slackErr
	}

	// Update request message
	message.Content = &messageText
	message.Timestamp = &slackMessage.Ts
	messageRepo.Update(&message)

	return nil
}
