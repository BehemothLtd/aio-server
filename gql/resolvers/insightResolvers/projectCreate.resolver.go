package insightResolvers

import (
	"aio-server/enums"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs/insightInputs"
	"context"
)

func (r *Resolver) ProjectCreate(ctx context.Context, args insightInputs.ProjectCreateInput) (*globalTypes.ProjectType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeUserGroups.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}

	return nil, nil
}
