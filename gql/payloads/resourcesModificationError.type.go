package payloads

import (
	"aio-server/models"
	"context"
)

type ResourceModifyErrorResolver struct {
	Error *models.ResourceModifyErrors
}

func (rmer *ResourceModifyErrorResolver) Column(context.Context) *string {
	return &rmer.Error.Column
}

func (rmer *ResourceModifyErrorResolver) Errors(context.Context) *[]*string {
	errors := rmer.Error.Errors

	result := make([]*string, len(errors))

	for i := range errors {
		result[i] = &errors[i]
	}

	return &result
}
