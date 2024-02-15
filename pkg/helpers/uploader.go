package helpers

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Uploader struct {
	Ctx *gin.Context
	Local bool
}

func (u *Uploader) Upload() (*string, error) {
	file, err := u.Ctx.FormFile("file")

	if err != nil {
		return nil, err
	}

	if u.Local {
		uploadDst := os.Getenv("UPLOAD_LOCALLY_PATH") + NewUUID() +file.Filename

		err :=  u.Ctx.SaveUploadedFile(file, uploadDst)

		if err != nil {
			return nil, err
		}

		return &uploadDst, nil
	} else {
		print("UPLOAD TO LOCALLY \n")
		// TODO: upload to GCS
		return nil, nil
	}
}
