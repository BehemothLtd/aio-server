package insightServices

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"aio-server/validators"
	"context"

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

	return nil
}
