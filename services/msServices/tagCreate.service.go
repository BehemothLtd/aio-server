package msServices

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/msInputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/repository"
	"aio-server/validators"
	"context"

	"gorm.io/gorm"
)

// TagCreateService handles the creation of tags.
type TagCreateService struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args msInputs.TagCreateInput
}

// Execute creates a new tag.
func (tcs *TagCreateService) Execute() (*models.Tag, error) {
	// Authenticate the user
	user, err := auths.AuthUserFromCtx(*tcs.Ctx)
	if err != nil {
		return nil, exceptions.NewUnauthorizedError("")
	}

	tag := models.Tag{
		UserId: user.Id,
	}

	form := validators.NewTagFormValidator(
		&tcs.Args.Input,
		repository.NewTagRepository(tcs.Ctx, tcs.Db),
		&tag,
	)

	if formErr := form.Save(); formErr != nil {
		return nil, formErr
	} else {
		return &tag, nil
	}
}
