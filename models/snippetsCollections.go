package models

import "time"

type SnippetsCollection struct {
	Id           int32
	SnippetId    int32
	CollectionId int32
	Snippet      Snippet
	Collection   Collection
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
