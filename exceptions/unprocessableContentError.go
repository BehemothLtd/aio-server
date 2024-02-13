package exceptions

import (
	"aio-server/pkg/constants"
	"fmt"
)

type ResourceModifyErrors struct {
	Field  string   `json:"field"`
	Errors []string `json:"errors"`
}

type UnprocessableContentError struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Errors  []ResourceModifyErrors `json:"errors"`
}

func (e UnprocessableContentError) Error() string {
	return fmt.Sprintf("error [%d]: %s", e.Code, e.Message)
}

func (e UnprocessableContentError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
		"errors":  e.Errors,
	}
}

func NewUnprocessableContentError(message *string, errors *[]ResourceModifyErrors) *UnprocessableContentError {
	var returnMessage string

	if message == nil {
		returnMessage = constants.UnprocessableContentErrorMsg
	} else {
		returnMessage = *message
	}

	return &UnprocessableContentError{
		Code:    constants.UnprocessableContentErrorCode,
		Message: returnMessage,
		Errors:  *errors,
	}
}

func (e *UnprocessableContentError) AddError(rme ResourceModifyErrors) {
	e.Errors = append(e.Errors, rme)
}
