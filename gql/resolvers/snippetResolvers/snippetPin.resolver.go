package snippetResolvers

import (
	"aio-server/gql/inputs/msInputs"
	"aio-server/services/msServices"
	"context"
)

func (r *Resolver) SnippetPin(ctx context.Context, args msInputs.SnippetPinInput) (*bool, error) {
	service := msServices.SnippetPinService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if result, err := service.Execute(); err != nil {
		return nil, err
	} else {
		return result, err
	}
}
