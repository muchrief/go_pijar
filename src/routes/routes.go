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

	addUserRoutes(app)
	addCampusRoutes(app)
}

func addUserRoutes(app *echo.Echo) {
	g := app.Group("/users")

	g.POST("", controller.AddUser)
	g.GET("", controller.GetAllUser)
	g.GET("/:id", controller.GetUser)
	g.DELETE("/:id", controller.DeleteUser)
}

func addCampusRoutes(app *echo.Echo) {
	g := app.Group("/campus")

	g.GET("/:id", controller.GetCampusInfo)
	g.GET("/:id/faculties", controller.GetCampusFaculties)
}
