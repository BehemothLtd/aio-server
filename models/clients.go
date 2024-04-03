package models

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	Id             int32
	Name           string
	ShowOnHomePage bool
	LockVersion    int32 `gorm:"not nul;autoIncrement;default:0"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (r *Client) BeforeUpdate(tx *gorm.DB)(err error){
	if tx.Statement.Changed(){
		tx.Statement.SetColumn("lock_version", r.LockVersion+1)
	}

	return
}
