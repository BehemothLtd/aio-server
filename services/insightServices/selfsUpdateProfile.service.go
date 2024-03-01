package insightServices

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type SelfsUpdateProfileService struct {
	Ctx *context.Context
	Db  *gorm.DB

	Args insightInputs.SelfsUpdateProfileInput
	User *models.User
}

func (sups *SelfsUpdateProfileService) Execute() error {
	// TODO
	sups.validate()

	return nil
}

func (sups *SelfsUpdateProfileService) validate() error {
	exception := exceptions.NewUnprocessableContentError("", nil)

	if len(exception.Errors) > 0 {
		return exception
	}
	return nil
}
