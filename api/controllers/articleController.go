package controllers

import (
	"net/http"
	"fmt"
	"portfolio-api/database"

	"github.com/labstack/echo"
)

type (
	article struct {
		Title  string `json:"title"`
		Detail string `json:"detail"`
		Href   string `json:"href"`
	}
)

// CreateArticle PUT:/articles
func CreateArticle() echo.HandlerFunc {
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
		database.CreateDynamo(title, detail, href, "portfolio-article")
		return c.String(http.StatusOK, "created item!")
	}
}

//ArticleList GET:/articles
func ArticleList() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, articles, _ := database.GetDynamo("portfolio-article")
		return c.JSON(http.StatusOK, articles)
	}
}
