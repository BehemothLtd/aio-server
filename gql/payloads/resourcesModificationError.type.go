package payloads

import (
	"aio-server/models"
	"context"
)

type ResourceModifyErrorResolver struct {
	Error *models.ResourceModifyErrors
}

func (r *ResourceModifyErrorResolver) Column(context.Context) *string {
	return &r.Error.Column
}

func (r *ResourceModifyErrorResolver) Errors(context.Context) *[]*string {
	errors := r.Error.Errors

	result := make([]*string, len(errors))

	for i := range errors {
		result[i] = &errors[i]
	}

	return &result
}
