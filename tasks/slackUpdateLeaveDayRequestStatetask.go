package tasks

import (
	"aio-server/database"
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/models"
	"aio-server/repository"
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"

	"github.com/hibiken/asynq"
)

type SlackUpdateLeaveDayRequestStateTaskPayload struct {
	Payload models.SlackInteractivePayload
	User    models.User
}

func NewSlackUpdateLeaveDayRequestStateTask(interactivePayload models.SlackInteractivePayload, user models.User) (*asynq.Task, error) {
	if payload, err := json.Marshal(SlackUpdateLeaveDayRequestStateTaskPayload{Payload: interactivePayload, User: user}); err != nil {
		return nil, err
	} else {
		return asynq.NewTask(SlackUpdateLeaveDayRequestStateTask, payload), nil
	}
}

func HandleSlackUpdateLeaveDayRequestStateTask(ctx context.Context, t *asynq.Task) error {
	db := database.Db

	var p SlackUpdateLeaveDayRequestStateTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	params := p.Payload
	message := params.OriginalMessage
	originalText := message.Text
	action := params.Action[0].Value
	re := regexp.MustCompile(`id=(\+\d+)`)
	match := re.FindStringSubmatch(originalText)

	if len(match) > 1 {
		matchId := match[1]
		requestId, err := strconv.ParseInt(matchId, 10, 32)
		if err != nil {
			return err
		}

		request := models.LeaveDayRequest{Id: int32(requestId)}
		repo := repository.NewLeaveDayRequestRepository(&ctx, db)
		if err := repo.Find(&request); err != nil {
			return err
		}

		request.ApproverId = &p.User.Id
		state, err := enums.ParseRequestStateType(action)
		if err != nil {
			return err
		}

		request.RequestState = state
		repo.Update(&request)

		return nil
	} else {
		return exceptions.NewBadRequestError("Invalid request value")
	}
}
