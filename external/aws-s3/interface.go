package awss3

import (
	"mime/multipart"
)

type MyS3Client interface {
	DoUpload(region, bucket string, file *multipart.FileHeader) (string, error)
	DoDelete(fileName, bucket string) error
}
