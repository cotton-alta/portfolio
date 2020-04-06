package controllers

import (
	"net/http"

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
		articles := [...]article{
			{
				Title:  "初投稿",
				Detail: "こちら初投稿です！",
				Href:   "#",
			},
			{
				Title:  "ポートフォリオサイトを公開",
				Detail: "ポートフォリオサイトを公開しました。",
				Href:   "#",
			},
		}
		return c.JSON(http.StatusOK, articles)
	}
}
