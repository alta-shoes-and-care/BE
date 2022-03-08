package uploader

import (
	"errors"
	awss3 "final-project/services/aws-s3"
	"mime/multipart"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/labstack/gommon/log"
)

func Uploader(awsSess *session.Session, region string, file *multipart.FileHeader) (string, error) {
	file.Filename = strings.ReplaceAll(file.Filename, " ", "_")

	link, err := awss3.DoUpload(awsSess, region, *file)
	if err != nil {
		log.Info(err)
		return "", errors.New("gagal mengunggah gambar")
	}
	return link, nil
}
