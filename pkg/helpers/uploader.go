package helpers

import (
	"aio-server/pkg/constants"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

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
	bucketName string
	uploadPath string
}

func (u *Uploader) Upload() (*string, error) {
	if _, signedIn := u.Ctx.Get(constants.ContextCurrentUser); !signedIn {
		return nil, errors.New("Unauthorized")
	}

	file, err := u.Ctx.FormFile("file")
	if err != nil {
		return nil, err
	}

	fileName := NewUUID() + file.Filename

	if u.Local {
		return u.uploadLocally(file, fileName)
	}
	return u.uploadToGCS(file, fileName)
}

func (u *Uploader) uploadLocally(file *multipart.FileHeader, fileName string) (*string, error) {
	uploadDst := os.Getenv("UPLOAD_LOCALLY_PATH") + fileName
	err := u.Ctx.SaveUploadedFile(file, uploadDst)
	if err != nil {
		return nil, err
	}
	return &uploadDst, nil
}

func (u *Uploader) uploadToGCS(file *multipart.FileHeader, fileName string) (*string, error) {
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
		uploadPath: u.UploadPath,
	}

	blobFile, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer blobFile.Close()

	err = uploadClient.UploadFile(blobFile, fileName)
	if err != nil {
		return nil, err
	}

	filePublicUrl := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, fileName)
	return &filePublicUrl, nil
}

func (c *ClientUploader) UploadFile(file multipart.File, object string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	wc := c.cl.Bucket(c.bucketName).Object(c.uploadPath + object).NewWriter(ctx)
	defer wc.Close()

	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	return nil
}
