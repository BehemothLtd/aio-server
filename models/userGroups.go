package models

import (
	"time"
)

type UserGroup struct {
	Id        int32  `gorm:"not null;autoIncrement"`
	Title     string `gorm:"not null;type:varchar(255);default:null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserGroupsQuery struct {
	TitleCont string
}

type UserGroupCollection struct {
	Collection []*UserGroup
	Metadata   *Metadata
}
