package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

// Home GET /home
func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World 2")
}
