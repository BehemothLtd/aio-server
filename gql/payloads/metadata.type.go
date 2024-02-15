package payloads

import (
	"aio-server/models"
	"aio-server/pkg/helpers"

	graphql "github.com/graph-gophers/graphql-go"
)

// MetadataResolver resolves the fields of the Metadata type.
type MetadataResolver struct {
	Metadata *models.Metadata
}

// Total returns the total value of Metadata as a graphql.ID pointer.
func (mr *MetadataResolver) Total() *graphql.ID {
	return helpers.IDPointer(*helpers.GqlIDP(mr.Metadata.Total))
}

// PerPage returns the per page value of Metadata as an int32 pointer.
func (mr *MetadataResolver) PerPage() *int32 {
	return helpers.Int32Pointer(int32(mr.Metadata.PerPage))
}

// Page returns the page value of Metadata as an int32 pointer.
func (mr *MetadataResolver) Page() *int32 {
	return helpers.Int32Pointer(int32(mr.Metadata.Page))
}

// Pages returns the pages value of Metadata as an int32 pointer.
func (mr *MetadataResolver) Pages() *int32 {
	return helpers.Int32Pointer(int32(mr.Metadata.Pages))
}

// Count returns the count value of Metadata as an int32 pointer.
func (mr *MetadataResolver) Count() *int32 {
	return helpers.Int32Pointer(int32(mr.Metadata.Count))
}

// Next returns the next value of Metadata as an int32 pointer.
func (mr *MetadataResolver) Next() *int32 {
	return helpers.Int32Pointer(int32(mr.Metadata.Next))
}

// Prev returns the prev value of Metadata as an int32 pointer.
func (mr *MetadataResolver) Prev() *int32 {
	return helpers.Int32Pointer(int32(mr.Metadata.Prev))
}

// From returns the from value of Metadata as an int32 pointer.
func (mr *MetadataResolver) From() *int32 {
	return helpers.Int32Pointer(int32(mr.Metadata.From))
}

// To returns the to value of Metadata as an int32 pointer.
func (mr *MetadataResolver) To() *int32 {
	return helpers.Int32Pointer(int32(mr.Metadata.To))
}
