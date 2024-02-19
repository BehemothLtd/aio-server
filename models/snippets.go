package models

import (
	"aio-server/exceptions"
	"aio-server/pkg/cryption"
	"fmt"
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

func (s *Snippet) EncryptContent(passKey string) error {
	if s.SnippetType == 1 {
		return nil
	}

	encryptedContent, err := cryption.Encrypt(s.Content, cryption.StringTo32Bytes(passKey))

	if err != nil {
		return exceptions.NewUnprocessableContentError(fmt.Sprintf("Unable to encrypt content %s", err.Error()), nil)
	}

	s.Content = encryptedContent
	return nil
}
