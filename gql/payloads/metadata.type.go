package payloads

import (
	"aio-server/models"
	"aio-server/pkg/helpers"

	graphql "github.com/graph-gophers/graphql-go"
)

type MetadataResolver struct {
	M *models.Metadata
}

func (m *MetadataResolver) Total() *graphql.ID {
	return helpers.GqlIDP(m.M.Total)
}

func (m *MetadataResolver) PerPage() *int32 {
	r := int32(m.M.PerPage)
	return &r
}

func (m *MetadataResolver) Page() *int32 {
	r := int32(m.M.Page)
	return &r
}

func (m *MetadataResolver) Pages() *int32 {
	r := int32(m.M.Pages)
	return &r
}

func (m *MetadataResolver) Count() *int32 {
	r := int32(m.M.Count)
	return &r
}

func (m *MetadataResolver) Next() *int32 {
	r := int32(m.M.Next)
	return &r
}

func (m *MetadataResolver) Prev() *int32 {
	r := int32(m.M.Prev)
	return &r
}

func (m *MetadataResolver) From() *int32 {
	r := int32(m.M.From)
	return &r
}

func (m *MetadataResolver) To() *int32 {
	r := int32(m.M.To)
	return &r
}
