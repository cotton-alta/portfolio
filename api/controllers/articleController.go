package controllers

import (
	"net/http"
	"fmt"
	"regexp"
	"portfolio-api/database"

	"github.com/labstack/echo"
)

// CreateArticle PUT:/articles
func CreateArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		article := map[string]string{
			"title": c.FormValue("title"),
			"detail": c.FormValue("detail"),
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
		database.CreateDynamo(article, href, "portfolio-article")
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

//GetArticle GET:/atricles/:article
func GetArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		articleName := c.Param("article")
		fmt.Println("articleName ", articleName)
		_, article, _ := database.GetDynamoSingle("portfolio-article", articleName)
		fmt.Println("article ", article)
		return c.JSON(http.StatusOK, article)
	}
}

//DeleteArticle DELETE:/articles
func DeleteArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		articleName := c.Param("article")
		_, article, _ := database.GetDynamoSingle("portfolio-article", articleName)
		reg := regexp.MustCompile(`.+/(.+?)([\?#;].*)?$`)
		itemName := reg.ReplaceAllString(article[0].Image, "$1")
		fmt.Println("itemName", itemName)
		err := database.DeleteObject(itemName)
		err = database.DeleteDynamo("portfolio-article", articleName, article[0].Timestamp)
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, "item deleted!")
	}
}