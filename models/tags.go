package models

import "time"

type Tag struct {
	Id          int32
	Name        string
	UserId      int32
	Snippets    []*Snippet `gorm:"many2many:snippets_tags;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	LockVersion int32 `gorm:"default:1"`
}
