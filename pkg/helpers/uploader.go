package helpers

import (
	"aio-server/database"
	"aio-server/models"
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

type UploadedBlob struct {
	Url string `json:"url"`
	Key string `json:"key"`
}

func (u *Uploader) Upload() ([]*UploadedBlob, error) {
	if _, signedIn := u.Ctx.Get(constants.ContextCurrentUser); !signedIn {
		return nil, errors.New("Unauthorized")
	}

	err := u.Ctx.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		return nil, err
	}

	var uploadedBlobs []*UploadedBlob

	form := u.Ctx.Request.MultipartForm
	files := form.File["files[]"]

	for _, file := range files {
		filename := file.Filename

		var uploadedBlob *UploadedBlob
		var uploadErr error

		if u.Local {
			uploadedBlob, uploadErr = u.uploadLocally(file, filename)
		} else {
			uploadedBlob, uploadErr = u.uploadToGCS(file, filename)
		}

		if uploadErr != nil {
			return nil, uploadErr
		}

		uploadedBlobs = append(uploadedBlobs, uploadedBlob)
	}

	return uploadedBlobs, nil
}

func (u *Uploader) uploadLocally(file *multipart.FileHeader, fileName string) (*UploadedBlob, error) {
	blobKey := NewUUID()

	uploadDst := os.Getenv("UPLOAD_LOCALLY_PATH") + blobKey
	err := u.Ctx.SaveUploadedFile(file, uploadDst)
	if err != nil {
		return nil, err
	}

	newBlob := models.AttachmentBlob{
		Key:      blobKey,
		Filename: fileName,
	}

	if err := database.Db.Create(&newBlob); err != nil {
		return &UploadedBlob{
			Url: uploadDst,
			Key: newBlob.Key,
		}, nil
	} else {
		return nil, err.Error
	}
}

func (u *Uploader) uploadToGCS(file *multipart.FileHeader, fileName string) (*UploadedBlob, error) {
	bucketName := os.Getenv("GCS_BUCKET_NAME")
	projectId := os.Getenv("GCS_PROJECT_ID")
	gcsAccountService := os.Getenv("GCS_ACCOUNT_SERVICE")

	if bucketName == "" || projectId == "" || gcsAccountService == "" {
		return nil, errors.New("invalid Setting for Upload")
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
	blobKey := NewUUID()

	err = uploadClient.UploadFile(blobFile, blobKey)
	if err != nil {
		return nil, err
	}

	newBlob := models.AttachmentBlob{
		Key:      blobKey,
		Filename: fileName,
	}

	if err := database.Db.Create(&newBlob); err != nil {
		uploadClient.DeleteFile(blobKey)

		return nil, err.Error
	} else {
		filePublicUrl := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, newBlob.Key)

		return &UploadedBlob{
			Url: filePublicUrl,
			Key: newBlob.Key,
		}, nil
	}

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

func (c *ClientUploader) DeleteFile(key string) {
	ctx := context.Background()
	o := c.cl.Bucket(c.bucketName).Object(key)
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	o.Delete(ctx)
}
