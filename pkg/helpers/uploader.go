package helpers

import (
	"aio-server/pkg/constants"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"mime/multipart"

	"cloud.google.com/go/storage"

	"github.com/gin-gonic/gin"
)

type Uploader struct {
	Ctx        *gin.Context
	Local      bool
	UploadPath string
}

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

func (u *Uploader) Upload() (*string, error) {
	_, signedIn := u.Ctx.Get(constants.ContextCurrentUser)

	if !signedIn {
		return nil, errors.New("Unauthorized")
	}

	file, err := u.Ctx.FormFile("file")

	if err != nil {
		return nil, err
	}

	fileName := NewUUID() + file.Filename

	if u.Local {
		uploadDst := os.Getenv("UPLOAD_LOCALLY_PATH") + fileName

		err := u.Ctx.SaveUploadedFile(file, uploadDst)

		if err != nil {
			return nil, err
		}

		return &uploadDst, nil
	} else {
		bucketName := os.Getenv("GCS_BUCKET_NAME")
		projectId := os.Getenv("GCS_PROJECT_ID")
		gcsAccountService := os.Getenv("GCS_ACCOUNT_SERVICE")

		if bucketName == "" || projectId == "" || gcsAccountService == "" {
			return nil, errors.New("Invalid Setting for Upload")
		}

		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", gcsAccountService)
		client, err := storage.NewClient(context.Background())
		if err != nil {
			return nil, err
		}

		uploadClient := &ClientUploader{
			cl:         client,
			bucketName: bucketName,
			projectID:  projectId,
			uploadPath: u.UploadPath,
		}

		blobFile, err := file.Open()
		if err != nil {
			return nil, err
		}

		err = uploadClient.UploadFile(blobFile, fileName)
		if err != nil {

			return nil, err
		}

		filePublicUrl := "https://storage.googleapis.com/" + bucketName + "/" + fileName

		return &filePublicUrl, nil
	}
}

func (c *ClientUploader) UploadFile(file multipart.File, object string) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := c.cl.Bucket(c.bucketName).Object(c.uploadPath + object).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}
