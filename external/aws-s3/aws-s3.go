package awss3

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/gommon/log"
)

func InitS3(key, secret, region string) *session.Session {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(region),
			Credentials: credentials.NewStaticCredentials(
				key, secret, "",
			),
		},
	)
	if err != nil {
		log.Error("konfigurasi AWS S3 Error:", err)
	}
	return sess
}

func DoUpload(sess *session.Session, region, bucket string, file multipart.FileHeader) (string, error) {
	manager := s3manager.NewUploader(sess)
	src, err := file.Open()
	if err != nil {
		log.Info(err)
		return "", errors.New("error ketika membuka file upload")
	}
	defer src.Close()

	buffer := make([]byte, file.Size)
	src.Read(buffer)
	body, _ := file.Open()

	res, err := manager.Upload(
		&s3manager.UploadInput{
			ACL:         aws.String("public-read"),
			Body:        body,
			Bucket:      aws.String(bucket),
			ContentType: aws.String(http.DetectContentType(buffer)),
			Key:         aws.String(file.Filename),
		},
	)
	if err != nil {
		log.Info("upload output:", res)
		log.Error("error upload:", err)
		return "", errors.New("error upload gambar ke AWS S3")
	}

	url := "https://%s.s3.%s.amazonaws.com/%s"
	link := fmt.Sprintf(url, os.Getenv("S3_BUCKET"), region, file.Filename)
	return link, nil
}

func DoDelete(sess *session.Session, fileName, bucket string) error {
	svc := s3.New(sess)

	deleteInput := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
	}

	res, err := svc.DeleteObject(deleteInput)
	if err != nil {
		log.Info(res)
		log.Error("error delete:", err)
		return errors.New("error delete gambar di AWS S3")
	}
	return nil
}
