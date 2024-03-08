package models

import (
	"time"
)

type Collection struct {
	Id        int32      `gorm:"not null;autoIncrement"`
	Title     string     `gorm:"not null;type:varchar(255);default:null"`
	UserId    int32      `gorm:"not null;type:bigint;default:null"`
	Snippets  []*Snippet `gorm:"many2many:snippets_collections"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CollectionsCollection struct {
	Collection []*Collection
	Metadata   *Metadata
}
