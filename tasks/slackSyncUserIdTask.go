package tasks

import (
	"aio-server/database"
	"aio-server/exceptions"
	"aio-server/models"
	"aio-server/repository"
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

type SlackSyncUserIdTaskPayload struct {
	UserId int32
}

func NewSlackSyncUserIdTask(userId int32) (*asynq.Task, error) {
	if payload, err := json.Marshal(SlackSyncUserIdTaskPayload{UserId: userId}); err != nil {
		return nil, err
	} else {
		return asynq.NewTask(SlackSyncUserIdTask, payload), nil
	}
}

func HandleSlackSyncUserIdTask(ctx context.Context, t *asynq.Task) error {
	db := database.Db

	var p SlackSyncUserIdTaskPayload

	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	userRepo := repository.NewUserRepository(&ctx, db)

	user := models.User{Id: p.UserId}
	if err := userRepo.Find(&user); err != nil {
		if err == gorm.ErrRecordNotFound {
			return exceptions.NewRecordNotFoundError()
		}

	}
}
