package services

import (
	"aio-server/exceptions"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"

	"gorm.io/gorm"
)

type AuthService struct {
	Email    string
	Password string
	Ctx      *context.Context
	Db       *gorm.DB

	Token *string
}

func (a *AuthService) Execute() error {
	validationErr := a.validate()

	if validationErr != nil {
		return validationErr
	}
	user, userErr := a.findUser()

	if userErr != nil {
		return userErr
	}

	token := a.generateToken(user)

	if token != nil {
		a.Token = token
	}

	return nil
}

func (a *AuthService) validate() error {
	exception := exceptions.NewUnprocessableContentError(nil, nil)

	if a.Email == "" {
		exception.AddError(exceptions.ResourceModifyErrors{
			Field:  "email",
			Errors: []string{"Cant be empty"},
		})
	}

	if a.Password == "" {
		exception.AddError(exceptions.ResourceModifyErrors{
			Field:  "password",
			Errors: []string{"Cant be empty"},
		})
	}

	if len(exception.Errors) > 0 {
		return exception
	}

	return nil
}

func (a *AuthService) findUser() (*models.User, error) {
	exception := exceptions.NewUnprocessableContentError(nil, nil)

	repo := repository.NewUserRepository(a.Ctx, a.Db)

	user, err := repo.AuthUser(a.Email, a.Password)

	if err != nil {
		exception.AddError(exceptions.ResourceModifyErrors{
			Field:  "base",
			Errors: []string{err.Error()},
		})

		return nil, exception
	}

	return user, nil
}

func (a *AuthService) generateToken(user *models.User) *string {
	token, _ := helpers.GenerateJwtToken(user.GenerateJwtClaims())

	return &token
}
