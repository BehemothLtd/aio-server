package globalTypes

import (
	"aio-server/models"
	"context"
)

type AttachmentType struct {
	Attachment *models.Attachment
}

func (at *AttachmentType) Key(ctx context.Context) *string {
	if at.Attachment.AttachmentBlob != nil {
		return &at.Attachment.AttachmentBlob.Key
	}

	return nil
}

func (at *AttachmentType) Url(ctx context.Context) *string {
	url := at.Attachment.Url()

	return url
}
