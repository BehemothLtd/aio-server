package insightResolvers

import (
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/gqlTypes/insightTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/repository"
	"context"
)

func (r *Resolver) SelfAttendances(ctx context.Context, args insightInputs.SelfAttendancesInput) (*insightTypes.AttendancesType, error) {
	user, err := auths.AuthInsightUserFromCtx(ctx)

	if err != nil {
		return nil, err
	}

	var attendances []*models.Attendance
	query, paginationData := args.ToPaginationAndQueryData()

	repo := repository.NewAttendanceRepository(&ctx, r.Db)
	if err := repo.ListByUser(&attendances, &paginationData, query, user); err != nil {
		return nil, err
	}

	return &insightTypes.AttendancesType{
		Collection: r.AttendanceSliceToType(attendances),
		Metadata: &globalTypes.MetadataType{
			Metadata: &paginationData.Metadata,
		},
	}, nil
}
