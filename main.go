package main

import (
	r "usuarios/routers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/middleware"
	_ "upper.io/db.v3"
)

func main() {
	e := r.App
	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(":3000"))
}
