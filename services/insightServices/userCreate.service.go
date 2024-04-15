package insightServices

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"aio-server/validators"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type UserCreateService struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args insightInputs.UserFormInput
	User *models.User
}

func (uc *UserCreateService) Execute() error {
	form := validators.NewUserFormValidator(
		&uc.Args,
		repository.NewUserRepository(uc.Ctx, uc.Db),
		uc.User,
	)

	if err := form.Save(); err != nil {
		return err
	}

	return nil
}
