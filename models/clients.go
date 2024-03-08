package models

import "time"

type Client struct {
	Id             int32
	Name           string
	ShowOnHomePage int
	LockVersion    int32
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
