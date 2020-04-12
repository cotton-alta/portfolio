package main

import (
	"portfolio-api/controllers"
	"portfolio-api/interceptors"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/api/works", controllers.WorkList(), interceptor.Auth())
	e.PUT("/api/works", controllers.CreateWork())
	e.GET("/api/articles", controllers.ArticleList())
	e.PUT("/api/articles", controllers.CreateArticle())

	e.Start(":3001")
}
