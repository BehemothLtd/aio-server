package services

import (
	"aio-server/exceptions"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"

	"gorm.io/gorm"
)

// AuthService handles user authentication.
type AuthService struct {
	Email    string
	Password string
	Ctx      *context.Context
	Db       *gorm.DB

	Token *string
}

// Execute performs the authentication process.
func (a *AuthService) Execute() error {
	if err := a.validate(); err != nil {
		return err
	}
	user, err := a.findUser()
	if err != nil {
		return err
	}
	token, _ := helpers.GenerateJwtToken(user.GenerateJwtClaims())
	a.Token = &token
	return nil
}

// validate checks if the email and password are provided.
func (a *AuthService) validate() error {
	exception := exceptions.NewUnprocessableContentError("", nil)
	if a.Email == "" {
		exception.AddBaseError("email", []string{"cannot be empty"})
	}
	if a.Password == "" {
		exception.AddBaseError("password", []string{"cannot be empty"})
	}
	if len(exception.Errors) > 0 {
		return exception
	}
	return nil
}

// findUser tries to authenticate the user.
func (a *AuthService) findUser() (*models.User, error) {
	repo := repository.NewUserRepository(a.Ctx, a.Db)
	if user, err := repo.Auth(a.Email, a.Password); err != nil {
		exception := exceptions.NewUnprocessableContentError("Unauthorized", nil)
		exception.AddBaseError("base", []string{"User Or Password is incorrect"})

		return nil, exception
	} else {
		return user, nil
	}
}
