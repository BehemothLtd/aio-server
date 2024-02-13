package models

import (
	"time"
)

type Snippet struct {
	Id             int32  `gorm:"not null;autoIncrement"`
	Title          string `gorm:"not null;type:varchar(255);default:null"`
	Content        string `gorm:"not null;type:longtext;default:null"`
	UserId         int32  `gorm:"not null;type:bigint;default:null"`
	Slug           string
	SnippetType    int    `gorm:"not null;default:1"`
	FavoritesCount int    `gorm:"not null;default:0"`
	FavoritedUsers []User `gorm:"many2many:snippets_favorites"`
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
