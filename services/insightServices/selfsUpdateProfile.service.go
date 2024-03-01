package insightServices

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"aio-server/validators"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type SelfsUpdateProfileService struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args insightInputs.SelfsUpdateProfileInput
	User *models.User

	repo *repository.UserRepository
}

func (sups *SelfsUpdateProfileService) Execute() error {
	sups.repo = repository.NewUserRepository(sups.Ctx, sups.Db)
	form := validators.NewUserProfileFormValidator(
		&sups.Args.Input,
		sups.repo,
		sups.User,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
