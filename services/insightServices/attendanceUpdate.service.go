package insightServices

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"aio-server/validators"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type AttendanceUpdateService struct {
	Ctx        *context.Context
	Db         *gorm.DB
	Args       insightInputs.AttendanceUpdateInput
	Attendance *models.Attendance
}

func (aus *AttendanceUpdateService) Execute() error {
	repo := repository.NewAttendanceRepository(aus.Ctx, aus.Db)

	if err := repo.Find(aus.Attendance); err != nil {
		return exceptions.NewRecordNotFoundError()
	}

	form := validators.NewAttendanceFormValidator(
		&aus.Args.Input,
		repository.NewAttendanceRepository(aus.Ctx, aus.Db),
		aus.Attendance,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
