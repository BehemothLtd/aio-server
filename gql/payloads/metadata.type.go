package payloads

import (
	"aio-server/models"
	"aio-server/pkg/helpers"

	graphql "github.com/graph-gophers/graphql-go"
)

type MetadataResolver struct {
	Metadata *models.Metadata
}

func (m *MetadataResolver) Total() *graphql.ID {
	return helpers.GqlIDP(m.Metadata.Total)
}

func (m *MetadataResolver) PerPage() *int32 {
	r := int32(m.Metadata.PerPage)
	return &r
}

func (m *MetadataResolver) Page() *int32 {
	r := int32(m.Metadata.Page)
	return &r
}

func (m *MetadataResolver) Pages() *int32 {
	r := int32(m.Metadata.Pages)
	return &r
}

func (m *MetadataResolver) Count() *int32 {
	r := int32(m.Metadata.Count)
	return &r
}

func (m *MetadataResolver) Next() *int32 {
	r := int32(m.Metadata.Next)
	return &r
}

func (m *MetadataResolver) Prev() *int32 {
	r := int32(m.Metadata.Prev)
	return &r
}

func (m *MetadataResolver) From() *int32 {
	r := int32(m.Metadata.From)
	return &r
}

func (m *MetadataResolver) To() *int32 {
	r := int32(m.Metadata.To)
	return &r
}
