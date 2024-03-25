package insightServices

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"aio-server/validators"
	"context"

	"gorm.io/gorm"
)

type LeaveDayRequestUpdateServive struct {
	Ctx     *context.Context
	Db      *gorm.DB
	Args    insightInputs.LeaveDayRequestUpdateInput
	Request *models.LeaveDayRequest
}

func (rus *LeaveDayRequestUpdateServive) Execute() error {
	repo := repository.NewLeaveDayRequestRepository(rus.Ctx, rus.Db.Preload("User"))

	if err := repo.Find(rus.Request); err != nil {
		return exceptions.NewRecordNotFoundError()
	}

	form := validators.NewLeaveDayrequestFormValidator(
		&rus.Args.Input,
		repo,
		rus.Request,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
