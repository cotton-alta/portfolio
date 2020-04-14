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

	e.GET("/api/login", controllers.Login(), interceptor.Auth())
	e.GET("/api/works", controllers.WorkList())
	e.PUT("/api/works", controllers.CreateWork())
	e.GET("/api/articles", controllers.ArticleList())
	e.PUT("/api/articles", controllers.CreateArticle())
	e.GET("/api/articles/:article", controllers.GetArticle())

	e.Start(":3001")
}
