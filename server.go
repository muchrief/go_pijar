package main

import (
	"github.com/muchrief/go_pijar/src/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	app.Use(middleware.CORS())

	routes.AddRoutes(app)

	app.Logger.Fatal(app.Start(":1323"))
}
