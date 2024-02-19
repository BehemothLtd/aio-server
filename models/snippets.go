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
	LockVersion    int32 `gorm:"not null;autoIncrement;default:0"`
}

type SnippetsQuery struct {
	TitleCont string
}

type SnippetsCollection struct {
	Collection []*Snippet
	Metadata   *Metadata
}

func (s *Snippet) EncryptContent(Passkey string) error {
	if s.SnippetType == 1 {
		return nil
	}

	encryptedContent, err := cryption.Encrypt(s.Content, cryption.StringTo32Bytes(Passkey))

	if err != nil {
		return exceptions.NewUnprocessableContentError(fmt.Sprintf("Unable to encrypt content %s", err.Error()), nil)
	}

	s.Content = encryptedContent
	return nil
}

func (s *Snippet) DecryptContent(Passkey string) (*string, error) {
	if s.SnippetType == 1 {
		return nil, nil
	}

	decryptedContent, err := cryption.Decrypt(s.Content, cryption.StringTo32Bytes(Passkey))

	if err != nil {
		return nil, exceptions.NewUnprocessableContentError("Incorrect Passkey", nil)
	}

	return &decryptedContent, nil
}
