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

	e.GET("/api/works", controllers.WorkList())
	e.GET("/api/articles", controllers.ArticleList())

	e.Start(":3001")
}
