package repository

import (
	"aio-server/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

// PinRepository handles operations related to pins
type PinRepository struct {
	Repository
}

// New PinRepository initializes a new PinRepository instance.
func NewPinRepository(c *context.Context, db *gorm.DB) *PinRepository {
	return &PinRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

// FindByUserAndSnippet finds pin record by its User and Snippet
func (r *PinRepository) FindByUserAndSnippet(
	pin *models.Pin,
) (pinnedRec *models.Pin, err error) {
	dbTables := r.db.Table("pins")

	result := dbTables.Where(&pin).First(&pinnedRec)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return pinnedRec, nil
}

// Toggle toggles a snippet pinned
func (r *PinRepository) Toggle(snippet models.Snippet, user models.User) (result bool, err error) {
	dbTables := r.db.Table("pins")

	pinned := models.Pin{
		ParentType: 1,
		ParentID:   snippet.Id,
		UserId:     user.Id,
	}

	found, err := r.FindByUserAndSnippet(&pinned)

	if err != nil {
		return false, err
	}

	if found != nil {
		// Found -> delete -> unpin
		deleteResult := dbTables.Where(&pinned).Delete(&pinned)

		if deleteResult.Error != nil {
			return false, deleteResult.Error
		}

		return false, nil
	}

	// Not Found -> create -> pinned
	createResult := dbTables.Create(&pinned)
	if createResult.Error != nil {
		return false, createResult.Error
	}

	return true, nil
}
