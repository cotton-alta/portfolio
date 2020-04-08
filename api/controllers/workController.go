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
		database.CreateDynamo(title, detail)
		return c.String(http.StatusOK, "created item!")
	}
}

//WorkList GET:/works
func WorkList() echo.HandlerFunc {
	return func(c echo.Context) error {
		works, _ := database.GetDynamo()
		// works := [...]work{
		// 	{
		// 		Title:  "MonsterCounter",
		// 		Detail: "魔剤でランキングを競うアプリです。",
		// 		Href:   "#",
		// 	},
		// 	{
		// 		Title:  "OsiCounter",
		// 		Detail: "推しキャラを愛でながら時間を計るアプリです。",
		// 		Href:   "#",
		// 	},
		// }
		return c.JSON(http.StatusOK, works)
	}
}
