package models

import "time"

type TimesheetTemplate struct {
	Id          int32
	Name        string
	UserId      int32
	User        User
	Settings    string // will need to be changed to a typed struct later
	LockVersion int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// HERE IS THE DEMO CODE FOR RELATION RECORDS LOAD. PLEASE DELETE THESE CODE AFTER
// YOU FINISH IMPLEMENT

// timesheetTemplate load
// timesheetTemplate := models.TimessheetTemplate{
// 	Id: 1,
// }

// db.Model(&timesheetTemplate).Preload("User").First(&timesheetTemplate)

// fmt.Printf("RECORD %+v", timesheetTemplate)
