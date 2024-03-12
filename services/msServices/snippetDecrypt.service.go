package msServices

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/gql/inputs/msInputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/pkg/helpers"
	"aio-server/pkg/specialTypes"
	"aio-server/repository"
	"context"

	"gorm.io/gorm"
)

// SnippetDecryptService handles correct snippet's user to decrypt their own
// snippet content using a Passkey
type SnippetDecryptService struct {
	Ctx  *context.Context
	Db   *gorm.DB
	Args msInputs.SnippetDecryptContentInput

	snippet models.Snippet
}

func (sds *SnippetDecryptService) Execute() (*string, error) {
	if err := sds.validate(); err != nil {
		return nil, err
	}

	decyptedContent, err := sds.snippet.DecryptContent(sds.Args.Passkey)
	if err != nil {
		return nil, err
	}

	return decyptedContent, nil
}

// validate validates the input and retrieves the user information.
func (sds *SnippetDecryptService) validate() error {
	if sds.Args.Id == "" {
		return exceptions.NewBadRequestError("Invalid Id")
	}

	snippetId, err := helpers.GqlIdToInt32(sds.Args.Id)
	if err != nil {
		return exceptions.NewBadRequestError("Invalid Id")
	}

	// Authenticate user
	user, err := auths.AuthUserFromCtx(*sds.Ctx)
	if err != nil {
		return exceptions.NewUnauthorizedError("")
	}

	sds.snippet = models.Snippet{
		Id: snippetId,
	}

	// Retrieve the snippet
	snippetRepo := repository.NewSnippetRepository(sds.Ctx, sds.Db)
	if err := snippetRepo.FindById(&sds.snippet, sds.snippet.Id); err != nil {
		return exceptions.NewRecordNotFoundError()
	}

	if user.Id != sds.snippet.UserId {
		return exceptions.NewRecordNotFoundError()
	}

	if sds.snippet.SnippetType == enums.SnippetTypePublic {
		return exceptions.NewUnprocessableContentError("Unable to perform this action", nil)
	}

	if sds.Args.Passkey == "" {
		return exceptions.NewUnprocessableContentError("Passkey is required", exceptions.ResourceModificationError{
			"passkey": &specialTypes.FieldAttributeErrorType{
				Base: []string{"Cant be empty"},
			},
		})
	}

	return nil
}
