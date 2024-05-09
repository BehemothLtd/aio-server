package globalTypes

import (
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"
	"fmt"
	"os"

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

func (rr *RequestReportType) PendingTime(ctx context.Context) *float64 {
	return helpers.Float64Pointer(rr.RequestReport.PendingTime)
}

func (rr *RequestReportType) RejectedTime(ctx context.Context) *float64 {
	return helpers.Float64Pointer(rr.RequestReport.RejectedTime)
}

func (rr *RequestReportType) UserName(ctx context.Context) *string {
	return &rr.RequestReport.UserName
}

func (rr *RequestReportType) FullName(ctx context.Context) *string {
	return &rr.RequestReport.FullName
}

func (rr *RequestReportType) AvatarKey(ctx context.Context) *string {
	var url string

	key := rr.RequestReport.AvatarKey

	if key == "" {
		return nil
	}

	if os.Getenv("UPLOAD_LOCALLY_PATH") != "" {
		url = os.Getenv("UPLOAD_LOCALLY_PATH") + key
	} else {
		bucketName := os.Getenv("GCS_BUCKET_NAME")
		url = fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, key)
	}

	return &url
}
