package models

import "time"

type ProjectExperience struct {
	Id          int32
	Title       string
	Description string
	UserId      int32
	ProjectId   int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
