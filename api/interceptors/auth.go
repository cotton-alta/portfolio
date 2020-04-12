package interceptor

import (
	// "os"
	"errors"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
)

func Auth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username string, password string, context echo.Context) (bool, error) {
		// if username == os.Getenv("USERNAME") && password == os.Getenv("PASSWORD") {
		if username == "username" && password == "password" {
			return true, nil
		}
		err := errors.New("no login")
		return false, err
	})
}