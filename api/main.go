package main

import (
	"portfolio-api/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello", controllers.MainPage())

	e.Start(":1234")
}
