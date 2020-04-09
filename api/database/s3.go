package database

import (
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

//CreateObject create s3 bucket
func CreateObject(src multipart.File) (string, error) {
	db, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)
	if err != nil {
		return "", err
	}
	bucket := "cotton-portfolio"
	filename := "originalname"
	// file, err := os.Open(filename)

	uploader := s3manager.NewUploader(db)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   src,
	})
	return result.Location, err
}
