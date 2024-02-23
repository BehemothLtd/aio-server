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

type SnippetAddTagService struct {
	Ctx  context.Context
	Db   gorm.DB
	Args msInputs.SnippetAddTagInput

	snippet *models.Snippet
	tag     *models.Tag
}

func (sats *SnippetAddTagService) Execute() (bool, error) {
	if err := sats.validate(); err != nil {
		return false, err
	}

	snippetsTag := models.SnippetsTag{
		SnippetId: sats.snippet.Id,
		TagId:     sats.tag.Id,
	}

	repo := repository.NewSnippetsTagRepository(&sats.Ctx, &sats.Db)

	repo.FindBySnippetAndTag(&snippetsTag)

	if snippetsTag.Id != 0 {
		return false, exceptions.NewUnprocessableContentError("already has this tag", nil)
	} else {
		if err := repo.Create(&snippetsTag); err != nil {
			return false, exceptions.NewUnprocessableContentError(fmt.Sprintf("false to add tag to snippet %s", err.Error()), nil)
		} else {
			return true, nil
		}
	}
}

func (sats *SnippetAddTagService) validate() error {
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

	snippet := models.Snippet{}
	snippetRepo := repository.NewSnippetRepository(&sats.Ctx, &sats.Db)
	err = snippetRepo.FindById(&snippet, snippetId)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return exceptions.NewRecordNotFoundError()
		}
		return err
	}
	// check self snippet
	if snippet.UserId != user.Id {
		return exceptions.NewRecordNotFoundError()
	}
	// check tag valid
	tagId, err := helpers.GqlIdToInt32(sats.Args.TagId)
	if err != nil {
		return err
	}

	tag := models.Tag{}
	tagRepo := repository.NewTagRepository(&sats.Ctx, &sats.Db)
	err = tagRepo.FindById(&tag, tagId)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return exceptions.NewRecordNotFoundError()
		}
		return err
	}

	sats.snippet = &snippet
	sats.tag = &tag

	return nil
}
