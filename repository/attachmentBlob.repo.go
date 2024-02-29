package repository

import (
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
