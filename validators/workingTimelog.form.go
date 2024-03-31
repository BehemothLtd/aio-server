package validators

import (
	"aio-server/exceptions"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/constants"
	"aio-server/pkg/helpers"
	"aio-server/repository"
)

type WorkingTimelogForm struct {
	Form
	insightInputs.SelfWorkingTimelogFormInput
	User           *models.User
	Issue          *models.Issue
	Repo           *repository.WorkingTimelogRepository
	WorkingTimelog *models.WorkingTimelog
}

func NewWorkingTimelogFormValidator(
	input *insightInputs.SelfWorkingTimelogFormInput,
	user *models.User,
	repo *repository.WorkingTimelogRepository,
	workingTimelog *models.WorkingTimelog,
	issue *models.Issue,
) WorkingTimelogForm {
	form := WorkingTimelogForm{
		Form:                        Form{},
		SelfWorkingTimelogFormInput: *input,
		User:                        user,
		Repo:                        repo,
		WorkingTimelog:              workingTimelog,
		Issue:                       issue,
	}

	form.assignAttributes()

	return form
}

func (form *WorkingTimelogForm) assignAttributes() {
	form.AddAttributes(
		&StringAttribute{
			FieldAttribute: FieldAttribute{
				Code: "description",
			},
			Value: helpers.GetStringOrDefault(form.Description),
		},
		&TimeAttribute{
			FieldAttribute: FieldAttribute{
				Code: "loggedAt",
			},
			Value: helpers.GetStringOrDefault(form.LoggedAt),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "minutes",
			},
			Value: helpers.GetInt32OrDefault(form.Minutes),
		},
		&IntAttribute[int32]{
			FieldAttribute: FieldAttribute{
				Code: "lockVersion",
			},
			Value: helpers.GetInt32OrDefault(form.LockVersion),
		},
	)
}

func (form *WorkingTimelogForm) validate() error {
	form.validateDescription().validateLoggedAt().validateTimes().assignIssueInfo().validateLockVersion().summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}

	return nil
}

func (form *WorkingTimelogForm) validateDescription() *WorkingTimelogForm {
	field := form.FindAttrByCode("description")

	field.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

	if field.IsClean() {
		form.WorkingTimelog.Description = form.Description
	}

	return form
}

func (form *WorkingTimelogForm) validateLoggedAt() *WorkingTimelogForm {
	field := form.FindAttrByCode("loggedAt")

	field.ValidateRequired()
	field.ValidateFormat("2006-01-02", "%y-%m-%d")

	if field.IsClean() {
		form.WorkingTimelog.LoggedAt = *field.Time()
	}

	return form
}

func (form *WorkingTimelogForm) validateTimes() *WorkingTimelogForm {
	field := form.FindAttrByCode("minutes")

	field.ValidateRequired()
	field.ValidateMin(int64(15))

	if field.IsClean() {
		if *form.Minutes%15 != 0 {
			field.AddError("must a multiple of 15 minutes")
		}

		if form.FindAttrByCode("loggedAt").IsClean() {
			var workingTimelogsByLoggedAt []*models.WorkingTimelog
			form.Repo.GetWorkingTimelogsByLoggedAt(&workingTimelogsByLoggedAt, *form.FindAttrByCode("loggedAt").Time())

			totalLogtime := 0
			for _, wt := range workingTimelogsByLoggedAt {
				if wt.Id == form.WorkingTimelog.Id {
					continue
				}
				totalLogtime = totalLogtime + wt.Minutes
			}

			if totalLogtime+int(*form.Minutes) >= constants.MaximumLogMinutesPerDay {
				field.AddError("logged maximum hours in a day")
			}
		}
	}

	if field.IsClean() {
		form.WorkingTimelog.Minutes = int(*form.Minutes)
	}
	return form
}

func (form *WorkingTimelogForm) validateLockVersion() *WorkingTimelogForm {
	newRecord := form.WorkingTimelog.Id == 0
	currentLockVersion := form.WorkingTimelog.LockVersion
	newLockVersion := helpers.GetInt32OrDefault(form.LockVersion) + 1
	formAttribute := form.FindAttrByCode("lockVersion")

	formAttribute.ValidateMin(interface{}(int64(currentLockVersion)))

	if newRecord {
		return form
	}

	if currentLockVersion >= newLockVersion {
		formAttribute.AddError("Attempted to update stale object")
	}

	return form
}

func (form *WorkingTimelogForm) assignIssueInfo() *WorkingTimelogForm {
	form.WorkingTimelog.ProjectId = form.Issue.ProjectId
	form.WorkingTimelog.IssueId = form.Issue.Id
	form.WorkingTimelog.UserId = form.User.Id

	return form
}

func (form *WorkingTimelogForm) Create() error {
	if err := form.validate(); err != nil {
		return err
	}

	if err := form.Repo.Create(form.WorkingTimelog); err != nil {
		return err
	}

	return nil
}

func (form *WorkingTimelogForm) Update() error {
	if err := form.validate(); err != nil {
		return err
	}

	if err := form.Repo.Update(form.WorkingTimelog); err != nil {
		return err
	}

	return nil
}
