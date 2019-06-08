package routers

import (
	"net/http"

	"github.com/labstack/echo"
)

var App *echo.Echo

// creata a construction function
func init() {
	App = echo.New()

	App.GET("/", home)
}
func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World 2")
}
