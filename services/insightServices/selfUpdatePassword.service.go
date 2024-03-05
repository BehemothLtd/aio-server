package insightServices

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/repository"
	"aio-server/validators"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type SelfUpdatePasswordService struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args insightInputs.SelfUpdatePasswordInput
	User models.User
}

func (sups *SelfUpdatePasswordService) Execute() error {
	form := validators.NewUserPasswordFormValidator(
		&sups.Args.Input,
		repository.NewUserRepository(sups.Ctx, sups.Db),
		&sups.User,
	)

	if err := form.Save(); err != nil {
		return err
	} else {
		return nil
	}
}
