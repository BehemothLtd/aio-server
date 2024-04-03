package tasks

import (
	"aio-server/database"
	"aio-server/models"
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

type DemoTaskPayload struct {
	UserId int32
}

//----------------------------------------------
// Write a function NewXXXTask to create a task.
// A task consists of a type and a payload.
//----------------------------------------------

func NewDemoTask(userId int32) (*asynq.Task, error) {
	if payload, err := json.Marshal(DemoTaskPayload{UserId: userId}); err != nil {
		return nil, err
	} else {
		return asynq.NewTask(TypeDemoTask, payload), nil
	}
}

//---------------------------------------------------------------
// Write a function HandleXXXTask to handle the input task.
// Note that it satisfies the asynq.HandlerFunc interface.
//
//---------------------------------------------------------------

func HandleDemoTask(ctx context.Context, t *asynq.Task) error {
	db := database.Db

	var p DemoTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	user := models.User{Id: p.UserId}
	db.Model(&user).First(&user)

	db.Model(&user).Updates(&models.User{Name: "bachdx"})

	return nil
}
