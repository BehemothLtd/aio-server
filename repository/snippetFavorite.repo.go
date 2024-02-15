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
			db: db,
			ctx: c,
		},
	}
}

// FindByUserAndSnippet finds snippetFavorite record by it User and Snippet
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
func (r *SnippetFavoriteRepository) Toggle(snippet *models.Snippet, user *models.User) (favorited bool, err error) {
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
	}

	// Not Found -> create -> favorited
	createResult := dbTables.Create(&snippetFavorited)
	if createResult.Error != nil {
		return false, createResult.Error
	}

	return true, nil
}
