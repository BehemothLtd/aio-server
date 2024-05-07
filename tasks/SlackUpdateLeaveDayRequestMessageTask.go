package tasks

import (
	"aio-server/database"
	"aio-server/models"
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

	request := p.Request
	db.Model(&request).First(&request)
	messageText := request.GetMessage(db, nil)

	message := models.RequestMessage{Id: request.Id}
	db.Model(&message).First(&message)

	slackClient := models.NewSlackClient()

	err := slackClient.UpdateMessage(os.Getenv("SLACK_LEAVE_WFH_REQUEST_CHANNEL_ID"), *message.Timestamp, messageText, nil)
	if err != nil {
		return err
	}

	return nil
}
