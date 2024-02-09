package services

import (
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

	Result *models.Authentication
}

func (a *AuthService) Execute() {
	a.validate()

	if len(a.Result.Errors) > 0 {
		return
	}

	user := a.findUser()
	if user != nil {
		token := a.genarateToken(user)

		if token != nil {
			a.Result.Token = *token
			a.Result.Message = "Successfully Signed In"
		}
	}
}

func (a *AuthService) validate() {
	if a.Email == "" {
		a.Result.Errors = append(a.Result.Errors, &models.ResourceModifyErrors{
			Column: "email",
			Errors: []string{"Email cant be empty"},
		})
	}

	if a.Password == "" {
		a.Result.Errors = append(a.Result.Errors, &models.ResourceModifyErrors{
			Column: "email",
			Errors: []string{"Password cant be empty"},
		})
	}
}

func (a *AuthService) findUser() *models.User {
	repo := repository.NewUserRepository(a.Ctx, a.Db)

	user, err := repo.AuthUser(a.Email, a.Password)
	if err != nil {
		a.Result.Errors = append(a.Result.Errors, &models.ResourceModifyErrors{
			Column: "base",
			Errors: []string{err.Error()},
		})

		return nil
	}

	return user
}

func (a *AuthService) genarateToken(user *models.User) *string {
	token, err := helpers.GenerateJwtToken(user.GenerateJwtClaims())

	if err != nil {
		a.Result.Errors = append(a.Result.Errors, &models.ResourceModifyErrors{
			Column: "base",
			Errors: []string{"error with decode password"},
		})
	}

	return &token
}
