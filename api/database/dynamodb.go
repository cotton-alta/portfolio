package database

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

//Work dynamodb model
type Work struct {
	WorkID    string
	Timestamp time.Time
	Title     string
	Detail    string
	Image     string
}

//CreateDynamo create item
func CreateDynamo(title string, content string, href string) error {
	db := dynamo.New(session.New(), &aws.Config{
		Region: aws.String("ap-northeast-1"),
	})
	table := db.Table("portfolio-work")

	work := Work{WorkID: "001",
		Timestamp: time.Now().UTC(),
		Title:     title,
		Detail:    content,
		Image:     href}

	err := table.Put(work).Run()
	fmt.Println("OK")
	return err
}

//GetDynamo get items list
func GetDynamo() ([]Work, error) {
	db := dynamo.New(session.New(), &aws.Config{
		Region: aws.String("ap-northeast-1"),
	})
	table := db.Table("portfolio-work")

	var results []Work

	err := table.Get("WorkID", "001").All(&results)
	return results, err
}
