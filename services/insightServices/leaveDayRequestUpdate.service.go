package insightServices

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"aio-server/validators"
	"context"

	"gorm.io/gorm"
)

type LeaveDayRequestUpdateService struct {
	Ctx     *context.Context
	Db      *gorm.DB
	Args    insightInputs.LeaveDayRequestUpdateInput
	Request *models.LeaveDayRequest
}

func (rus *LeaveDayRequestUpdateService) Execute() error {
	repo := repository.NewLeaveDayRequestRepository(rus.Ctx, rus.Db)

	if err := repo.Find(rus.Request); err != nil {
		return exceptions.NewRecordNotFoundError()
	}

	if rus.Request.RequestState != enums.RequestStateTypePending {
		return exceptions.NewBadRequestError("Can not update this request!")
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
