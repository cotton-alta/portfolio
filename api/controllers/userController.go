package controllers

import (
	"github.com/labstack/echo"
	"github.com/ashwanthkumar/slack-go-webhook"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "login user!")
	}
}

func SendMessage() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := godotenv.Load()
    if err != nil {
        return err
    }
		emailField := slack.Field{Title: "メールアドレス", Value: c.FormValue("email")}
		contentField := slack.Field{Title: "お問い合わせ内容", Value: c.FormValue("content")}
		
		attachment := slack.Attachment{}
		attachment.AddField(emailField).AddField(contentField)
		color := "good"
		attachment.Color = &color
		payload := slack.Payload{
			Username: os.Getenv("SLACKUSER"),
			Channel: os.Getenv("CHANNEL"),
			Attachments: []slack.Attachment{attachment},
		}
		slack.Send(os.Getenv("WEBHOOKURL"), "", payload)
		return c.String(http.StatusOK, "send message!")
	}
}