package payloads

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type MetadataResolver struct {
	Db  *gorm.DB
	Ctx *context.Context
}

func (m *MetadataResolver) Total() *graphql.ID {
	return nil
}

func (m *MetadataResolver) PerPage() *int32 {
	return nil
}

func (m *MetadataResolver) Page() *int32 {
	return nil
}

func (m *MetadataResolver) Pages() *int32 {
	return nil
}

func (m *MetadataResolver) Count() *int32 {
	return nil
}

func (m *MetadataResolver) Next() *int32 {
	return nil
}

func (m *MetadataResolver) Prev() *int32 {
	return nil
}

func (m *MetadataResolver) From() *int32 {
	return nil
}

func (m *MetadataResolver) To() *int32 {
	return nil
}
