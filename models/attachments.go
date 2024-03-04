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

func (a Attachment) Url() string {
	key := a.AttachmentBlob.Key

	if os.Getenv("UPLOAD_LOCALLY_PATH") != "" {
		return os.Getenv("UPLOAD_LOCALLY_PATH") + key
	} else {
		bucketName := os.Getenv("GCS_BUCKET_NAME")
		return fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, key)
	}
}
