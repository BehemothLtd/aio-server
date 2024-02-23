package msServices

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/msInputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"aio-server/validators"
	"context"

	"gorm.io/gorm"
)

// TagUpdateService handles the updation of tag
type TagUpdateService struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args msInputs.TagUpdateInput
}

// Execute updates a new tag.
func (tus *TagUpdateService) Execute() (*models.Tag, error) {
	// Authenticate the user
	user, err := auths.AuthUserFromCtx(*tus.Ctx)
	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	tagRepo := repository.NewTagRepository(tus.Ctx, tus.Db)

	tagId, err := helpers.GqlIdToInt32(tus.Args.Id)

	if err != nil {
		return nil, exceptions.NewBadRequestError("Invalid Id")
	}

	tag := models.Tag{}
	snippetErr := tagRepo.FindById(&tag, tagId)

	if snippetErr != nil {
		return nil, exceptions.NewRecordNotFoundError()
	}

	if tag.UserId != user.Id {
		return nil, exceptions.NewRecordNotFoundError()
	}

	form := validators.NewTagFormValidator(
		&tus.Args.Input,
		repository.NewTagRepository(tus.Ctx, tus.Db),
		&tag,
	)

	if formErr := form.Save(); formErr != nil {
		return nil, formErr
	} else {
		return &tag, nil
	}
}
