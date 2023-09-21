package main

import (
	"github.com/muchrief/go_pijar/database"
	"github.com/muchrief/go_pijar/src/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	database.InitializeDB()
}

func main() {
	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.Use(middleware.CORS())
	app.Use(middleware.
		RateLimiter(
			middleware.
				NewRateLimiterMemoryStore(1000),
		),
	)

	routes.AddRoutes(app)

	app.Logger.Fatal(app.Start(":1323"))
}
