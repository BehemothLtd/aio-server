package resolvers

import (
	"aio-server/gql/inputs/msInputs"
	"aio-server/services/msServices"
	"context"
)

func (r *Resolver) MsSnippetDecryptContent(ctx context.Context, args msInputs.SnippetDecryptContentInput) (*string, error) {
	service := msServices.SnippetDecryptService{
		Ctx:  &ctx,
		Db:   r.Db,
		Args: args,
	}

	if decryptedContent, err := service.Execute(); err != nil {
		return nil, err
	} else {
		return decryptedContent, nil
	}
}
