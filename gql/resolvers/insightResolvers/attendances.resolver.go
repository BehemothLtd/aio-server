package insightResolvers

import (
	"aio-server/enums"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"context"
)

func (r *Resolver) Attendances(ctx context.Context, args insightInputs.AttendancesInput) (*insightTypes.AttendancesType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeAttendances.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}

	var attendances []*models.Attendance
	query, paginationData := args.ToPaginationAndQueryData()

	repo := repository.NewAttendanceRepository(&ctx, r.Db.Preload("User"))
	if err := repo.List(&attendances, &paginationData, query); err != nil {
		return nil, err
	}

	return &insightTypes.AttendancesType{
		Collection: r.AttendanceSliceToType(attendances),
		Metadata: &globalTypes.MetadataType{
			Metadata: &paginationData.Metadata,
		},
	}, nil
}
