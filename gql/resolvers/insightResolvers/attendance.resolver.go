package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"

	"gorm.io/gorm"
)

func (r *Resolver) Attendance(ctx context.Context, args insightInputs.AttendanceInput) (*globalTypes.AttendanceType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeClients.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}

	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	attendanceId, err := helpers.GqlIdToInt32(args.Id)
	if err != nil {
		return nil, err
	}
	attendance := models.Attendance{Id: attendanceId}
	repo := repository.NewAttendanceRepository(&ctx, r.Db)
	err = repo.Find(&attendance)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewRecordNotFoundError()
		}
		return nil, err
	}

	return &globalTypes.AttendanceType{Attendance: &attendance}, nil
}
