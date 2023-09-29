package main

import (
	"fmt"

	"github.com/muchrief/go_pijar/database"
	"github.com/muchrief/go_pijar/src/api"
	"github.com/muchrief/go_pijar/src/helper"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	database.InitializeDB(database.POSTGRESQL)
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

	api.RegisterApi(app)

	port := helper.LoadEnv("PORT", "3000")

	app.Logger.Fatal(app.Start(fmt.Sprintf("0.0.0.0:%s", port)))
}
