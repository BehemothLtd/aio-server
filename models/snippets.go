package models

import (
	"time"
)

type Snippet struct {
	Id             int32
	Title          string
	Content        string
	UserId         int32
	Slug           string
	SnippetType    int
	FavoritesCount int
	FavoritedUsers []*User `gorm:"many2many:snippets_favorites"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type SnippetsQuery struct {
	TitleCont string
}

type SnippetsCollection struct {
	Collection []*Snippet
	Metadata   *Metadata
}
