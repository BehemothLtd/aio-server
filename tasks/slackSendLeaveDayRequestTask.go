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

type SlackSendLeaveDayRequestTaskPayload struct {
	RequestId int32
	Mentions  *[]*string
}

func NewSlackSendLeaveDayRequestTask(requestId int32, mentions *[]*string) (*asynq.Task, error) {
	if payload, err := json.Marshal(SlackSendLeaveDayRequestTaskPayload{RequestId: requestId, Mentions: mentions}); err != nil {
		return nil, err
	} else {
		return asynq.NewTask(SlackSendLeaveDayRequestMessagetask, payload), nil
	}
}

func HandleSlackSendLeaveDayRequestTask(ctx context.Context, t *asynq.Task) error {
	db := database.Db

	var p SlackSendLeaveDayRequestTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	request := models.LeaveDayRequest{Id: p.RequestId}
	db.Model(&request).First(&request)
	message := request.GetMessage(db, p.Mentions)
	callback := "change_state_rq"
	slacClient := models.NewSlackClient()

	slacClient.SendMessage(message, os.Getenv("SLACK_LEAVE_WFH_REQUEST_CHANNEL_ID"), &callback)

	return nil
}
