package globalTypes

import (
	"aio-server/models"
	"context"
)

type AttachmentType struct {
	Attachment *models.Attachment
}

func (at *AttachmentType) Key(ctx context.Context) *string {
	return &at.Attachment.AttachmentBlob.Key
}

func (at *AttachmentType) Url(ctx context.Context) *string {
	url := at.Attachment.Url()

	return url
}
