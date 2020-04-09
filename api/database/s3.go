package database

import (
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

//CreateObject create s3 bucket
func CreateObject(src multipart.File, originalname string, contentType string) (string, error) {
	db, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)
	if err != nil {
		return "", err
	}
	bucket := "cotton-portfolio"
	filename := originalname
	fileType := contentType

	uploader := s3manager.NewUploader(db)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(filename),
		ACL:         aws.String("public-read"),
		ContentType: aws.String(fileType),
		Body:        src,
	})
	return result.Location, err
}
