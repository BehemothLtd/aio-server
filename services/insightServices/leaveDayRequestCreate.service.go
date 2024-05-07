package insightServices

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"aio-server/tasks"
	"aio-server/validators"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type LeaveDayRequestService struct {
	Ctx     *context.Context
	Db      *gorm.DB
	Args    insightInputs.LeaveDayRequestCreateInput
	Request *models.LeaveDayRequest
}

func (rs *LeaveDayRequestService) Excecute() error {
	form := validators.NewLeaveDayrequestFormValidator(
		&rs.Args.Input,
		repository.NewLeaveDayRequestRepository(rs.Ctx, rs.Db),
		rs.Request,
	)

	if err := form.Save(); err != nil {
		return err
	}

	// Send slack message task
	mentions := rs.Args.Input.Mentions
	task, err := tasks.NewSlackSendLeaveDayRequestTask(rs.Request.Id, mentions)
	if err != nil {
		return err
	}

	info, err := tasks.AsynqClient.Enqueue(task)
	if err != nil {
		return err
	}

	fmt.Printf("Task ID: %+v => completed at %v\n", info.ID, info.CompletedAt)

	return nil
}
