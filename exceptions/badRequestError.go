package exceptions

import (
	"aio-server/pkg/constants"
	"fmt"
)

type BadRequestError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e BadRequestError) Error() string {
	return fmt.Sprintf("error [%d]: %s", e.Code, e.Message)
}

func (e BadRequestError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}

func NewBadRequestError(message string) BadRequestError {
	var returnMessage string

	if message == "" {
		returnMessage = constants.BadRequestErrorMsg
	} else {
		returnMessage = message
	}

	return BadRequestError{
		Code:    constants.BadRequestErrorCode,
		Message: returnMessage,
	}
}
