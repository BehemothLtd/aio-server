package constants

const (
	AuthorizationHeader = "BhmAIO-Authorization"
	GinContextKey       = "AIOContextKey"
	ContextCurrentUser  = "CurrentUser"

	MaxStringLength   = 255
	MaxLongTextLength = 4294967295

	BadRequestErrorCode = 400
	BadRequestErrorMsg  = "Bad Request"

	NotFoundErrorCode = 404
	NotFoundErrorMsg  = "Not Found"

	UnauthorizedErrorCode = 401
	UnauthorizedErrorMsg  = "You need to sign in to perform this action"

	UnprocessableContentErrorCode = 422
	UnprocessableContentErrorMsg  = "Please check your input"

	DDMMYYYY_DateFormat       = "02-01-2006" // "Date-Month-Year"
	HUMAN_DDMMYYYY_DateFormat = "%d-%m-%y"

	DDMMYYYY_DateSplashFormat       = "02/01/2006"
	YYYYMMDD_DateSplashFormat       = "2006/01/02"
	HUMAN_DDMMYYYY_DateSplashFormat = "%d/%m/%y"

	YYYYMMDD_DateFormat       = "2006-01-02" // "Month-Date-Year"
	HUMAN_YYYYMMDD_DateFormat = "%y-%m-%d"

	DateTimeFormat           = "2006-01-02 15:04:05.000"
	DateTimeZoneFormat       = "2006-01-02 15:04:05 -0700"
	HUMAN_DateTimeZoneFormat = "%d-%m-%y %H:%M -%Z"

	DDMMYYY_HHMM_DateFormat       = "2-1-2006 15:04"
	HUMAN_DDMMYYY_HHMM_DateFormat = "%d-%m-%y %H:%M"
	MMDD_DateFormatForChart       = "Jan 02"

	RequestTimeOut = 20
	Get            = "GET"
	Post           = "POST"

	MaximumLogMinutesPerDay = 840

	BODGroup = "BOD"

	// Slack callback id
	SlackChangeStateRq = "change_state_rq"

	Charset     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	EmailFormat = `\A([^@\s]+)@((?:[-a-z0-9]+\.)+[a-z]{2,})\z`
)

func RequiredIssueStatusIdsForKanbanProject() []int32 {
	return []int32{2, 3, 7}
}

func RequiredIssueStatusIdsForScrumProject() []int32 {
	return []int32{1, 2, 3, 7}
}

func ScrumDefaultIssueStatus() []string {
	return []string{"Backlog", "To Do", "Doing", "Done"}
}
