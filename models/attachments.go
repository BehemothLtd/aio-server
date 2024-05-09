package models

import (
	"aio-server/pkg/utilities"
	"time"
)

type Attachment struct {
	Id               int32
	OwnerID          int32 `gorm:"column:owner_id"`
	OwnerType        string
	AttachmentBlobId int32
	AttachmentBlob   AttachmentBlob
	Name             string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a Attachment) Url() *string {
	key := a.AttachmentBlob.Key

	return utilities.GetAvatarUrl(key)
}
