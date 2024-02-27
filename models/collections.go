package models

import (
	"time"
)

type Collection struct {
	Id          int32  `gorm:"not null;autoIncrement"`
	Title       string `gorm:"not null;type:varchar(255);default:null"`
	UserId      int32  `gorm:"not null;type:bigint;default:null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	LockVersion int32 `gorm:"not null;autoIncrement;default:0"`
}

type CollectionsCollection struct {
	Collection []*Collection
	Metadata   *Metadata
}

type CollectionQuery struct {
	TitleCont string
}
