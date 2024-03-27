package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"context"
)

func (r *Resolver) LeaveDayRequestStateChange(ctx context.Context, args insightInputs.LeaveDayRequestStateChangeInput) (*string, error) {
	user, err := r.Authorize(ctx, string(enums.PermissionTargetTypeLeaveDayRequests), string(enums.PermissionActionTypeChangeState))

	if err != nil {
		return nil, err
	}

	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}
	if args.RequestState != enums.RequestStateType()
}
