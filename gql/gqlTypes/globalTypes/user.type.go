package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

// UserType resolves self information.
type UserType struct {
	Ctx *context.Context
	Db  *gorm.DB

	User *models.User
}

type UserUpdatedType struct {
	User *UserType
}

// ID returns the ID of the user.
func (ut *UserType) ID(context.Context) *graphql.ID {
	return helpers.GqlIDP(ut.User.Id)
}

// Email returns the email of the user.
func (ut *UserType) Email(context.Context) *string {
	return &ut.User.Email
}

// FullName returns the full name of the user.
func (ut *UserType) FullName(context.Context) *string {
	return &ut.User.FullName
}

// Name returns the name of the user.
func (ut *UserType) Name(context.Context) *string {
	return &ut.User.Name
}

// About returns the about of the user.
func (ut *UserType) About(context.Context) *string {
	return ut.User.About
}

// AvatarURL returns the AvatarURL of the user.
func (ut *UserType) AvatarUrl(context.Context) *string {
	if ut.User.Avatar != nil {
		url := ut.User.Avatar.Url()
		return url
	}
	return nil
}

// CreatedAt returns the CreatedAt of the user.
func (ut *UserType) CreatedAt(context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&ut.User.CreatedAt)
}

// CompanyLevelId returns the CompanyLevelId of the user.
func (ut *UserType) CompanyLevelId(context.Context) *graphql.ID {
	if ut.User.CompanyLevelId != nil {
		return helpers.GqlIDP(*ut.User.CompanyLevelId)
	} else {
		return nil
	}
}

// Address returns the Address of the user.
func (ut *UserType) Address(context.Context) *string {
	return ut.User.Address
}

// Phone returns the Phone of the user.
func (ut *UserType) Phone(context.Context) *string {
	return ut.User.Phone
}

// TimingActivedAt returns the TimingActivedAt of the user.
func (ut *UserType) TimingActiveAt(context.Context) *graphql.Time {
	timing := ut.User.Timing

	if timing != nil && timing.ActiveAt != "" {
		return helpers.RubyTimeStringToGqlTime(timing.ActiveAt)
	}

	return nil
}

// timingDeactiveAt returns the timingDeactiveAt of the user.
func (ut *UserType) TimingDeactiveAt(context.Context) *graphql.Time {
	timing := ut.User.Timing

	if timing != nil && timing.InactiveAt != "" {
		return helpers.RubyTimeStringToGqlTime(timing.InactiveAt)
	}
	return nil
}

// Gender returns the Gender of the user.
func (ut *UserType) Gender(context.Context) *string {
	if ut.User.Gender != nil {
		gender := ut.User.Gender.String()
		return &gender
	}

	return nil
}

// Birthday returns the Birthday of the user.
func (ut *UserType) Birthday(context.Context) *graphql.Time {
	return helpers.GqlTimePointer(&ut.User.Birthday)
}

// State returns the State of the user.
func (ut *UserType) State(context.Context) string {
	return ut.User.State.String()
}

// SlackId returns the SlackId of the user.
func (ut *UserType) SlackId(context.Context) *string {
	return ut.User.SlackId
}

// LockVersion returns the lock version of the user.
func (ut *UserType) LockVersion(context.Context) int32 {
	return ut.User.LockVersion
}

func (ut *UserType) IssuesCount(context.Context) *int32 {
	var issuesCount int32

	ut.Db.Table("issue_assignees").Select("Count(distinct issue_id)").Where("user_id = ?", ut.User.Id).Scan(&issuesCount)

	return &issuesCount
}

func (ut *UserType) ProjectsCount(context.Context) *int32 {
	var projectsCount int32

	ut.Db.Table("project_assignees").Select("Count(distinct project_id)").Where("user_id = ?", ut.User.Id).Scan(&projectsCount)

	return &projectsCount
}

func (ut *UserType) ThisMonthWorkingHours(context.Context) *ThisMonthWorkingHoursType {
	result := models.ThisMonthWorkingHours{}

	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	var thisMonthHours, lastMonthHours float64

	ut.Db.Model(&models.WorkingTimelog{}).Select("(IFNULL(SUM(minutes), 0) / 60)").Where("user_id = ?", ut.User.Id).Where(
		"logged_at between ? AND ?", firstOfMonth, lastOfMonth,
	).Scan(&thisMonthHours)

	firstOfLastMonth := time.Date(currentYear, currentMonth-1, 1, 0, 0, 0, 0, currentLocation)
	lastOfLastMonth := firstOfLastMonth.AddDate(0, 1, -1)
	ut.Db.Model(&models.WorkingTimelog{}).Select("(IFNULL(SUM(minutes), 0) / 60)").Where("user_id = ?", ut.User.Id).Where(
		"logged_at between ? AND ?", firstOfLastMonth, lastOfLastMonth,
	).Scan(&lastMonthHours)

	result.Hours = thisMonthHours
	result.UpFromLastMonth = thisMonthHours > lastMonthHours

	var comparation float64
	if lastMonthHours == 0.0 {
		comparation = 1.0
	} else {
		comparation = lastMonthHours
	}

	result.PercentCompareToLastMonth = ((thisMonthHours - lastMonthHours) / comparation) * 100

	timeGraphOnProjects := models.TimeGraphOnProjects{Labels: []string{}, Series: []float64{}}
	ProjectsWorkingHours := []models.ProjectsWorkingHours{}
	ut.Db.Model(&models.WorkingTimelog{}).Select("SUM(minutes) / 60 as hours, projects.name").
		Joins("left join projects on projects.id = working_timelogs.project_id").
		Where("projects.state = 1").
		Where("working_timelogs.logged_at between ? AND ?", firstOfMonth, lastOfMonth).
		Group("project_id").Find(&ProjectsWorkingHours)

	for _, ProjectWorkingHour := range ProjectsWorkingHours {
		timeGraphOnProjects.Labels = append(timeGraphOnProjects.Labels, ProjectWorkingHour.Name)
		timeGraphOnProjects.Series = append(timeGraphOnProjects.Series, ProjectWorkingHour.Hours)
	}

	result.TimeGraphOnProjects = timeGraphOnProjects

	return &ThisMonthWorkingHoursType{
		ThisMonthWorkingHours: &result,
	}
}

type ThisMonthWorkingHoursType struct {
	ThisMonthWorkingHours *models.ThisMonthWorkingHours
}

func (tmwht *ThisMonthWorkingHoursType) Hours(context.Context) *float64 {
	return &tmwht.ThisMonthWorkingHours.Hours
}

func (tmwht *ThisMonthWorkingHoursType) PercentCompareToLastMonth(context.Context) *float64 {
	return &tmwht.ThisMonthWorkingHours.PercentCompareToLastMonth
}

func (tmwht *ThisMonthWorkingHoursType) UpFromLastMonth(context.Context) *bool {
	return &tmwht.ThisMonthWorkingHours.UpFromLastMonth
}

func (tmwht *ThisMonthWorkingHoursType) TimeGraphOnProjects(context.Context) *TimeGraphOnProjectType {
	return &TimeGraphOnProjectType{
		TimeGraphOnProject: &tmwht.ThisMonthWorkingHours.TimeGraphOnProjects,
	}
}

type TimeGraphOnProjectType struct {
	TimeGraphOnProject *models.TimeGraphOnProjects
}

func (tgopt *TimeGraphOnProjectType) Series(context.Context) *[]float64 {
	return &tgopt.TimeGraphOnProject.Series
}

func (tgopt *TimeGraphOnProjectType) Labels(context.Context) *[]string {
	return &tgopt.TimeGraphOnProject.Labels
}
