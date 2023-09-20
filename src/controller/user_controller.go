package controller

import (
	"net/http"

	"github.com/muchrief/go_pijar/src/model"

	"github.com/labstack/echo/v4"
)

func GetAllUser(c echo.Context) error {
	resp := model.ResponeApi[[]string]{
		Code:    0,
		Message: "Success",
		Data:    []string{"jabranboy", "kiboy", "microboy"},
	}

	return c.JSON(http.StatusOK, resp)
}

func AddUser(c echo.Context) error {
	user := new(model.AddUserPayload)
	err := c.Bind(user)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			model.ResponeApi[interface{}]{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		)
	}

	err = c.Validate(user)
	if err != nil {
		return err
	}
	resp := model.ResponeApi[[]string]{
		Code:    0,
		Message: "Success",
		Data:    []string{"jabranboy", "kiboy", "microboy"},
	}

	return c.JSON(http.StatusOK, resp)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	c.Logger().Info(id)

	resp := model.ResponeApi[string]{
		Code:    0,
		Message: "Success",
		Data:    "Jabranboy",
	}

	return c.JSON(http.StatusOK, resp)
}
