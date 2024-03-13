package insightResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/pkg/auths"
	"aio-server/pkg/systems"
	"context"
)

func (r *Resolver) DevelopmentRoles(ctx context.Context) (*[]*globalTypes.DevelopentRoleType, error) {
	_, err := auths.AuthInsightUserFromCtx(ctx)

	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	developmentRoles := systems.GetDevelopmentRoles()

	result := make([]*globalTypes.DevelopentRoleType, len(developmentRoles))

	for i := range developmentRoles {
		result[i] = &globalTypes.DevelopentRoleType{
			DevelopmentRole: &developmentRoles[i],
		}
	}

	return &result, nil
}
