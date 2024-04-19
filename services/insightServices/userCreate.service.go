package insightServices

import (
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/utilities"
	"aio-server/repository"
	"aio-server/validators"
	"fmt"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type UserCreateService struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args insightInputs.UserFormInput
	User *models.User
}

func (uc *UserCreateService) Execute() (string, error) {
	password := utilities.RandomToken(8)
	uc.Args.Password = &password

	form := validators.NewUserCreateFormValidator(
		&uc.Args,
		repository.NewUserRepository(uc.Ctx, uc.Db),
		uc.User,
	)

	if err := form.Save(); err != nil {
		return "", err
	}

	return fmt.Sprintf("Email: %s, Password: %s", password, *uc.Args.Email), nil
}
