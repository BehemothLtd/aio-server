package models

import "time"

type SnippetsCollection struct {
	Id           int32 `gorm:"not null;autoIncrement"`
	SnippetId    int32 `gorm:"not null"`
	CollectionId int32 `gorm:"not null"`
	Snippet      Snippet
	Collection   Collection
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
