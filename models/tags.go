package models

type Tag struct {
	Id     int32  `gorm:"not null;autoIncrement"`
	Name   string `gorm:"not null;"`
	UserId int32  `gorm:"not null;"`
}
