package insightServices

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"aio-server/validators"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type AttendanceCreateService struct {
	Ctx        *context.Context
	Db         *gorm.DB
	Args       insightInputs.AttendanceCreateInput
	Attendance *models.Attendance
}

func (acs *AttendanceCreateService) Execute() error {
	form := validators.NewAttendanceFormValidator(
		&acs.Args.Input,
		repository.NewAttendanceRepository(acs.Ctx, acs.Db),
		acs.Attendance,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
