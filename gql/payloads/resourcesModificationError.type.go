package payloads

import (
	"aio-server/exceptions"
	"context"
)

// ResourceModifyErrorResolver resolves errors encountered during resource modification.
type ResourceModifyErrorResolver struct {
	Error *exceptions.ResourceModifyErrors
}

// Column returns the column associated with the error.
func (rmer *ResourceModifyErrorResolver) Column(ctx context.Context) *string {
	return &rmer.Error.Field
}

// Errors returns a list of errors encountered during resource modification.
func (rmer *ResourceModifyErrorResolver) Errors(ctx context.Context) *[]*string {
	errors := rmer.Error.Errors

	result := make([]*string, len(errors))

	for i, err := range errors {
		result[i] = &err
	}

	return &result
}
