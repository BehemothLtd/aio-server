package insightResolvers

import (
	"aio-server/enums"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/services/insightServices"
	"context"
)

func (r *Resolver) AttendanceCreate(ctx context.Context, args insightInputs.AttendanceCreateInput) (*insightTypes.AttendanceType, error) {
	currentUser, err := r.Authorize(ctx, enums.PermissionTargetTypeClients.String(), enums.PermissionActionTypeWrite.String())
	if err != nil {
		return nil, err
	}

	attendance := models.Attendance{CreatedUserId: currentUser.Id}
	service := insightServices.AttendanceCreateService{
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
