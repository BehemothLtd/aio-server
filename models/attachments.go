package models

import (
	"fmt"
	"os"
	"time"
)

type Attachment struct {
	Id               int32
	OwnerID          int32 `gorm:"column:owner_id"`
	OwnerType        string
	AttachmentBlobId int32
	AttachmentBlob   *AttachmentBlob

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a Attachment) Url() *string {
	if a.AttachmentBlob != nil {
		var url string

		key := a.AttachmentBlob.Key

		if os.Getenv("UPLOAD_LOCALLY_PATH") != "" {
			url = os.Getenv("UPLOAD_LOCALLY_PATH") + key
		} else {
			bucketName := os.Getenv("GCS_BUCKET_NAME")
			url = fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, key)
		}

		return &url
	}

	return nil
}
