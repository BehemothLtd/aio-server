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
func (r *SnippetFavoriteRepository) Toggle(snippet *models.Snippet, user *models.User) (result *models.SnippetFavorited, err error) {
	dbTables := r.db.Table("snippets_favorites")

	snippetFavorited := models.SnippetsFavorite{UserId: user.Id, SnippetId: snippet.Id}

	found, err := r.FindByUserAndSnippet(&snippetFavorited)

	if err != nil {
		return nil, err
	}

	if found != nil {
		// Found -> delete -> unfavorited
		deleteResult := dbTables.Where(&snippetFavorited).Delete(&snippetFavorited)

		if deleteResult.Error != nil {
			return nil, deleteResult.Error
		}

		return &models.SnippetFavorited{
			Id:             snippet.Id,
			Favorited:      false,
			FavoritesCount: int32(snippet.FavoritesCount - 1),
		}, nil
	}

	// Not Found -> create -> favorited
	createResult := dbTables.Create(&snippetFavorited)
	if createResult.Error != nil {
		return nil, createResult.Error
	}

	return &models.SnippetFavorited{
		Id:             snippet.Id,
		Favorited:      true,
		FavoritesCount: int32(snippet.FavoritesCount + 1),
	}, nil
}
