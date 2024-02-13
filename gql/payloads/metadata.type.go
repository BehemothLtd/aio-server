package payloads

import (
	"aio-server/models"
	"aio-server/pkg/helpers"

	graphql "github.com/graph-gophers/graphql-go"
)

type MetadataResolver struct {
	Metadata *models.Metadata
}

func (mr *MetadataResolver) Total() *graphql.ID {
	return helpers.GqlIDP(mr.Metadata.Total)
}

func (mr *MetadataResolver) PerPage() *int32 {
	r := int32(mr.Metadata.PerPage)
	return &r
}

func (mr *MetadataResolver) Page() *int32 {
	r := int32(mr.Metadata.Page)
	return &r
}

func (mr *MetadataResolver) Pages() *int32 {
	r := int32(mr.Metadata.Pages)
	return &r
}

func (mr *MetadataResolver) Count() *int32 {
	r := int32(mr.Metadata.Count)
	return &r
}

func (mr *MetadataResolver) Next() *int32 {
	r := int32(mr.Metadata.Next)
	return &r
}

func (mr *MetadataResolver) Prev() *int32 {
	r := int32(mr.Metadata.Prev)
	return &r
}

func (mr *MetadataResolver) From() *int32 {
	r := int32(mr.Metadata.From)
	return &r
}

func (mr *MetadataResolver) To() *int32 {
	r := int32(mr.Metadata.To)
	return &r
}
