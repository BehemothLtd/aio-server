package resolvers

import (
	"aio-server/gql/inputs"
	"aio-server/services"
	"context"
)

func (r *Resolver) MsSnippetDecryptContent(ctx context.Context, args inputs.MsSnippetDecryptContentInput) (*string, error) {
	service := services.SnippetDecryptService{
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
