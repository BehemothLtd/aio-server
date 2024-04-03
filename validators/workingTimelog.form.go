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
	User                 *models.User
	Issue                *models.Issue
	Repo                 *repository.WorkingTimelogRepository
	WorkingTimelog       *models.WorkingTimelog
	WorkingTimelogUpdate *models.WorkingTimelog
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
		WorkingTimelogUpdate:        &models.WorkingTimelog{},
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
	)
}

func (form *WorkingTimelogForm) validate() error {
	form.assignIssueInfo().validateAndFindRecordWithLoggedAt().validateDescription().validateTimes().summaryErrors()

	if form.Errors != nil {
		return exceptions.NewUnprocessableContentError("", form.Errors)
	}

	return nil
}

func (form *WorkingTimelogForm) validateDescription() *WorkingTimelogForm {
	field := form.FindAttrByCode("description")

	field.ValidateMax(interface{}(int64(constants.MaxLongTextLength)))

	if field.IsClean() {
		form.WorkingTimelogUpdate.Description = *form.Description
	}

	return form
}

func (form *WorkingTimelogForm) validateAndFindRecordWithLoggedAt() *WorkingTimelogForm {
	field := form.FindAttrByCode("loggedAt")

	field.ValidateRequired()
	field.ValidateFormat(constants.DDMMYYY_DateFormat, constants.HUMAN_DD_MM_YY_DateFormat)

	if field.IsClean() {
		form.WorkingTimelog.LoggedAt = *field.Time()
		form.Repo.FindByAttr(form.WorkingTimelog)
	}

	return form
}

func (form *WorkingTimelogForm) validateTimes() *WorkingTimelogForm {
	field := form.FindAttrByCode("minutes")

	field.ValidateRequired()
	field.ValidateMin(int64(15))

	if field.IsClean() {
		if (*form.Minutes % 15) != 0 {
			field.AddError("must a multiple of 15 minutes")
		}

		if form.FindAttrByCode("loggedAt").IsClean() {
			var workingTimelogsByLoggedAt []*models.WorkingTimelog
			form.Repo.GetWorkingTimelogsByLoggedAt(&workingTimelogsByLoggedAt, *form.FindAttrByCode("loggedAt").Time(), form.WorkingTimelog.Id)

			totalLogtime := 0
			for _, wt := range workingTimelogsByLoggedAt {
				totalLogtime = totalLogtime + wt.Minutes
			}

			if totalLogtime+int(*form.Minutes) >= constants.MaximumLogMinutesPerDay {
				field.AddError("logged maximum hours in a day")
			}
		}
	}

	if field.IsClean() {
		form.WorkingTimelogUpdate.Minutes = int(*form.Minutes)
	}
	return form
}

func (form *WorkingTimelogForm) assignIssueInfo() *WorkingTimelogForm {
	form.WorkingTimelog.ProjectId = form.Issue.ProjectId
	form.WorkingTimelog.IssueId = form.Issue.Id
	form.WorkingTimelog.UserId = form.User.Id

	return form
}

func (form *WorkingTimelogForm) Save() error {
	if err := form.validate(); err != nil {
		return err
	}

	if form.WorkingTimelog.Id != 0 {
		err := form.Repo.Update(form.WorkingTimelog, *form.WorkingTimelogUpdate)
		if err != nil {
			return err
		}
	} else {
		form.WorkingTimelog.Description = form.WorkingTimelogUpdate.Description
		form.WorkingTimelog.Minutes = form.WorkingTimelogUpdate.Minutes
		err := form.Repo.Create(form.WorkingTimelog)
		if err != nil {
			return err
		}
	}

	return nil
}
