package controllers

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"portfolio-api/database"

	"github.com/labstack/echo"
)

type (
	work struct {
		Title  string `json:"title"`
		Detail string `json:"detail"`
		Href   string `json:"href"`
	}
)

//GetContentType detect file content type
func GetContentType(file multipart.File) (string, error) {
	buffer := make([]byte, 512)

	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)
	return contentType, err
}

//CreateWork PUT:/works
func CreateWork() echo.HandlerFunc {
	return func(c echo.Context) error {
		title := c.FormValue("title")
		detail := c.FormValue("detail")

		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		originalname := file.Filename

		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		contentType, err := GetContentType(src)
		fmt.Println("contentType", contentType)
		if err != nil {
			return err
		}

		href, err := database.CreateObject(src, originalname, contentType)
		database.CreateDynamo(title, detail, href)
		return c.String(http.StatusOK, "created item!")
	}
}

//WorkList GET:/works
func WorkList() echo.HandlerFunc {
	return func(c echo.Context) error {
		works, _ := database.GetDynamo()
		return c.JSON(http.StatusOK, works)
	}
}
