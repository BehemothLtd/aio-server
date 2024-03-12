package insightServices

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"aio-server/validators"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type SelfUpdateProfileService struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args insightInputs.SelfUpdateProfileInput
	User *models.User
}

func (sups *SelfUpdateProfileService) Execute() error {
	repo := repository.NewUserRepository(sups.Ctx, sups.Db)
	repo.FindWithAvatar(sups.User)

	form := validators.NewUserProfileFormValidator(
		&sups.Args.Input,
		repo,
		sups.User,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
