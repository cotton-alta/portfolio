package controllers

import (
	"net/http"
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

//ArticleList GET:/articles
func ArticleList() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, articles, _ := database.GetDynamo("portfolio-article")
		return c.JSON(http.StatusOK, articles)
	}
}
