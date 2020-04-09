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

//Article dynamodb model
type Article struct {
	ArticleID string
	Timestamp time.Time
	Title     string
	Detail    string
	Image     string
}

//CreateDynamo create item
func CreateDynamo(title string, content string, href string, tableName string) error {
	db := dynamo.New(session.New(), &aws.Config{
		Region: aws.String("ap-northeast-1"),
	})
	table := db.Table(tableName)
	fmt.Println("tableName", tableName)
	if tableName == "portfolio-work" {
		work := Work{WorkID: "001",
			Timestamp: time.Now().UTC(),
			Title:     title,
			Detail:    content,
			Image:     href}

		err := table.Put(work).Run()
		return err
	} else {
		article := Article{ArticleID: "001",
			Timestamp: time.Now().UTC(),
			Title:     title,
			Detail:    content,
			Image:     href}
		fmt.Println("created table!")
		err := table.Put(article).Run()
		return err
	}
}

//GetDynamo get items list
func GetDynamo(tableName string) ([]Work, []Article, error) {
	db := dynamo.New(session.New(), &aws.Config{
		Region: aws.String("ap-northeast-1"),
	})
	table := db.Table(tableName)
	if tableName == "portfolio-work" {
		var results []Work
		err := table.Get("WorkID", "001").All(&results)
		return results, nil, err
	} else {
		var results []Article
		err := table.Get("ArticleID", "001").All(&results)
		return nil, results, err
	}
}
