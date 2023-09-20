package routes

import (
	"net/http"

	"github.com/muchrief/go_pijar/src/controller"
	"github.com/muchrief/go_pijar/src/helper"
	"github.com/muchrief/go_pijar/src/model"

	"github.com/labstack/echo/v4"
)

func AddRoutes(app *echo.Echo) {
	app.Validator = helper.CreateCustomValidator()
	app.RouteNotFound("/*", func(c echo.Context) error {
		return c.JSON(
			http.StatusNotFound,
			model.ResponeApi[interface{}]{
				Code:    http.StatusNotFound,
				Message: "Not Found",
			},
		)
	})
	app.GET("/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, model.ResponeApi[interface{}]{
			Code:    0,
			Message: "Success",
		})
	})

	g := app.Group("/users")

	g.GET("", controller.GetAllUser)
	g.POST("", controller.AddUser)

	g.GET("/:id", controller.GetUser)
}
