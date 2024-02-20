package snippetResolvers

import (
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/pkg/auths"
	"context"
)

// SelfInfo resolves the query for retrieving self information.
func (r *Resolver) SelfInfo(ctx context.Context) (*globalTypes.UserType, error) {
	user, err := auths.AuthUserFromCtx(ctx)
	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	return &globalTypes.UserType{
		User: &user,
	}, nil
}
