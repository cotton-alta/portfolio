package interceptor

import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
	"errors"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
)

func Auth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username string, password string, context echo.Context) (bool, error) {
		err := godotenv.Load()
    if err != nil {
        return false, err
    }
		if username == os.Getenv("USER") && password == os.Getenv("PASSWORD") {
			return true, nil
		}
		err = errors.New("no login")
		fmt.Println(os.Getenv("USER"))
		return false, err
	})
}