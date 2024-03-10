package models

import "time"

type SnippetsTag struct {
	Id        int32
	SnippetId int32
	TagId     int32
	CreatedAt time.Time
	UpdatedAt time.Time
}
