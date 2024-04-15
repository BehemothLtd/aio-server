package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/services/insightServices"
	"context"
)

func (r *Resolver) AttendanceUpdate(ctx context.Context, args insightInputs.AttendanceUpdateInput) (*insightTypes.AttendanceType, error) {
	currentUser, err := r.Authorize(ctx, enums.PermissionTargetTypeClients.String(), enums.PermissionActionTypeWrite.String())
	if err != nil {
		return nil, err
	}

	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	attendanceId, err := helpers.GqlIdToInt32(args.Id)
	if err != nil {
		return nil, err
	}

	attendance := models.Attendance{Id: attendanceId, CreatedUserId: currentUser.Id}

	service := insightServices.AttendanceUpdateService{
		Ctx:        &ctx,
		Db:         r.Db,
		Args:       args,
		Attendance: &attendance,
	}

	if err := service.Execute(); err != nil {
		return nil, err
	} else {
		return &insightTypes.AttendanceType{
			Attendance: &globalTypes.AttendanceType{
				Attendance: &attendance,
			},
		}, nil
	}
}
