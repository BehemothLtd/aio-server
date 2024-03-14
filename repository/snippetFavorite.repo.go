package repository

import (
	"aio-server/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

// SnippetFavoriteRepository handles operations related to snippetFavorites.
type SnippetFavoriteRepository struct {
	Repository
}

// NewSnippetRepository initializes a new NewSnippetFavoriteRepository instance.
func NewSnippetFavoriteRepository(c *context.Context, db *gorm.DB) *SnippetFavoriteRepository {
	return &SnippetFavoriteRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

// FindByUserAndSnippet finds snippetFavorite record by its User and Snippet
func (r *SnippetFavoriteRepository) FindByUserAndSnippet(
	snippetFavorited *models.SnippetsFavorite,
) (favoritedRec *models.SnippetsFavorite, err error) {
	dbTables := r.db.Table("snippets_favorites")

	result := dbTables.Where(&snippetFavorited).First(&favoritedRec)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return favoritedRec, nil
}

// Toggle toggles a snippet favorite.
func (r *SnippetFavoriteRepository) Toggle(snippet *models.Snippet, user *models.User) (result *models.Snippet, err error) {
	err = r.db.Transaction(func(tx *gorm.DB) error {
		snippetFavorited := models.SnippetsFavorite{UserId: user.Id, SnippetId: snippet.Id}
		found, err := r.FindByUserAndSnippet(&snippetFavorited)
		if err != nil {
			return err
		}

		snippetFavoriteDb := tx.Table("snippets_favorites")
		snippetDb := tx.Table("snippets")
		if found != nil {
			// Found -> delete -> unfavorited
			deleteResult := snippetFavoriteDb.Where(&snippetFavorited).Delete(&snippetFavorited)
			if deleteResult.Error != nil {
				return deleteResult.Error
			}

			updateResult := snippetDb.Where(&snippet).UpdateColumn("favorites_count", gorm.Expr("favorites_count - ?", 1))
			if updateResult.Error != nil {
				return updateResult.Error
			}

			result = &models.Snippet{
				Favorited: false,
			}
		} else {
			// Not Found -> create -> favorited
			createResult := snippetFavoriteDb.Create(&snippetFavorited)
			if createResult.Error != nil {
				return createResult.Error
			}

			updateResult := snippetDb.Where(&snippet).UpdateColumn("favorites_count", gorm.Expr("favorites_count + ?", 1))
			if updateResult.Error != nil {
				return updateResult.Error
			}

			result = &models.Snippet{
				Favorited: true,
			}
		}

		return nil
	})

	return result, err
}
