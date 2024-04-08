package tasks

import (
	"aio-server/models"
	"context"
	"encoding/json"

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

	return nil
}
