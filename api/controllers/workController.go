package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	work struct {
		Title  string `json:"title"`
		Detail string `json:"detail"`
		Href   string `json:"href"`
	}
)

//WorkList GET:/hello
func WorkList() echo.HandlerFunc {
	return func(c echo.Context) error {
		works := [...]work{
			{
				Title:  "MonsterCounter",
				Detail: "魔剤でランキングを競うアプリです。",
				Href:   "#",
			},
			{
				Title:  "OsiCounter",
				Detail: "推しキャラを愛でながら時間を計るアプリです。",
				Href:   "#",
			},
		}
		return c.JSON(http.StatusOK, works)
	}
}
