package repository

import (
	"aio-server/models"
	"context"

	"gorm.io/gorm"
)

type AttachmentBlobRepository struct {
	Repository
}

// NewAttachmentBlobRepository initializes a new PinRepository instance.
func NewAttachmentBlobRepository(c *context.Context, db *gorm.DB) *AttachmentBlobRepository {
	return &AttachmentBlobRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

// Find find blob record by its attributes
func (r *AttachmentBlobRepository) Find(blob *models.AttachmentBlob) error {
	dbTables := r.db.Table("attachment_blobs")

	return dbTables.Where(&blob).First(&blob).Error
}
