package insightResolvers

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/gqlTypes/globalTypes"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"

	"gorm.io/gorm"
)

func (r *Resolver) Client(ctx context.Context, args insightInputs.ClientInput) (*globalTypes.ClientType, error) {
	if _, err := r.Authorize(ctx, enums.PermissionTargetTypeClients.String(), enums.PermissionActionTypeRead.String()); err != nil {
		return nil, err
	}

	if args.Id == "" {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	clientId, err := helpers.GqlIdToInt32(args.Id)

	if err != nil {
		return nil, err
	}

	client := models.Client{Id: clientId}
	repo := repository.NewClientRepository(&ctx, r.Db.Preload("Avatar", "name = 'avatar'").Preload("Avatar.AttachmentBlob"))
	err = repo.Find(&client)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exceptions.NewRecordNotFoundError()
		}
		return nil, err
	}

	return &globalTypes.ClientType{Client: &client}, nil
}
