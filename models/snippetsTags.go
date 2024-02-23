package models

type SnippetsTag struct {
	Id        int32 `gorm:"not null;autoIncrement"`
	SnippetId int32 `gorm:"not null"`
	TagId     int32 `gorm:"not null"`
}
