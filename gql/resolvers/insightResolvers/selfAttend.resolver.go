package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/repository"
	"context"
	"time"
)

func (r *Resolver) SelfAttend(ctx context.Context) (*globalTypes.AttendanceType, error) {
	user, err := auths.AuthUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	currentTime := time.Now()

	repo := repository.NewAttendanceRepository(&ctx, r.Db)
	attendance := models.Attendance{UserId: user.Id, CreatedUserId: user.Id}

	if err := repo.AtDateOfUser(&attendance, user, currentTime); err != nil {
		// Didnt checkin yet, execute checkin
		attendance.CheckinAt = &currentTime

		if createErr := repo.Create(&attendance); createErr != nil {
			return nil, exceptions.NewBadRequestError(createErr.Error())
		} else {
			return &globalTypes.AttendanceType{
				Attendance: &attendance,
			}, nil
		}
	} else {
		// checked in
		// if already checkedOut -> error
		if attendance.CheckoutAt != nil {
			return nil, exceptions.NewBadRequestError("Cant perform this action")
		}
		// if didnt checkout -> execute checkout
		attendanceUpdate := map[string]interface{}{
			"CheckoutAt": currentTime,
		}
		if updateErr := repo.Update(&attendance, attendanceUpdate); updateErr != nil {
			return nil, exceptions.NewBadRequestError(updateErr.Error())
		}

		return &globalTypes.AttendanceType{
			Attendance: &attendance,
		}, nil
	}
}
