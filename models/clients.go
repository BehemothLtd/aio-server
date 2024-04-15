package models

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	Id             int32
	Name           string
	ShowOnHomePage bool
	LockVersion    int32 `gorm:"not null;default:0"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (r *Client) BeforeUpdate(tx *gorm.DB)(err error){
	if tx.Statement.Changed(){
		tx.Statement.SetColumn("lock_version", r.LockVersion+1)
	}

	return
}

func (r *Client) AfterDelete(tx *gorm.DB)(err error){
	return tx.Model(&Project{}).Where(`client_id = ?`,r.Id).Update("client_id",nil).Error
}
