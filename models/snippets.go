package models

import (
	"aio-server/enums"
	"aio-server/exceptions"
	"aio-server/pkg/cryption"
	"fmt"
	"time"
)

type Snippet struct {
	Id                int32  `gorm:"not null;"`
	Title             string `gorm:"not null;type:varchar(255);default:null"`
	Content           string `gorm:"not null;type:longtext;default:null"`
	UserId            int32  `gorm:"not null;type:bigint;default:null"`
	Slug              string
	SnippetType       enums.SnippetType `gorm:"not null;"`
	FavoritesCount    int               `gorm:"not null;default:0"`
	FavoritedUsers    []User            `gorm:"many2many:snippets_favorites"`
	Pins              []Pin             `gorm:"polymorphic:Parent;polymorphicValue:1"`
	Tags              []*Tag            `gorm:"many2many:snippets_tags;"`
	SnippetsFavorites []*SnippetsFavorite
	CreatedAt         time.Time
	UpdatedAt         time.Time
	LockVersion       int32 `gorm:"not null;default:0"`
}

func (s *Snippet) EncryptContent(Passkey string) error {
	if s.SnippetType == enums.SnippetTypePublic {
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
	if s.SnippetType == enums.SnippetTypePublic {
		return nil, nil
	}

	decryptedContent, err := cryption.Decrypt(s.Content, cryption.StringTo32Bytes(Passkey))

	if err != nil {
		return nil, exceptions.NewUnprocessableContentError("Incorrect Passkey", nil)
	}

	return &decryptedContent, nil
}
