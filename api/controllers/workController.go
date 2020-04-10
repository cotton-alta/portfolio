package controllers

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"portfolio-api/database"

	"github.com/labstack/echo"
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
		work := map[string]string{
			"title": c.FormValue("title"),
			"detail": c.FormValue("detail"),
			"link": c.FormValue("link"),
			"github": c.FormValue("github"),
		}

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
		database.CreateDynamo(work, href, "portfolio-work")
		return c.String(http.StatusOK, "created item!")
	}
}

//WorkList GET:/works
func WorkList() echo.HandlerFunc {
	return func(c echo.Context) error {
		works, _, _ := database.GetDynamo("portfolio-work")
		return c.JSON(http.StatusOK, works)
	}
}
