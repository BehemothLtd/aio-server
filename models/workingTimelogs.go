package models

import "time"

type WorkingTimelog struct {
	Id          int32
	UserId      int32
	User        User
	ProjectId   int32
	Project     Project
	IssueId     int32
	Issue       Issue
	Minutes     int
	Description string
	LoggedAt    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// HERE IS THE DEMO CODE FOR RELATION RECORDS LOAD. PLEASE DELETE THESE CODE AFTER
// YOU FINISH IMPLEMENT

// workingTimelog load
// workingTimelog := models.WorkingTimelog{
// 	Id: 1,
// }

// db.Model(&workingTimelog).Preload("User").Preload("Project").Preload("Issue").First(&workingTimelog)

// fmt.Printf("WORKING TIME LOG %+v", workingTimelog)

// User Loaded with workingTimelog

// user := models.User{
// 	Id: 1,
// }

// db.Model(&user).Preload("WorkingTimelogs").First(&user)

// fmt.Printf("RECORD %+v", user)
