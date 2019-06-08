package main

import (
	r "usuarios/routers"
)

func main() {
	e := r.App

	e.Start(":3000")
}
