package insightServices

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"aio-server/validators"
	"context"

	"gorm.io/gorm"
)

type UserUpdateService struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args insightInputs.UserUpdateInput
	User *models.User
}

func (us *UserUpdateService) Execute() error {
	repo := repository.NewUserRepository(us.Ctx, us.Db)
	err := repo.FindWithAvatar(us.User)
	if err != nil {
		return err
	}

	form := validators.NewUserUpdateFormValidator(
		&us.Args.Input,
		repo,
		us.User,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
