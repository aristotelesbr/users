package routers

import (
	"usuarios/controllers"

	"github.com/labstack/echo"
)

// App is a instance of echo
var App *echo.Echo

// creata a construction function
func init() {
	App = echo.New()

	App.GET("/", controllers.Home)

	api := App.Group("/v1")

	api.GET("/new", controllers.New)
	api.POST("/create", controllers.Create)
	api.DELETE("/delete/:id", controllers.Delete)
}
