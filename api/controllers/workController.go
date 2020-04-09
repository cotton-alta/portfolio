package controllers

import (
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

//CreateWork PUT:/works
func CreateWork() echo.HandlerFunc {
	return func(c echo.Context) error {
		title := c.FormValue("title")
		detail := c.FormValue("detail")

		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		href, err := database.CreateObject(src)
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
