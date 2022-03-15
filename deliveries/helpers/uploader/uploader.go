package uploader

import (
	awss3 "final-project/external/aws-s3"
	"mime/multipart"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/labstack/gommon/log"
)

func Uploader(awsSess *session.Session, region, bucket string, file multipart.FileHeader) (string, error) {
	file.Filename = strings.ReplaceAll(file.Filename, " ", "_")

	link, err := awss3.DoUpload(awsSess, region, bucket, file)
	if err != nil {
		log.Info(err)
		return "", err
	}
	return link, nil
}
