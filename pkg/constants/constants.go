package constants

const (
	AuthorizationHeader = "BhmAIO-Authorization"
	GinContextKey       = "AIOContextKey"
	ContextCurrentUser  = "CurrentUser"

	MaxStringLength         = 255
	MaxLongTextLength       = 4294967295

	BadRequestErrorCode = 400
	BadRequestErrorMsg  = "Bad Request"

	NotFoundErrorCode = 404
	NotFoundErrorMsg  = "Not Found"

	UnauthorizedErrorCode = 401
	UnauthorizedErrorMsg  = "You need to sign in to perform this action"

	UnprocessableContentErrorCode = 422
	UnprocessableContentErrorMsg  = "Please check your input"
)
