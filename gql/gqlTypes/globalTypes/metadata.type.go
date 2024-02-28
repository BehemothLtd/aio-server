package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"

	graphql "github.com/graph-gophers/graphql-go"
)

// MetadataType resolves the fields of the Metadata type.
type MetadataType struct {
	Metadata *models.Metadata
}

// Total returns the total value of Metadata as a graphql.ID pointer.
func (mt *MetadataType) Total() *graphql.ID {
	return helpers.IDPointer(helpers.GqlIDValue(mt.Metadata.Total))
}

// PerPage returns the per page value of Metadata as an int32 pointer.
func (mt *MetadataType) PerPage() *int32 {
	return helpers.Int32Pointer(int32(mt.Metadata.PerPage))
}

// Page returns the page value of Metadata as an int32 pointer.
func (mt *MetadataType) Page() *int32 {
	return helpers.Int32Pointer(int32(mt.Metadata.Page))
}

// Pages returns the pages value of Metadata as an int32 pointer.
func (mt *MetadataType) Pages() *int32 {
	return helpers.Int32Pointer(int32(mt.Metadata.Pages))
}

// Count returns the count value of Metadata as an int32 pointer.
func (mt *MetadataType) Count() *int32 {
	return helpers.Int32Pointer(int32(mt.Metadata.Count))
}

// Next returns the next value of Metadata as an int32 pointer.
func (mt *MetadataType) Next() *int32 {
	return helpers.Int32Pointer(int32(mt.Metadata.Next))
}

// Prev returns the prev value of Metadata as an int32 pointer.
func (mt *MetadataType) Prev() *int32 {
	return helpers.Int32Pointer(int32(mt.Metadata.Prev))
}

// From returns the from value of Metadata as an int32 pointer.
func (mt *MetadataType) From() *int32 {
	return helpers.Int32Pointer(int32(mt.Metadata.From))
}

// To returns the to value of Metadata as an int32 pointer.
func (mt *MetadataType) To() *int32 {
	return helpers.Int32Pointer(int32(mt.Metadata.To))
}
