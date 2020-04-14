package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	// "fmt"
)

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		// username := c.FormValue("username")
		// fmt.Println(username)
		return c.String(http.StatusOK, "login user!")
	}
}