package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"
	"fmt"
)

func (r *Resolver) AttendanceDelete(ctx context.Context, args insightInputs.AttendanceInput) (*string, error) {
	_, err := r.Authorize(ctx, string(enums.PermissionTargetTypeAttendances), string(enums.PermissionActionTypeWrite))
	if err != nil {
		return nil, err
	}
	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}
	Id, err := helpers.GqlIdToInt32(args.Id)
	if err != nil {
		return nil, err
	}

	attendance := models.Attendance{
		Id: Id,
	}
	repo := repository.NewAttendanceRepository(&ctx, r.Db)

	if err := repo.Find(&attendance); err != nil {
		return nil, exceptions.NewRecordNotFoundError()
	}

	if err = repo.Destroy(&attendance); err != nil {
		return nil, exceptions.NewBadRequestError(fmt.Sprintf("Can not delete this attendance %s", err.Error()))
	} else {
		message := "Deleted"
		return &message, nil
	}

}
