package main

import (
	r "usuarios/routers"

	_ "github.com/flosch/pongo2"

	"github.com/ikeikeikeike/pongor"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := r.App

	p := pongor.GetRenderer()
	p.Directory = "views"
	e.Renderer = p

	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(":3000"))
}
