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

	DDMMYYY_DateFormat            = "2-1-2006" // "Month-Date-Year"
	YYMMDD_DateFormat             = "2006-01-02"
	HUMAN_DD_MM_YY_DateFormat     = "%d-%m-%y"
	DDMMYYY_HHMM_DateFormat       = "2-1-2006 15:04"
	HUMAN_DDMMYYY_HHMM_DateFormat = "%d-%m-%y %H:%M"

	RequestTimeOut = 20
	Get            = "GET"
	Post           = "POST"
)

func RequiredIssueStatusIdsForKanbanProject() []int32 {
	return []int32{2, 3, 7}
}

func RequiredIssueStatusIdsForScrumProject() []int32 {
	return []int32{1, 2, 3, 7}
}
