package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"

	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type RequestReportType struct {
	Ctx *context.Context
	Db  *gorm.DB

	RequestReport *models.RequestReport
}

func (rr *RequestReportType) UserId(ctx context.Context) *graphql.ID {
	return helpers.GqlIDP(rr.RequestReport.UserId)
}

func (rr *RequestReportType) User(ctx context.Context) *UserType {
	return &UserType{User: &rr.RequestReport.User}
}

func (rr *RequestReportType) ApprovedTime(ctx context.Context) *float64 {
	return helpers.Float64Pointer(rr.RequestReport.ApprovedTime)
}

func (rr *RequestReportType) PeddingTime(ctx context.Context) *float64 {
	return helpers.Float64Pointer(rr.RequestReport.PeddingTime)
}

func (rr *RequestReportType) RejectedTime(ctx context.Context) *float64 {
	return helpers.Float64Pointer(rr.RequestReport.RejectedTime)
}
