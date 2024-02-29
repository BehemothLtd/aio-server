package models

import "time"

type AttachmentBlob struct {
	Id        int32
	Filename  string
	Key       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
