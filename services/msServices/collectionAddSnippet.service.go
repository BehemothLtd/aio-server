package msServices

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/msInputs"
	"aio-server/models"
	"aio-server/pkg/auths"
	"aio-server/pkg/helpers"
	"aio-server/repository"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type CollectionAddSnippetService struct {
	Ctx  context.Context
	Db   gorm.DB
	Args msInputs.CollectionAddSnippetInput

	collection *models.Collection
	snippet    *models.Snippet
}

func (sats *CollectionAddSnippetService) Execute() (bool, error) {
	if err := sats.validate(); err != nil {
		return false, err
	}

	snippetId, _ := helpers.GqlIdToInt32(sats.Args.SnippetId)
	collectionId, _ := helpers.GqlIdToInt32(sats.Args.Id)

	snippetsCollection := models.SnippetsCollection{
		SnippetId:    snippetId,
		CollectionId: collectionId,
	}

	repo := repository.NewSnippetsCollectionRepository(&sats.Ctx, &sats.Db)

	if err := repo.Find(&snippetsCollection); err != nil {
		return false, exceptions.NewUnprocessableContentError("already has this snippet", nil)
	} else {
		if err := repo.Create(&snippetsCollection); err != nil {
			return false, exceptions.NewUnprocessableContentError(fmt.Sprintf("false to add snippet to collection %s", err.Error()), nil)
		} else {
			return true, nil
		}
	}
}

func (sats *CollectionAddSnippetService) validate() error {
	// Check auth user
	// Authenticate the user
	user, err := auths.AuthUserFromCtx(sats.Ctx)
	if err != nil {
		return exceptions.NewUnauthorizedError("")
	}

	// check snippet valid
	snippetId, err := helpers.GqlIdToInt32(sats.Args.Id)
	if err != nil {
		return err
	}

	snippet := models.Snippet{Id: snippetId}
	snippetRepo := repository.NewSnippetRepository(&sats.Ctx, &sats.Db)
	err = snippetRepo.FindSnippetByAttr(&snippet)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return exceptions.NewRecordNotFoundError()
		}
		return err
	}

	// check collection valid
	collectionId, err := helpers.GqlIdToInt32(sats.Args.Id)
	if err != nil {
		return err
	}

	collection := models.Collection{Id: collectionId}
	collectionRepo := repository.NewCollectionRepository(&sats.Ctx, &sats.Db)
	err = collectionRepo.Find(&collection)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return exceptions.NewRecordNotFoundError()
		}
		return err
	}

	// check self collection
	if collection.UserId != user.Id {
		return exceptions.NewRecordNotFoundError()
	}

	sats.snippet = &snippet
	sats.collection = &collection

	return nil
}
