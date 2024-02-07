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
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
