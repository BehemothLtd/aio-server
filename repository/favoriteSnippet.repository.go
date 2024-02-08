package repository

import (
	"aio-server/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

func NewFavoriteSnippetRepository(c *context.Context, db *gorm.DB) *Repository {
	return &Repository{
		db: db,
		c:  c,
	}
}

func (r *Repository) FindByUserAndSnippet(
	snippetFavorited *models.SnippetsFavorite,
) (favoritedRec *models.SnippetsFavorite, err error) {
	dbTables := r.db.Table("snippets_favorites")

	result := dbTables.Where(&snippetFavorited).First(&favoritedRec)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return favoritedRec, nil
}

func (r *Repository) ToggleFavoriteSnippet(
	snippet *models.Snippet,
	user *models.User,
) (favorited bool, err error) {
	dbTables := r.db.Table("snippets_favorites")

	snippetFavorited := models.SnippetsFavorite{UserId: user.Id, SnippetId: snippet.Id}

	found, err := r.FindByUserAndSnippet(&snippetFavorited)

	if err != nil {
		return false, err
	}

	if found != nil {
		// Found -> delete -> unfavorited
		deleteResult := dbTables.Where(&snippetFavorited).Delete(&snippetFavorited)

		if deleteResult.Error != nil {
			return false, deleteResult.Error
		}

		return false, nil
	} else {
		// Not Found -> create -> favorited
		createResult := dbTables.Create(&snippetFavorited)

		if createResult.Error != nil {
			return false, createResult.Error
		}

		return true, nil
	}
}
