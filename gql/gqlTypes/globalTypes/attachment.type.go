package globalTypes

import (
	"aio-server/models"
	"context"
)

type AttachmentType struct {
	Attachment *models.Attachment
}

func (at *AttachmentType) Key(ctx context.Context) *string {
	return nil
}

func (at *AttachmentType) Url(ctx context.Context) *string {
	return nil
}
