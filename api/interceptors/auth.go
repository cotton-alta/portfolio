package interceptor

import (
	"os"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
)

func Auth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username string, password string, context echo.Context) (bool, error) {
		if username == os.Getenv("USERNAME") && password == os.Getenv("PASSWORD") {
			return true, nil
		}
		return false, nil
	})
}